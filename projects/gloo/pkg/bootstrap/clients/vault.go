package clients

import (
	"context"
	"fmt"
	"os"

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

func VaultClientForSettings(vaultSettings *v1.Settings_VaultSecrets) (*vault.Client, error) {
	cfg, err := parseVaultSettings(vaultSettings)
	if err != nil {
		return nil, err
	}
	client, err := vault.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return configureVaultAuth(vaultSettings, client)
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

func configureVaultAuth(vaultSettings *v1.Settings_VaultSecrets, client *vault.Client) (*vault.Client, error) {
	// each case returns
	switch tlsCfg := vaultSettings.GetAuthMethod().(type) {
	case *v1.Settings_VaultSecrets_AccessToken:
		client.SetToken(tlsCfg.AccessToken)
		return client, nil
	case *v1.Settings_VaultSecrets_Aws:
		return configureAwsAuth(tlsCfg.Aws, client)
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
func configureAwsAuth(aws *v1.Settings_VaultAwsAuth, client *vault.Client) (*vault.Client, error) {
	return configureAwsIamAuth(aws, client)
}

func configureAwsIamAuth(aws *v1.Settings_VaultAwsAuth, client *vault.Client) (*vault.Client, error) {
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

	authInfo, err := client.Auth().Login(context.Background(), awsAuth)
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
	go renewToken(client, awsAuth, int(aws.GetWatcherIncrement()))

	return client, nil
}

// Once you've set the token for your Vault client, you will need to periodically renew its lease.
// taken from https://github.com/hashicorp/vault-examples/blob/main/examples/token-renewal/go/example.go
func renewToken(client *vault.Client, awsAuth *awsauth.AWSAuth, watcherIncrement int) {
	for {
		vaultLoginResp, err := client.Auth().Login(context.Background(), awsAuth)
		if err != nil {
			contextutils.LoggerFrom(context.Background()).Fatalf("unable to authenticate to Vault: %v", err)
		}
		tokenErr := manageTokenLifecycle(client, vaultLoginResp, watcherIncrement)
		if tokenErr != nil {
			contextutils.LoggerFrom(context.Background()).Fatalf("unable to start managing token lifecycle: %v", tokenErr)
		}
	}
}

// Starts token lifecycle management. Returns only fatal errors as errors,
// otherwise returns nil so we can attempt login again.
// taken from https://github.com/hashicorp/vault-examples/blob/main/examples/token-renewal/go/example.go
func manageTokenLifecycle(client *vault.Client, token *vault.Secret, watcherIncrement int) error {
	renew := token.Auth.Renewable // You may notice a different top-level field called Renewable. That one is used for dynamic secrets renewal, not token renewal.
	if !renew {
		contextutils.LoggerFrom(context.Background()).Infof("Token is not configured to be renewable. Re-attempting login.")
		return nil
	}

	lifetimeWatcherInput := &vault.LifetimeWatcherInput{
		Secret: token,
	}

	if watcherIncrement > 0 {
		lifetimeWatcherInput.Increment = watcherIncrement
	}

	watcher, err := client.NewLifetimeWatcher(lifetimeWatcherInput)
	if err != nil {
		return fmt.Errorf("unable to initialize new lifetime watcher for renewing auth token: %w", err)
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
				contextutils.LoggerFrom(context.Background()).Infof("Failed to renew token: %v. Re-attempting login.", err)
				return nil
			}
			// This occurs once the token has reached max TTL.
			contextutils.LoggerFrom(context.Background()).Infof("Token can no longer be renewed. Re-attempting login.")
			return nil

		// Successfully completed renewal
		case renewal := <-watcher.RenewCh():
			contextutils.LoggerFrom(context.Background()).Infof("Successfully renewed: %#v.", renewal)
		}
	}
}
