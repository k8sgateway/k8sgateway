package vault

import (
    "context"
    "errors"
    "github.com/rotisserie/eris"
    "github.com/solo-io/gloo/pkg/utils/awsutils"
    "time"

    "github.com/avast/retry-go"
    vault "github.com/hashicorp/vault/api"
    awsauth "github.com/hashicorp/vault/api/auth/aws"
    "github.com/solo-io/gloo/pkg/utils"
    v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
    "github.com/solo-io/go-utils/contextutils"
)

// In an ideal world, we would re-use the mocks provided by an external library.
// Since the vault.AuthMethod interface does not have corresponding mocks, we have to define our own.
//go:generate mockgen -destination mocks/mock_auth.go -package mocks github.com/hashicorp/vault/api AuthMethod

type ClientAuth interface {
    vault.AuthMethod
    StartRenewal(ctx context.Context, secret *vault.Secret) error
}

var _ ClientAuth = &staticTokenAuth{}
var _ ClientAuth = &remoteTokenAuth{}

var (
    ErrEmptyToken = errors.New("unable to authenticate to vault with empty token. check Settings configuration")
    ErrNoAuthInfo = errors.New("no auth info was returned after login")
)

// newClientAuthForSettings returns a vault ClientAuth based on the provided settings.
func newClientAuthForSettings(ctx context.Context, vaultSettings *v1.Settings_VaultSecrets) (ClientAuth, error) {
    switch tlsCfg := vaultSettings.GetAuthMethod().(type) {
    case *v1.Settings_VaultSecrets_AccessToken:
        return newStaticTokenAuth(tlsCfg.AccessToken), nil

    case *v1.Settings_VaultSecrets_Aws:
        awsAuth, err := newAwsAuthMethod(ctx, tlsCfg.Aws)
        if err != nil {
            return nil, err
        }

        return newRemoteTokenAuth(awsAuth), nil

    default:
        // AuthMethod is the preferred API to define the policy for authenticating to vault
        // If one is not defined, we fall back to the deprecated API
        return newStaticTokenAuth(vaultSettings.GetToken()), nil
    }
}

func newStaticTokenAuth(token string) ClientAuth {
    return &staticTokenAuth{
        token: token,
    }
}

type staticTokenAuth struct {
    token string
}

func (s *staticTokenAuth) StartRenewal(_ context.Context, _ *vault.Secret) error {
    // static tokens do not support renewal
    return nil
}

func (s *staticTokenAuth) Login(ctx context.Context, _ *vault.Client) (*vault.Secret, error) {
    if s.token == "" {
        utils.Measure(ctx, mLastLoginFailure, time.Now().Unix())
        utils.MeasureOne(ctx, mLoginFailures)
        return nil, ErrEmptyToken
    }

    contextutils.LoggerFrom(ctx).Debug("successfully authenticated to vault with static token")
    utils.Measure(ctx, mLastLoginSuccess, time.Now().Unix())
    utils.MeasureOne(ctx, mLoginSuccesses)
    return &vault.Secret{
        Auth: &vault.SecretAuth{
            ClientToken: s.token,
        },
    }, nil
}

func newRemoteTokenAuth(authMethod vault.AuthMethod, retryOptions ...retry.Option) ClientAuth {
    // Standard retry options, which can be overridden by the loginRetryOptions parameter
    defaultRetryOptions := []retry.Option{
        retry.Delay(1 * time.Second),
        retry.DelayType(retry.BackOffDelay),
        retry.Attempts(10),
        retry.LastErrorOnly(true),
    }
    loginRetryOptions := append(defaultRetryOptions, retryOptions...)

    return &remoteTokenAuth{
        authMethod:        authMethod,
        loginRetryOptions: loginRetryOptions,
    }
}

type remoteTokenAuth struct {
    authMethod        vault.AuthMethod
    loginRetryOptions []retry.Option
}

func (r *remoteTokenAuth) Login(ctx context.Context, client *vault.Client) (*vault.Secret, error) {
    var (
        loginResponse *vault.Secret
        loginErr      error
    )

    loginErr = retry.Do(func() error {
        loginResponse, loginErr = r.loginOnce(ctx, client)
        return loginErr
    }, r.loginRetryOptions...)

    return loginResponse, loginErr
}

func (r *remoteTokenAuth) loginOnce(ctx context.Context, client *vault.Client) (*vault.Secret, error) {
    loginResponse, loginErr := r.authMethod.Login(ctx, client)
    if loginErr != nil {
        contextutils.LoggerFrom(ctx).Errorf("unable to authenticate to vault: %v", loginErr)
        utils.Measure(ctx, mLastLoginFailure, time.Now().Unix())
        utils.MeasureOne(ctx, mLoginFailures)
        return nil, eris.Wrapf(loginErr, "unable to authenticate to vault")
    }

    if loginResponse == nil {
        contextutils.LoggerFrom(ctx).Error(ErrNoAuthInfo)
        utils.Measure(ctx, mLastLoginFailure, time.Now().Unix())
        utils.MeasureOne(ctx, mLoginFailures)
        return nil, ErrNoAuthInfo
    }

    contextutils.LoggerFrom(ctx).Debugf("successfully authenticated to vault %v", loginResponse)
    utils.Measure(ctx, mLastLoginSuccess, time.Now().Unix())
    utils.MeasureOne(ctx, mLoginSuccesses)
    return loginResponse, nil
}

func (r *remoteTokenAuth) StartRenewal(ctx context.Context, secret *vault.Secret) error {
    // todo - implement renewal
    return nil
}

func newAwsAuthMethod(_ context.Context, aws *v1.Settings_VaultAwsAuth) (*awsauth.AWSAuth, error) {
    // The AccessKeyID and SecretAccessKey are not required in the case of using temporary credentials from assumed roles with AWS STS or IRSA.
    // STS: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html
    // IRSA: https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
    var possibleErrStrings []string
    if accessKeyId := aws.GetAccessKeyId(); accessKeyId != "" {
        awsutils.SetAccessKeyEnv(accessKeyId)
    } else {
        possibleErrStrings = append(possibleErrStrings, "access key id must be defined for AWS IAM auth")
    }

    if secretAccessKey := aws.GetSecretAccessKey(); secretAccessKey != "" {
        awsutils.SetSecretAccessKeyEnv(secretAccessKey)
    } else {
        possibleErrStrings = append(possibleErrStrings, "secret access key must be defined for AWS IAM auth")
    }

    // if we have only partial configuration set
    if len(possibleErrStrings) == 1 {
        return nil, errors.New("only partial credentials were provided for AWS IAM auth: " + possibleErrStrings[0])
    }

    // At this point, we either have full auth configuration set, or are in an ec2 environment, where vault will infer the credentials.
    loginOptions := []awsauth.LoginOption{awsauth.WithIAMAuth()}

    if role := aws.GetVaultRole(); role != "" {
        loginOptions = append(loginOptions, awsauth.WithRole(role))
    }

    if region := aws.GetRegion(); region != "" {
        loginOptions = append(loginOptions, awsauth.WithRegion(region))
    }

    if iamServerIdHeader := aws.GetIamServerIdHeader(); iamServerIdHeader != "" {
        loginOptions = append(loginOptions, awsauth.WithIAMServerIDHeader(iamServerIdHeader))
    }

    if mountPath := aws.GetMountPath(); mountPath != "" {
        loginOptions = append(loginOptions, awsauth.WithMountPath(mountPath))
    }

    if sessionToken := aws.GetSessionToken(); sessionToken != "" {
        awsutils.SetSessionTokenEnv(sessionToken)
    }

    return awsauth.NewAWSAuth(loginOptions...)
}
