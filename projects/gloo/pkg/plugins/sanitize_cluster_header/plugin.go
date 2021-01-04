package sanitize_cluster_header

import (
	"github.com/rotisserie/eris"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

const (
	errEnterpriseOnly = "Could not load sanitize_cluster_header plugin - this is an Enterprise feature"
)

var (
	sanitizeFilterStage = plugins.BeforeStage(plugins.AuthNStage)
)

type plugin struct {
}

var (
	_ plugins.Plugin           = new(plugin)
	_ plugins.HttpFilterPlugin = new(plugin)
)

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) IsUpgrade() bool {
	return false
}

func (p *plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *plugin) HttpFilters(params plugins.Params, listener *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {
	if listener.GetOptions().GetSanitizeClusterHeader() != nil {
		return nil, eris.New(errEnterpriseOnly)
	}
	return nil, nil
}
