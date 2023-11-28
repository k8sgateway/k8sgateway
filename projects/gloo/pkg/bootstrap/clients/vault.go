package clients

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"

	vault "github.com/hashicorp/vault/api"
	awsauth "github.com/hashicorp/vault/api/auth/aws"
	errors "github.com/rotisserie/eris"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// The DefaultPathPrefix may be overridden to allow for non-standard vault mount paths
const DefaultPathPrefix = "secret"

type VaultClientInitFunc func() *vault.Client

func NoopVaultClientInitFunc(c *vault.Client) VaultClientInitFunc {
	return func() *vault.Client {
		return c
	}
}

var (
	ErrNilVaultClient = errors.New("vault API client failed to initialize")
)

// NewVaultSecretClientFactory consumes a vault client along with a set of basic configurations for retrieving info with the client
func NewVaultSecretClientFactory(clientInit VaultClientInitFunc, pathPrefix, rootKey string) factory.ResourceClientFactory {
	return &factory.VaultSecretClientFactory{
		Vault:      clientInit(),
		RootKey:    rootKey,
		PathPrefix: pathPrefix,
	}
}

func VaultClientForSettings(ctx context.Context, vaultSettings *v1.Settings_VaultSecrets) (*vault.Client, error) {
	cfg, err := parseVaultSettings(vaultSettings)
	if err != nil {
		return nil, err
	}
	client, err := vault.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return configureVaultAuth(ctx, vaultSettings, client)
}

func parseVaultSettings(vaultSettings *v1.Settings_VaultSecrets) (*vault.Config, error) {
	cfg := vault.DefaultConfig()

	if addr := vaultSettings.GetAddress(); addr != "" {
		cfg.Address = addr
	}
	if tlsConfig := parseTlsSettings(vaultSettings); tlsConfig != nil {
		if err := cfg.ConfigureTLS(tlsConfig); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func parseTlsSettings(vaultSettings *v1.Settings_VaultSecrets) *vault.TLSConfig {
	var tlsConfig *vault.TLSConfig

	// helper functions to avoid repeated nilchecking
	addStringSetting := func(s string, addSettingFunc func(string)) {
		if s == "" {
			return
		}
		if tlsConfig == nil {
			tlsConfig = &vault.TLSConfig{}
		}
		addSettingFunc(s)
	}
	addBoolSetting := func(b *wrapperspb.BoolValue, addSettingFunc func(bool)) {
		if b == nil {
			return
		}
		if tlsConfig == nil {
			tlsConfig = &vault.TLSConfig{}
		}
		addSettingFunc(b.GetValue())
	}

	setCaCert := func(s string) { tlsConfig.CACert = s }
	setCaPath := func(s string) { tlsConfig.CAPath = s }
	setClientCert := func(s string) { tlsConfig.ClientCert = s }
	setClientKey := func(s string) { tlsConfig.ClientKey = s }
	setTlsServerName := func(s string) { tlsConfig.TLSServerName = s }
	setInsecure := func(b bool) { tlsConfig.Insecure = b }

	// Add our settings to the vault TLS config, preferring settings set in the
	// new TlsConfig field if it is used to those in the deprecated fields
	if tlsSettings := vaultSettings.GetTlsConfig(); tlsSettings == nil {
		addStringSetting(vaultSettings.GetCaCert(), setCaCert)
		addStringSetting(vaultSettings.GetCaPath(), setCaPath)
		addStringSetting(vaultSettings.GetClientCert(), setClientCert)
		addStringSetting(vaultSettings.GetClientKey(), setClientKey)
		addStringSetting(vaultSettings.GetTlsServerName(), setTlsServerName)
		addBoolSetting(vaultSettings.GetInsecure(), setInsecure)
	} else {
		addStringSetting(vaultSettings.GetTlsConfig().GetCaCert(), setCaCert)
		addStringSetting(vaultSettings.GetTlsConfig().GetCaPath(), setCaPath)
		addStringSetting(vaultSettings.GetTlsConfig().GetClientCert(), setClientCert)
		addStringSetting(vaultSettings.GetTlsConfig().GetClientKey(), setClientKey)
		addStringSetting(vaultSettings.GetTlsConfig().GetTlsServerName(), setTlsServerName)
		addBoolSetting(vaultSettings.GetTlsConfig().GetInsecure(), setInsecure)
	}

	return tlsConfig

}

func configureVaultAuth(ctx context.Context, vaultSettings *v1.Settings_VaultSecrets, client *vault.Client) (*vault.Client, error) {
	// each case returns
	switch tlsCfg := vaultSettings.GetAuthMethod().(type) {
	case *v1.Settings_VaultSecrets_AccessToken:
		client.SetToken(tlsCfg.AccessToken)
		return client, nil
	case *v1.Settings_VaultSecrets_Aws:
		return configureAwsAuth(ctx, tlsCfg.Aws, client)
	default:
		// We don't have one of the defined auth methods, so try to fall back to the
		// deprecated token field before erroring
		token := vaultSettings.GetToken()
		if token == "" {
			return nil, errors.Errorf("unable to determine vault authentication method. check Settings configuration")
		}
		client.SetToken(token)
		return client, nil
	}
}

// This indirection function exists to more easily enable further extenstion of AWS auth
// to support EC2 auth method in the future
func configureAwsAuth(ctx context.Context, aws *v1.Settings_VaultAwsAuth, client *vault.Client) (*vault.Client, error) {
	return configureAwsIamAuth(ctx, aws, client)
}

func configureAwsIamAuth(ctx context.Context, aws *v1.Settings_VaultAwsAuth, client *vault.Client) (*vault.Client, error) {
	// The AccessKeyID and SecretAccessKey are not required in the case of using temporary credentials from assumed roles with AWS STS or IRSA.
	// STS: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html
	// IRSA: https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
	var possibleErrStrings []string
	if accessKeyId := aws.GetAccessKeyId(); accessKeyId != "" {
		os.Setenv("AWS_ACCESS_KEY_ID", accessKeyId)
	} else {
		possibleErrStrings = append(possibleErrStrings, "access key id must be defined for AWS IAM auth")
	}

	if secretAccessKey := aws.GetSecretAccessKey(); secretAccessKey != "" {
		os.Setenv("AWS_SECRET_ACCESS_KEY", secretAccessKey)
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
		os.Setenv("AWS_SESSION_TOKEN", sessionToken)
	}

	awsAuth, err := awsauth.NewAWSAuth(loginOptions...)
	if err != nil {
		return nil, err
	}

	authInfo, err := client.Auth().Login(ctx, awsAuth)
	if err != nil {
		err := errors.Wrapf(err, "unable to login to AWS auth method")
		// if using inferred credentials, add error information regarding setting credentials
		if len(possibleErrStrings) > 0 {
			err = errors.Wrapf(err, "using implicit credentials, consider setting aws secret access key and access key id")
		}

		return nil, err
	}
	if authInfo == nil {
		return nil, errors.New("no auth info was returned after login")
	}

	// set up auth token refreshing with client.NewLifetimeWatcher()
	go renewToken(ctx, client, awsAuth, int(aws.GetLeaseIncrement()))

	return client, nil
}

var (
	errVaultAuthentication = errors.New("unable to authenticate to Vault")
	errNoAuthInfo          = errors.New("no auth info was returned after login")
	//errTokenRenewal        = errors.New("unable to start managing token lifecycle")
)

// Once you've set the token for your Vault client, you will need to periodically renew its lease.
// taken from https://github.com/hashicorp/vault-examples/blob/main/examples/token-renewal/go/example.go
func renewToken(ctx context.Context, client *vault.Client, awsAuth *awsauth.AWSAuth, watcherIncrement int) {
	contextutils.LoggerFrom(ctx).Infof("Starting renewToken goroutine")

	// var count = 0
	for {
		var vaultLoginResp *vault.Secret
		err := contextutils.NewExponentialBackoff(contextutils.ExponentialBackoff{}).Backoff(ctx, func(ctx context.Context) error {
			var vaultErr error
			vaultLoginResp, vaultErr = client.Auth().Login(ctx, awsAuth)
			if vaultErr != nil {
				//contextutils.LoggerFrom(ctx).Fatalf("unable to authenticate to Vault: %v", vaultErr)
				contextutils.LoggerFrom(ctx).Errorf("unable to authenticate to Vault: %v", vaultErr)
				return errVaultAuthentication
			}
			if vaultLoginResp == nil {
				//contextutils.LoggerFrom(ctx).Fatalf("no auth info was returned after login")
				contextutils.LoggerFrom(ctx).Errorf("no auth info was returned after login")
				return errNoAuthInfo
			}
			// if count < 3 {
			// 	count++
			// 	contextutils.LoggerFrom(ctx).Errorf("ERROR: Testing retry %d", count)
			// 	return errors.Errorf("ERROR: Testing retry %d", count)
			// }

			return nil
		})
		// We should never hit these errors because the retry logic has no max time or retries, so it will retry until err is nil
		if err != nil {
			if ctx.Err() != nil {
				contextutils.LoggerFrom(ctx).Fatalf("renew token context error: %v", ctx.Err())
			} else {
				contextutils.LoggerFrom(ctx).Fatalf("unable to authenticate to Vault: %v", err)
			}
		}
		contextutils.LoggerFrom(ctx).Infof("Successfully authenticated to Vault %v", vaultLoginResp)

		retry, tokenErr := manageTokenLifecycle(ctx, client, vaultLoginResp, watcherIncrement)

		// The only error this function can return is if the vaultLoginResp is nil, and we have checked against that
		if tokenErr != nil {
			contextutils.LoggerFrom(ctx).Fatalf("unable to start managing token lifecycle: %v", tokenErr)
		}

		if !retry {
			// contextutils.LoggerFrom(ctx).Infof("Stopping renewToken goroutine")
			return
		}

	}

}

// Starts token lifecycle management. Returns only fatal errors as errors,
// otherwise returns nil so we can attempt login again.
// based on https://github.com/hashicorp/vault-examples/blob/main/examples/token-renewal/go/example.go
func manageTokenLifecycle(ctx context.Context, client *vault.Client, secret *vault.Secret, watcherIncrement int) (bool, error) {
	if renewable, err := secret.TokenIsRenewable(); !renewable || err != nil {
		if err != nil {
			contextutils.LoggerFrom(ctx).Debugf("Error in checking if token is renewable: %v. Re-attempting login.", err)
		} else {
			contextutils.LoggerFrom(ctx).Debugf("Token is not configured to be renewable. Re-attempting login.")
		}
		return true, nil
	}

	watcher, err := client.NewLifetimeWatcher(&vault.LifetimeWatcherInput{
		Secret:    secret,
		Increment: watcherIncrement,
		// Below this comment we are manually setting the parameters to current defaults to protect against future changes
		Rand:          rand.New(rand.NewSource(int64(time.Now().Nanosecond()))),
		RenewBuffer:   5, // equivalent to vault.DefaultLifetimeWatcherRenewBuffer,
		RenewBehavior: vault.RenewBehaviorIgnoreErrors,
	})

	// The only errors the constructor can return are if the input parameter is nil or if the secret is nil, and we
	// are always passing input and have validated the secret is not nil in the calling
	if err != nil {
		return false, fmt.Errorf("unable to initialize new lifetime watcher for renewing auth token: %w", err)
	}

	go watcher.Start()
	defer watcher.Stop()

	for {
		select {
		// `DoneCh` will return if renewal fails, or if the remaining lease
		// duration is under a built-in threshold and either renewing is not
		// extending it or renewing is disabled. In any case, the caller
		// needs to attempt to log in again.
		case err := <-watcher.DoneCh():
			if err != nil {
				contextutils.LoggerFrom(ctx).Debugf("Failed to renew token: %v. Re-attempting login.", err)
				return true, nil
			}
			// This occurs once the token has reached max TTL.
			contextutils.LoggerFrom(ctx).Debugf("Token can no longer be renewed. Re-attempting login.")
			return true, nil

		// Successfully completed renewal
		case renewal := <-watcher.RenewCh():
			contextutils.LoggerFrom(ctx).Debugf("Successfully renewed: %v.", renewal)

		case <-ctx.Done():
			return false, nil
		}

	}
}
