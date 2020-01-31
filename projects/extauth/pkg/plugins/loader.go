package plugins

import (
	"context"
	"fmt"
	"path/filepath"
	"plugin"
	"reflect"
	"strings"

	"github.com/gogo/protobuf/types"
	errors "github.com/rotisserie/eris"
	"github.com/solo-io/ext-auth-plugins/api"
	extauth "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/protoutils"
	"go.uber.org/zap"
)

//go:generate mockgen -destination mocks/loader_mock.go -package mocks github.com/solo-io/solo-projects/projects/extauth/pkg/plugins Loader

var (
	PluginFileOpenError = func(err error) error {
		return errors.Wrapf(err, "failed to open plugin file")
	}
	InvalidExportedSymbolError = func(err error) error {
		return errors.Wrapf(err, "failed to find exported plugin symbol")
	}
)

func NewPluginLoader(pluginDir string) Loader {
	return &loader{
		pluginDir: pluginDir,
	}
}

type Loader interface {
	LoadAuthPlugin(ctx context.Context, pluginConfig *extauth.AuthPlugin) (api.AuthService, error)
}

type loader struct {
	pluginDir string
}

func (l *loader) LoadAuthPlugin(ctx context.Context, pluginConfig *extauth.AuthPlugin) (api.AuthService, error) {
	logger := contextutils.LoggerFrom(ctx)

	// Load plugin
	logger.Debugw("Loading ext-auth plugin", "name", pluginConfig.Name)
	extAuthPlugin, err := l.loadPlugin(ctx, pluginConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load plugin")
	}

	// Get object to deserialize plugin config into
	logger.Debugw("Getting new instance of plugin configuration object")
	pluginCfg, err := extAuthPlugin.NewConfigInstance(ctx)
	if err != nil {
		return nil, err
	}

	// Deserialize plugin config into object
	logger.Debugw("Trying to unmarshal plugin config into configuration object",
		zap.Any("pluginConfig", pluginConfig.Config),
		zap.Any("configurationObject", pluginCfg),
	)
	if err = parsePluginConfig(pluginConfig.Config, pluginCfg); err != nil {
		return nil, errors.Wrapf(err, "failed to parse config for plugin [%s]. "+
			"Could not deserialize config: %v into configuration object %v", pluginConfig.Name, pluginConfig.Config, pluginCfg)
	}

	logger.Debugw("Getting AuthService instance from plugin", zap.Any("pluginConfig", pluginCfg))
	authService, err := extAuthPlugin.GetAuthService(ctx, pluginCfg)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get AuthService instance from plugin")
	}

	return authService, nil
}

func (l *loader) loadPlugin(ctx context.Context, authPlugin *extauth.AuthPlugin) (api.ExtAuthPlugin, error) {
	logger := contextutils.LoggerFrom(ctx)

	pluginName := authPlugin.Name
	if pluginName == "" {
		return nil, errors.New("plugin is missing required [Name] field")
	}

	// Default file name to <name>.so
	pluginFileName := fmt.Sprintf("%s.so", strings.TrimSuffix(pluginName, ".so"))
	if authPlugin.PluginFileName != "" {
		pluginFileName = authPlugin.PluginFileName
	}
	pluginFilePath := filepath.Join(l.pluginDir, pluginFileName)

	// Default symbol name to <name>
	symbolName := pluginName
	if authPlugin.ExportedSymbolName != "" {
		symbolName = authPlugin.ExportedSymbolName
	}

	logger.Debugw("Loading plugin", zap.Any("name", pluginName), zap.Any("path", pluginFilePath))
	goPlugin, err := plugin.Open(pluginFilePath)
	if err != nil {
		return nil, PluginFileOpenError(err)
	}

	logger.Debugw("Looking up exported plugin symbol", zap.Any("symbolName", symbolName))
	exportedSymbol, err := goPlugin.Lookup(symbolName)
	if err != nil {
		return nil, InvalidExportedSymbolError(err)
	}

	// A Symbol is a pointer to a variable or function.
	// We first check if the underlying object implements the plugin interface.
	// If not, we check if the symbol itself implements the plugin interface. This can happen if e.g. the exported
	// symbol is a concrete struct, but it implements the interface via pointer receivers. Let's be tolerant.
	logger.Debugw("Checking type of exported symbol", zap.Any("symbol", exportedSymbol))
	underlyingObj, err := unpack(exportedSymbol)
	if err != nil {
		// Should never happen, as Lookup always returns a pointer
		return nil, errors.Wrapf(err, "panicked while trying to unpack symbol [%s]", symbolName)
	}
	pluginInterfaceInstance, ok := underlyingObj.(api.ExtAuthPlugin)
	if !ok {
		pluginInterfaceInstance, ok = exportedSymbol.(api.ExtAuthPlugin)
		if !ok {
			return nil, errors.Errorf("invalid plugin interface implementation: type [%T] does not implement "+
				"the [api.ExtAuthPlugin] interface", exportedSymbol)
		}
	}

	return pluginInterfaceInstance, nil
}

func parsePluginConfig(rawConfig *types.Struct, typedConfig interface{}) error {
	return protoutils.UnmarshalStruct(rawConfig, typedConfig)
}

// A Symbol is a pointer to a variable or function. This function returns the object it is pointing to.
// Panics i.a. if the argument is not a pointer or an interface.
func unpack(symbol plugin.Symbol) (i interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			i = nil
			err = errors.Errorf("%v", r)
		}
	}()
	return reflect.ValueOf(symbol).Elem().Interface(), nil
}
