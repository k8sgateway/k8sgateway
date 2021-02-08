package flagutils

import (
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/solo-projects/projects/glooctl-extensions/fed/pkg/cmd/options"
	"github.com/spf13/pflag"
)

func AddNamespaceFlag(set *pflag.FlagSet, opt *options.Options) {
	set.StringVarP(&opt.Namespace, "namespace", "n", defaults.GlooFed, "namespace for reading or writing resources")
}
