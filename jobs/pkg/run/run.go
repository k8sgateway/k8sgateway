package run

import (
	"context"
	"time"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/jobs/pkg/certgen"
	"github.com/solo-io/gloo/jobs/pkg/kube"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/k8s-utils/certutils"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// Options to configure the rotation job
// Certgen job yaml files have opinionated defaults for these options
type Options struct {
	SvcName      string
	SvcNamespace string

	SecretName      string
	SecretNamespace string
	NextSecretName  string

	ServerCertSecretFileName    string
	ServerCertAuthorityFileName string
	ServerKeySecretFileName     string

	ValidatingWebhookConfigurationName string

	ForceRotation bool

	RenewBefore string

	// The duration waited after first updating a secret's cabundle before
	// updating the actual secrets.
	// Lower values make the rotation job run faster
	// Higher values make the rotation job more resilient to errors
	RotationDuration string
}

func Run(ctx context.Context, opts Options) error {
	if opts.SvcNamespace == "" {
		return eris.Errorf("must provide svc-namespace")
	}
	if opts.SvcName == "" {
		return eris.Errorf("must provide svc-name")
	}
	if opts.SecretNamespace == "" {
		return eris.Errorf("must provide secret-namespace")
	}
	if opts.SecretName == "" {
		return eris.Errorf("must provide secret-name")
	}
	if opts.NextSecretName == "" {
		opts.NextSecretName = opts.SecretName + "-next"
	}
	if opts.ServerCertSecretFileName == "" {
		return eris.Errorf("must provide name for the server cert entry in the secret data")
	}
	if opts.ServerCertAuthorityFileName == "" {
		return eris.Errorf("must provide name for the cert authority entry in the secret data")
	}
	if opts.ServerKeySecretFileName == "" {
		return eris.Errorf("must provide name for the server key entry in the secret data")
	}
	renewBeforeDuration, err := time.ParseDuration(opts.RenewBefore)
	if err != nil {
		return err
	}
	rotationDuration, err := time.ParseDuration(opts.RotationDuration)
	if err != nil {
		return err
	}

	logger := contextutils.LoggerFrom(ctx)
	logger.Info("starting cert rotation job, validated inputs")

	kubeClient := helpers.MustKubeClient()

	var secret *v1.Secret
	// check if there is an existing valid TLS secret
	secret, renewCurrent, err := kube.GetExistingValidTlsSecret(ctx, kubeClient, opts.SecretName, opts.SecretNamespace,
		opts.SvcName, opts.SvcNamespace, renewBeforeDuration)
	if err != nil {
		return eris.Wrapf(err, "failed validating existing secret, consider deleting it and rerunning the job")
	}
	nextSecret, renewNext, err := kube.GetExistingValidTlsSecret(ctx, kubeClient, opts.NextSecretName, opts.SecretNamespace,
		opts.SvcName, opts.SvcNamespace, renewBeforeDuration)
	if err != nil {
		return eris.Wrapf(err, "failed validating next secret, consider deleting it and rerunning the job")
	}

	logger.Infof("checked for existing secrets, currentSecretExists:%v shouldrenew:%t  nextSecretExists:%v shouldrenew:%t", secret != nil, renewCurrent, nextSecret == nil, renewNext)

	// If either secret is empty or invalid, generate two new secrets and save them.
	if secret == nil || nextSecret == nil {

		// standup the follow up secret.
		// this is needed to have a handle on its ca for the cabundle
		nextCerts, err := certgen.GenCerts(opts.SvcName, opts.SvcNamespace)
		if err != nil {
			return eris.Wrapf(err, "failed creating next secret")
		}

		nextSecretConfig := fillTlsSecret(opts, opts.NextSecretName, nextCerts.ServerCertKey, nextCerts.CaCertificate, nextCerts.CaCertificate)

		_, err = kube.CreateTlsSecret(ctx, kubeClient, nextSecretConfig)
		if err != nil {
			return eris.Wrapf(err, "error saving next secret")
		}

		var certs *certutils.Certificates
		if secret != nil && !renewCurrent {
			// In this case we have a valid unexpired secret
			// and we havent gotten to the force rotation flag check yet
			// DO_NOT_SUBMIT: We need to make sure that we can do a smooth run here if the next secret is mistyped however for now we dont do anyhting in this

			// Here we should pull the secret out, update its cabundle like in swap secrets
			// then store it back

			// secret.CaBundle = append(secret.CaBundle, nextCerts.CaCertificate...)
			// return persistWebhook(ctx, opts, kubeClient, secret)
		}
		certs, err = certgen.GenCerts(opts.SvcName, opts.SvcNamespace)
		if err != nil {
			return eris.Wrapf(err, "failed creating secret")
		}

		newSecretConfig := fillTlsSecret(opts, opts.SecretName, certs.ServerCertKey, certs.CaCertificate, append(certs.CaCertificate, nextCerts.CaCertificate...))
		secret, err = kube.CreateTlsSecret(ctx, kubeClient, newSecretConfig)
		if err != nil {
			return eris.Wrapf(err, "error saving secret")
		}

		return persistWebhook(ctx, opts, kubeClient, secret)
	}
	// Rotate out the older cert and add a newer one
	if opts.ForceRotation || renewCurrent || renewNext {
		contextutils.LoggerFrom(ctx).Infow("Rotating secrets regardless of expiration")
		certs, err := certgen.GenCerts(opts.SvcName, opts.SvcNamespace)
		if err != nil {
			return eris.Wrapf(err, "generating self-signed certs and key")
		}
		nextSecretConfig := parseTlsSecret(nextSecret, opts.ServerKeySecretFileName, opts.ServerCertSecretFileName, opts.ServerCertAuthorityFileName)
		secretConfig := parseTlsSecret(secret, opts.ServerKeySecretFileName, opts.ServerCertSecretFileName, opts.ServerCertAuthorityFileName)

		newSecretConfig := fillTlsSecret(opts, opts.NextSecretName, certs.ServerCertKey, certs.CaCertificate, certs.CaCertificate)

		secret, err = kube.SwapSecrets(ctx, rotationDuration, kubeClient, secretConfig, nextSecretConfig, newSecretConfig)
		if err != nil {
			return eris.Wrapf(err, "failed creating or rotating secret")
		}
		return persistWebhook(ctx, opts, kubeClient, secret)
	} else {
		contextutils.LoggerFrom(ctx).Infow("existing TLS secret found, skipping update to TLS secret since the old TLS secret is still valid",
			zap.String("secretName", opts.SecretName),
			zap.String("secretNamespace", opts.SecretNamespace))
	}

	return nil
}

// fill a kube tls secret with the provided data
// a helper as we reuse the options and it gets bloaty
func fillTlsSecret(opt Options, name string, key, certChain, caBundle []byte) kube.TlsSecret {
	return kube.TlsSecret{
		SecretName:         name,
		SecretNamespace:    opt.SecretNamespace,
		PrivateKeyFileName: opt.ServerKeySecretFileName,
		CertFileName:       opt.ServerCertSecretFileName,
		CaBundleFileName:   opt.ServerCertAuthorityFileName,
		PrivateKey:         key,
		Cert:               certChain, // should be server with ca appended
		CaBundle:           caBundle,
	}
}

func persistWebhook(ctx context.Context, opts Options, kubeClient kubernetes.Interface, secret *v1.Secret) error {

	vwcName := opts.ValidatingWebhookConfigurationName
	if vwcName == "" {
		contextutils.LoggerFrom(ctx).Infof("no ValidatingWebhookConfiguration provided. finished successfully.")
		return nil
	}

	safeBundle := make([]byte, len(secret.Data[opts.ServerCertAuthorityFileName]))
	copy(safeBundle, secret.Data[opts.ServerCertAuthorityFileName])
	vwcConfig := kube.WebhookTlsConfig{
		ServiceName:      opts.SvcName,
		ServiceNamespace: opts.SvcNamespace,
		CaBundle:         safeBundle,
	}

	if err := kube.UpdateValidatingWebhookConfigurationCaBundle(ctx, kubeClient, vwcName, vwcConfig); err != nil {
		return eris.Wrapf(err, "failed patching validating webhook config")
	}

	contextutils.LoggerFrom(ctx).Infof("finished successfully.")
	return nil
}
func parseTlsSecret(args *v1.Secret, privateKeyFilename, certFileName, caBundleFileName string) kube.TlsSecret {
	return kube.TlsSecret{
		SecretName:         args.GetObjectMeta().GetName(),
		SecretNamespace:    args.GetObjectMeta().GetNamespace(),
		PrivateKeyFileName: privateKeyFilename,
		CertFileName:       certFileName,
		CaBundleFileName:   caBundleFileName,
		PrivateKey:         args.Data[privateKeyFilename],
		Cert:               args.Data[certFileName],
		CaBundle:           args.Data[caBundleFileName],
	}
}
