package prerun

import (
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/go-utils/errors"
)

func EnableConsulClients(opts *options.Options) error {
	consul := opts.Top.Consul
	if consul.UseConsul {
		client, err := consul.Client()
		if err != nil {
			return errors.Wrapf(err, "creating Consul client")
		}
		helpers.UseConsulClients(client, consul.RootKey)
	}
	return nil
}

func EnableVaultClients(vault options.Vault) error {
	if vault.UseVault {
		client, err := vault.Client()
		if err != nil {
			return errors.Wrapf(err, "creating Vault client")
		}
		helpers.UseVaultClients(client, vault.RootKey)
	}
	return nil
}
