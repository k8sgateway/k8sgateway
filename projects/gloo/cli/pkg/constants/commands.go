package constants

import "github.com/spf13/cobra"

var (
	VIRTUAL_SERVICE_COMMAND = cobra.Command{
		Use:     "virtualservice",
		Aliases: []string{"vs", "virtualservices"},
	}

	UPSTREAM_COMMAND = cobra.Command{
		Use:     "upstream",
		Aliases: []string{"u", "us", "upstreams"},
	}

	PROXY_COMMAND = cobra.Command{
		Use:     "proxy",
		Aliases: []string{"p", "proxies"},
	}

	ROUTE_COMMAND = cobra.Command{
		Use:     "route",
		Aliases: []string{"r", "routes"},
	}

	GATEWAY_COMMAND = cobra.Command{
		Use:     "gateway",
		Aliases: []string{"gw"},
		Short:   "interact with the Gloo Gateway/Ingress",
	}

	ADD_COMMAND = cobra.Command{
		Use:     "add",
		Aliases: []string{"a"},
		Short:   "adds configuration to a top-level Gloo resource",
	}

	CREATE_COMMAND = cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a Gloo resource",
		Long:    "Gloo resources be created from files (including stdin)",
	}

	DELETE_COMMAND = cobra.Command{
		Use:     "delete",
		Aliases: []string{"d"},
		Short:   "Delete a Gloo resource",
	}

	GET_COMMAND = cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Display one or a list of Gloo resources",
	}

	INSTALL_COMMAND = cobra.Command{
		Use:   "install",
		Short: "install gloo on different platforms",
		Long:  "choose which system to install Gloo onto. options include: kubernetes",
	}

	UPGRADE_COMMAND = cobra.Command{
		Use:     "upgrade",
		Aliases: []string{"ug"},
		Short:   "upgrade glooctl binary",
	}


)
