/* eslint-disable */
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');
var extproto_ext_pb = require('../../../../../../../extproto/ext_pb.js');
var github_com_solo$io_solo$kit_api_v1_solo$kit_pb = require('../../../../../../../github.com/solo-io/solo-kit/api/v1/solo-kit_pb.js');
var github_com_solo$io_solo$kit_api_v1_ref_pb = require('../../../../../../../github.com/solo-io/solo-kit/api/v1/ref_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_extensions_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/extensions_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/ratelimit/ratelimit_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_caching_caching_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/caching/caching_pb.js');
var github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/enterprise.gloo/v1/auth_config_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_rbac_rbac_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/rbac/rbac_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_circuit_breaker_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/circuit_breaker_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_ssl_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/ssl_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_extensions_aws_filter_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/aws/filter_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_consul_query_options_pb = require('../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/consul/query_options_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
goog.exportSymbol('proto.gloo.solo.io.ConsoleOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.GatewayOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.GatewayOptions.ValidationOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.GlooOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.GlooOptions.AWSOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy', null, global);
goog.exportSymbol('proto.gloo.solo.io.GraphqlOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsNamespacedStatuses', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.ConsulConfiguration', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.ConsulKv', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.Directory', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.DiscoveryOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.FdsMode', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.KnativeOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.KubernetesCrds', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.KubernetesSecrets', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.ObservabilityOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsSpec.VaultSecrets', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsStatus', null, global);
goog.exportSymbol('proto.gloo.solo.io.SettingsStatus.State', null, global);
goog.exportSymbol('proto.gloo.solo.io.UpstreamOptions', null, global);
goog.exportSymbol('proto.gloo.solo.io.VirtualServiceOptions', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.gloo.solo.io.SettingsSpec.repeatedFields_, proto.gloo.solo.io.SettingsSpec.oneofGroups_);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.displayName = 'proto.gloo.solo.io.SettingsSpec';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.gloo.solo.io.SettingsSpec.repeatedFields_ = [2];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.gloo.solo.io.SettingsSpec.oneofGroups_ = [[4,5,21],[6,7,8],[9,10,23]];

/**
 * @enum {number}
 */
proto.gloo.solo.io.SettingsSpec.ConfigSourceCase = {
  CONFIG_SOURCE_NOT_SET: 0,
  KUBERNETES_CONFIG_SOURCE: 4,
  DIRECTORY_CONFIG_SOURCE: 5,
  CONSUL_KV_SOURCE: 21
};

/**
 * @return {proto.gloo.solo.io.SettingsSpec.ConfigSourceCase}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getConfigSourceCase = function() {
  return /** @type {proto.gloo.solo.io.SettingsSpec.ConfigSourceCase} */(jspb.Message.computeOneofCase(this, proto.gloo.solo.io.SettingsSpec.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.gloo.solo.io.SettingsSpec.SecretSourceCase = {
  SECRET_SOURCE_NOT_SET: 0,
  KUBERNETES_SECRET_SOURCE: 6,
  VAULT_SECRET_SOURCE: 7,
  DIRECTORY_SECRET_SOURCE: 8
};

/**
 * @return {proto.gloo.solo.io.SettingsSpec.SecretSourceCase}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getSecretSourceCase = function() {
  return /** @type {proto.gloo.solo.io.SettingsSpec.SecretSourceCase} */(jspb.Message.computeOneofCase(this, proto.gloo.solo.io.SettingsSpec.oneofGroups_[1]));
};

/**
 * @enum {number}
 */
proto.gloo.solo.io.SettingsSpec.ArtifactSourceCase = {
  ARTIFACT_SOURCE_NOT_SET: 0,
  KUBERNETES_ARTIFACT_SOURCE: 9,
  DIRECTORY_ARTIFACT_SOURCE: 10,
  CONSUL_KV_ARTIFACT_SOURCE: 23
};

/**
 * @return {proto.gloo.solo.io.SettingsSpec.ArtifactSourceCase}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getArtifactSourceCase = function() {
  return /** @type {proto.gloo.solo.io.SettingsSpec.ArtifactSourceCase} */(jspb.Message.computeOneofCase(this, proto.gloo.solo.io.SettingsSpec.oneofGroups_[2]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    discoveryNamespace: jspb.Message.getFieldWithDefault(msg, 1, ""),
    watchNamespacesList: jspb.Message.getRepeatedField(msg, 2),
    kubernetesConfigSource: (f = msg.getKubernetesConfigSource()) && proto.gloo.solo.io.SettingsSpec.KubernetesCrds.toObject(includeInstance, f),
    directoryConfigSource: (f = msg.getDirectoryConfigSource()) && proto.gloo.solo.io.SettingsSpec.Directory.toObject(includeInstance, f),
    consulKvSource: (f = msg.getConsulKvSource()) && proto.gloo.solo.io.SettingsSpec.ConsulKv.toObject(includeInstance, f),
    kubernetesSecretSource: (f = msg.getKubernetesSecretSource()) && proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.toObject(includeInstance, f),
    vaultSecretSource: (f = msg.getVaultSecretSource()) && proto.gloo.solo.io.SettingsSpec.VaultSecrets.toObject(includeInstance, f),
    directorySecretSource: (f = msg.getDirectorySecretSource()) && proto.gloo.solo.io.SettingsSpec.Directory.toObject(includeInstance, f),
    kubernetesArtifactSource: (f = msg.getKubernetesArtifactSource()) && proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.toObject(includeInstance, f),
    directoryArtifactSource: (f = msg.getDirectoryArtifactSource()) && proto.gloo.solo.io.SettingsSpec.Directory.toObject(includeInstance, f),
    consulKvArtifactSource: (f = msg.getConsulKvArtifactSource()) && proto.gloo.solo.io.SettingsSpec.ConsulKv.toObject(includeInstance, f),
    refreshRate: (f = msg.getRefreshRate()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    devMode: jspb.Message.getFieldWithDefault(msg, 13, false),
    linkerd: jspb.Message.getFieldWithDefault(msg, 17, false),
    knative: (f = msg.getKnative()) && proto.gloo.solo.io.SettingsSpec.KnativeOptions.toObject(includeInstance, f),
    discovery: (f = msg.getDiscovery()) && proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.toObject(includeInstance, f),
    gloo: (f = msg.getGloo()) && proto.gloo.solo.io.GlooOptions.toObject(includeInstance, f),
    gateway: (f = msg.getGateway()) && proto.gloo.solo.io.GatewayOptions.toObject(includeInstance, f),
    consul: (f = msg.getConsul()) && proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.toObject(includeInstance, f),
    consuldiscovery: (f = msg.getConsuldiscovery()) && proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.toObject(includeInstance, f),
    kubernetes: (f = msg.getKubernetes()) && proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.toObject(includeInstance, f),
    extensions: (f = msg.getExtensions()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_extensions_pb.Extensions.toObject(includeInstance, f),
    ratelimit: (f = msg.getRatelimit()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.ServiceSettings.toObject(includeInstance, f),
    ratelimitServer: (f = msg.getRatelimitServer()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.Settings.toObject(includeInstance, f),
    rbac: (f = msg.getRbac()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_rbac_rbac_pb.Settings.toObject(includeInstance, f),
    extauth: (f = msg.getExtauth()) && github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.Settings.toObject(includeInstance, f),
    namedExtauthMap: (f = msg.getNamedExtauthMap()) ? f.toObject(includeInstance, proto.enterprise.gloo.solo.io.Settings.toObject) : [],
    cachingServer: (f = msg.getCachingServer()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_caching_caching_pb.Settings.toObject(includeInstance, f),
    observabilityoptions: (f = msg.getObservabilityoptions()) && proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.toObject(includeInstance, f),
    upstreamoptions: (f = msg.getUpstreamoptions()) && proto.gloo.solo.io.UpstreamOptions.toObject(includeInstance, f),
    consoleOptions: (f = msg.getConsoleOptions()) && proto.gloo.solo.io.ConsoleOptions.toObject(includeInstance, f),
    graphqlOptions: (f = msg.getGraphqlOptions()) && proto.gloo.solo.io.GraphqlOptions.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec}
 */
proto.gloo.solo.io.SettingsSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec;
  return proto.gloo.solo.io.SettingsSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec}
 */
proto.gloo.solo.io.SettingsSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setDiscoveryNamespace(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addWatchNamespaces(value);
      break;
    case 4:
      var value = new proto.gloo.solo.io.SettingsSpec.KubernetesCrds;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.KubernetesCrds.deserializeBinaryFromReader);
      msg.setKubernetesConfigSource(value);
      break;
    case 5:
      var value = new proto.gloo.solo.io.SettingsSpec.Directory;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.Directory.deserializeBinaryFromReader);
      msg.setDirectoryConfigSource(value);
      break;
    case 21:
      var value = new proto.gloo.solo.io.SettingsSpec.ConsulKv;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.ConsulKv.deserializeBinaryFromReader);
      msg.setConsulKvSource(value);
      break;
    case 6:
      var value = new proto.gloo.solo.io.SettingsSpec.KubernetesSecrets;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.deserializeBinaryFromReader);
      msg.setKubernetesSecretSource(value);
      break;
    case 7:
      var value = new proto.gloo.solo.io.SettingsSpec.VaultSecrets;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.VaultSecrets.deserializeBinaryFromReader);
      msg.setVaultSecretSource(value);
      break;
    case 8:
      var value = new proto.gloo.solo.io.SettingsSpec.Directory;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.Directory.deserializeBinaryFromReader);
      msg.setDirectorySecretSource(value);
      break;
    case 9:
      var value = new proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.deserializeBinaryFromReader);
      msg.setKubernetesArtifactSource(value);
      break;
    case 10:
      var value = new proto.gloo.solo.io.SettingsSpec.Directory;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.Directory.deserializeBinaryFromReader);
      msg.setDirectoryArtifactSource(value);
      break;
    case 23:
      var value = new proto.gloo.solo.io.SettingsSpec.ConsulKv;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.ConsulKv.deserializeBinaryFromReader);
      msg.setConsulKvArtifactSource(value);
      break;
    case 12:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setRefreshRate(value);
      break;
    case 13:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDevMode(value);
      break;
    case 17:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setLinkerd(value);
      break;
    case 18:
      var value = new proto.gloo.solo.io.SettingsSpec.KnativeOptions;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.KnativeOptions.deserializeBinaryFromReader);
      msg.setKnative(value);
      break;
    case 19:
      var value = new proto.gloo.solo.io.SettingsSpec.DiscoveryOptions;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.deserializeBinaryFromReader);
      msg.setDiscovery(value);
      break;
    case 24:
      var value = new proto.gloo.solo.io.GlooOptions;
      reader.readMessage(value,proto.gloo.solo.io.GlooOptions.deserializeBinaryFromReader);
      msg.setGloo(value);
      break;
    case 25:
      var value = new proto.gloo.solo.io.GatewayOptions;
      reader.readMessage(value,proto.gloo.solo.io.GatewayOptions.deserializeBinaryFromReader);
      msg.setGateway(value);
      break;
    case 20:
      var value = new proto.gloo.solo.io.SettingsSpec.ConsulConfiguration;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.deserializeBinaryFromReader);
      msg.setConsul(value);
      break;
    case 30:
      var value = new proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.deserializeBinaryFromReader);
      msg.setConsuldiscovery(value);
      break;
    case 22:
      var value = new proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.deserializeBinaryFromReader);
      msg.setKubernetes(value);
      break;
    case 16:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_extensions_pb.Extensions;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_extensions_pb.Extensions.deserializeBinaryFromReader);
      msg.setExtensions(value);
      break;
    case 26:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.ServiceSettings;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.ServiceSettings.deserializeBinaryFromReader);
      msg.setRatelimit(value);
      break;
    case 27:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.Settings;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.Settings.deserializeBinaryFromReader);
      msg.setRatelimitServer(value);
      break;
    case 28:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_rbac_rbac_pb.Settings;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_rbac_rbac_pb.Settings.deserializeBinaryFromReader);
      msg.setRbac(value);
      break;
    case 29:
      var value = new github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.Settings;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.Settings.deserializeBinaryFromReader);
      msg.setExtauth(value);
      break;
    case 33:
      var value = msg.getNamedExtauthMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.enterprise.gloo.solo.io.Settings.deserializeBinaryFromReader, "");
         });
      break;
    case 36:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_caching_caching_pb.Settings;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_caching_caching_pb.Settings.deserializeBinaryFromReader);
      msg.setCachingServer(value);
      break;
    case 31:
      var value = new proto.gloo.solo.io.SettingsSpec.ObservabilityOptions;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.deserializeBinaryFromReader);
      msg.setObservabilityoptions(value);
      break;
    case 32:
      var value = new proto.gloo.solo.io.UpstreamOptions;
      reader.readMessage(value,proto.gloo.solo.io.UpstreamOptions.deserializeBinaryFromReader);
      msg.setUpstreamoptions(value);
      break;
    case 35:
      var value = new proto.gloo.solo.io.ConsoleOptions;
      reader.readMessage(value,proto.gloo.solo.io.ConsoleOptions.deserializeBinaryFromReader);
      msg.setConsoleOptions(value);
      break;
    case 37:
      var value = new proto.gloo.solo.io.GraphqlOptions;
      reader.readMessage(value,proto.gloo.solo.io.GraphqlOptions.deserializeBinaryFromReader);
      msg.setGraphqlOptions(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDiscoveryNamespace();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getWatchNamespacesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getKubernetesConfigSource();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.gloo.solo.io.SettingsSpec.KubernetesCrds.serializeBinaryToWriter
    );
  }
  f = message.getDirectoryConfigSource();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.gloo.solo.io.SettingsSpec.Directory.serializeBinaryToWriter
    );
  }
  f = message.getConsulKvSource();
  if (f != null) {
    writer.writeMessage(
      21,
      f,
      proto.gloo.solo.io.SettingsSpec.ConsulKv.serializeBinaryToWriter
    );
  }
  f = message.getKubernetesSecretSource();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.serializeBinaryToWriter
    );
  }
  f = message.getVaultSecretSource();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      proto.gloo.solo.io.SettingsSpec.VaultSecrets.serializeBinaryToWriter
    );
  }
  f = message.getDirectorySecretSource();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      proto.gloo.solo.io.SettingsSpec.Directory.serializeBinaryToWriter
    );
  }
  f = message.getKubernetesArtifactSource();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.serializeBinaryToWriter
    );
  }
  f = message.getDirectoryArtifactSource();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      proto.gloo.solo.io.SettingsSpec.Directory.serializeBinaryToWriter
    );
  }
  f = message.getConsulKvArtifactSource();
  if (f != null) {
    writer.writeMessage(
      23,
      f,
      proto.gloo.solo.io.SettingsSpec.ConsulKv.serializeBinaryToWriter
    );
  }
  f = message.getRefreshRate();
  if (f != null) {
    writer.writeMessage(
      12,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getDevMode();
  if (f) {
    writer.writeBool(
      13,
      f
    );
  }
  f = message.getLinkerd();
  if (f) {
    writer.writeBool(
      17,
      f
    );
  }
  f = message.getKnative();
  if (f != null) {
    writer.writeMessage(
      18,
      f,
      proto.gloo.solo.io.SettingsSpec.KnativeOptions.serializeBinaryToWriter
    );
  }
  f = message.getDiscovery();
  if (f != null) {
    writer.writeMessage(
      19,
      f,
      proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.serializeBinaryToWriter
    );
  }
  f = message.getGloo();
  if (f != null) {
    writer.writeMessage(
      24,
      f,
      proto.gloo.solo.io.GlooOptions.serializeBinaryToWriter
    );
  }
  f = message.getGateway();
  if (f != null) {
    writer.writeMessage(
      25,
      f,
      proto.gloo.solo.io.GatewayOptions.serializeBinaryToWriter
    );
  }
  f = message.getConsul();
  if (f != null) {
    writer.writeMessage(
      20,
      f,
      proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.serializeBinaryToWriter
    );
  }
  f = message.getConsuldiscovery();
  if (f != null) {
    writer.writeMessage(
      30,
      f,
      proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.serializeBinaryToWriter
    );
  }
  f = message.getKubernetes();
  if (f != null) {
    writer.writeMessage(
      22,
      f,
      proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.serializeBinaryToWriter
    );
  }
  f = message.getExtensions();
  if (f != null) {
    writer.writeMessage(
      16,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_extensions_pb.Extensions.serializeBinaryToWriter
    );
  }
  f = message.getRatelimit();
  if (f != null) {
    writer.writeMessage(
      26,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.ServiceSettings.serializeBinaryToWriter
    );
  }
  f = message.getRatelimitServer();
  if (f != null) {
    writer.writeMessage(
      27,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.Settings.serializeBinaryToWriter
    );
  }
  f = message.getRbac();
  if (f != null) {
    writer.writeMessage(
      28,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_rbac_rbac_pb.Settings.serializeBinaryToWriter
    );
  }
  f = message.getExtauth();
  if (f != null) {
    writer.writeMessage(
      29,
      f,
      github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.Settings.serializeBinaryToWriter
    );
  }
  f = message.getNamedExtauthMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(33, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.enterprise.gloo.solo.io.Settings.serializeBinaryToWriter);
  }
  f = message.getCachingServer();
  if (f != null) {
    writer.writeMessage(
      36,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_caching_caching_pb.Settings.serializeBinaryToWriter
    );
  }
  f = message.getObservabilityoptions();
  if (f != null) {
    writer.writeMessage(
      31,
      f,
      proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.serializeBinaryToWriter
    );
  }
  f = message.getUpstreamoptions();
  if (f != null) {
    writer.writeMessage(
      32,
      f,
      proto.gloo.solo.io.UpstreamOptions.serializeBinaryToWriter
    );
  }
  f = message.getConsoleOptions();
  if (f != null) {
    writer.writeMessage(
      35,
      f,
      proto.gloo.solo.io.ConsoleOptions.serializeBinaryToWriter
    );
  }
  f = message.getGraphqlOptions();
  if (f != null) {
    writer.writeMessage(
      37,
      f,
      proto.gloo.solo.io.GraphqlOptions.serializeBinaryToWriter
    );
  }
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.KubernetesCrds = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.KubernetesCrds, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.KubernetesCrds.displayName = 'proto.gloo.solo.io.SettingsSpec.KubernetesCrds';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesCrds.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.KubernetesCrds.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesCrds} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesCrds.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesCrds}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesCrds.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.KubernetesCrds;
  return proto.gloo.solo.io.SettingsSpec.KubernetesCrds.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesCrds} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesCrds}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesCrds.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesCrds.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.KubernetesCrds.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesCrds} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesCrds.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.KubernetesSecrets = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.KubernetesSecrets, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.displayName = 'proto.gloo.solo.io.SettingsSpec.KubernetesSecrets';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesSecrets} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesSecrets}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.KubernetesSecrets;
  return proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesSecrets} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesSecrets}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesSecrets} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesSecrets.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.VaultSecrets, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.VaultSecrets.displayName = 'proto.gloo.solo.io.SettingsSpec.VaultSecrets';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.VaultSecrets.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.VaultSecrets} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.toObject = function(includeInstance, msg) {
  var f, obj = {
    token: jspb.Message.getFieldWithDefault(msg, 1, ""),
    address: jspb.Message.getFieldWithDefault(msg, 2, ""),
    caCert: jspb.Message.getFieldWithDefault(msg, 3, ""),
    caPath: jspb.Message.getFieldWithDefault(msg, 4, ""),
    clientCert: jspb.Message.getFieldWithDefault(msg, 5, ""),
    clientKey: jspb.Message.getFieldWithDefault(msg, 6, ""),
    tlsServerName: jspb.Message.getFieldWithDefault(msg, 7, ""),
    insecure: (f = msg.getInsecure()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    rootKey: jspb.Message.getFieldWithDefault(msg, 9, ""),
    pathPrefix: jspb.Message.getFieldWithDefault(msg, 10, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.VaultSecrets}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.VaultSecrets;
  return proto.gloo.solo.io.SettingsSpec.VaultSecrets.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.VaultSecrets} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.VaultSecrets}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setCaCert(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setCaPath(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setClientCert(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setClientKey(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setTlsServerName(value);
      break;
    case 8:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setInsecure(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setRootKey(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setPathPrefix(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.VaultSecrets.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.VaultSecrets} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getCaCert();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCaPath();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getClientCert();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getClientKey();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getTlsServerName();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getInsecure();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getRootKey();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getPathPrefix();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
};


/**
 * optional string token = 1;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string address = 2;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setAddress = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string ca_cert = 3;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getCaCert = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setCaCert = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string ca_path = 4;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getCaPath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setCaPath = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string client_cert = 5;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getClientCert = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setClientCert = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string client_key = 6;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getClientKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setClientKey = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string tls_server_name = 7;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getTlsServerName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setTlsServerName = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional google.protobuf.BoolValue insecure = 8;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getInsecure = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 8));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setInsecure = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.clearInsecure = function() {
  this.setInsecure(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.hasInsecure = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional string root_key = 9;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getRootKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setRootKey = function(value) {
  jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string path_prefix = 10;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.getPathPrefix = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.VaultSecrets.prototype.setPathPrefix = function(value) {
  jspb.Message.setProto3StringField(this, 10, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.ConsulKv = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.ConsulKv, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.ConsulKv.displayName = 'proto.gloo.solo.io.SettingsSpec.ConsulKv';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.ConsulKv.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.ConsulKv.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulKv} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ConsulKv.toObject = function(includeInstance, msg) {
  var f, obj = {
    rootKey: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.ConsulKv}
 */
proto.gloo.solo.io.SettingsSpec.ConsulKv.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.ConsulKv;
  return proto.gloo.solo.io.SettingsSpec.ConsulKv.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulKv} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.ConsulKv}
 */
proto.gloo.solo.io.SettingsSpec.ConsulKv.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRootKey(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.ConsulKv.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.ConsulKv.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulKv} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ConsulKv.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRootKey();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string root_key = 1;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulKv.prototype.getRootKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulKv.prototype.setRootKey = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.displayName = 'proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps;
  return proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.Directory = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.Directory, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.Directory.displayName = 'proto.gloo.solo.io.SettingsSpec.Directory';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.Directory.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.Directory.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.Directory} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.Directory.toObject = function(includeInstance, msg) {
  var f, obj = {
    directory: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.Directory}
 */
proto.gloo.solo.io.SettingsSpec.Directory.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.Directory;
  return proto.gloo.solo.io.SettingsSpec.Directory.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.Directory} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.Directory}
 */
proto.gloo.solo.io.SettingsSpec.Directory.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setDirectory(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.Directory.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.Directory.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.Directory} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.Directory.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDirectory();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string directory = 1;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.Directory.prototype.getDirectory = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.Directory.prototype.setDirectory = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.KnativeOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.KnativeOptions.displayName = 'proto.gloo.solo.io.SettingsSpec.KnativeOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.KnativeOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.KnativeOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    clusterIngressProxyAddress: jspb.Message.getFieldWithDefault(msg, 1, ""),
    knativeExternalProxyAddress: jspb.Message.getFieldWithDefault(msg, 2, ""),
    knativeInternalProxyAddress: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.KnativeOptions}
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.KnativeOptions;
  return proto.gloo.solo.io.SettingsSpec.KnativeOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.KnativeOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.KnativeOptions}
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setClusterIngressProxyAddress(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setKnativeExternalProxyAddress(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setKnativeInternalProxyAddress(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.KnativeOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.KnativeOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClusterIngressProxyAddress();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getKnativeExternalProxyAddress();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getKnativeInternalProxyAddress();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string cluster_ingress_proxy_address = 1;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.prototype.getClusterIngressProxyAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.prototype.setClusterIngressProxyAddress = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string knative_external_proxy_address = 2;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.prototype.getKnativeExternalProxyAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.prototype.setKnativeExternalProxyAddress = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string knative_internal_proxy_address = 3;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.prototype.getKnativeInternalProxyAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.KnativeOptions.prototype.setKnativeInternalProxyAddress = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.DiscoveryOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.displayName = 'proto.gloo.solo.io.SettingsSpec.DiscoveryOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    fdsMode: jspb.Message.getFieldWithDefault(msg, 1, 0),
    udsOptions: (f = msg.getUdsOptions()) && proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.DiscoveryOptions;
  return proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.FdsMode} */ (reader.readEnum());
      msg.setFdsMode(value);
      break;
    case 2:
      var value = new proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.deserializeBinaryFromReader);
      msg.setUdsOptions(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFdsMode();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getUdsOptions();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.FdsMode = {
  BLACKLIST: 0,
  WHITELIST: 1,
  DISABLED: 2
};


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.displayName = 'proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    enabled: (f = msg.getEnabled()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    watchLabelsMap: (f = msg.getWatchLabelsMap()) ? f.toObject(includeInstance, undefined) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions;
  return proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setEnabled(value);
      break;
    case 2:
      var value = msg.getWatchLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEnabled();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getWatchLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(2, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * optional google.protobuf.BoolValue enabled = 1;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.prototype.getEnabled = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 1));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.prototype.setEnabled = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.prototype.clearEnabled = function() {
  this.setEnabled(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.prototype.hasEnabled = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * map<string, string> watch_labels = 2;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.prototype.getWatchLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 2, opt_noLazyCreate,
      null));
};


proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions.prototype.clearWatchLabelsMap = function() {
  this.getWatchLabelsMap().clear();
};


/**
 * optional FdsMode fds_mode = 1;
 * @return {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.FdsMode}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.prototype.getFdsMode = function() {
  return /** @type {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.FdsMode} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.FdsMode} value */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.prototype.setFdsMode = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional UdsOptions uds_options = 2;
 * @return {?proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.prototype.getUdsOptions = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions, 2));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.UdsOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.prototype.setUdsOptions = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.prototype.clearUdsOptions = function() {
  this.setUdsOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.DiscoveryOptions.prototype.hasUdsOptions = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.ConsulConfiguration, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.displayName = 'proto.gloo.solo.io.SettingsSpec.ConsulConfiguration';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.toObject = function(includeInstance, msg) {
  var f, obj = {
    address: jspb.Message.getFieldWithDefault(msg, 1, ""),
    datacenter: jspb.Message.getFieldWithDefault(msg, 2, ""),
    username: jspb.Message.getFieldWithDefault(msg, 3, ""),
    password: jspb.Message.getFieldWithDefault(msg, 4, ""),
    token: jspb.Message.getFieldWithDefault(msg, 5, ""),
    caFile: jspb.Message.getFieldWithDefault(msg, 6, ""),
    caPath: jspb.Message.getFieldWithDefault(msg, 7, ""),
    certFile: jspb.Message.getFieldWithDefault(msg, 8, ""),
    keyFile: jspb.Message.getFieldWithDefault(msg, 9, ""),
    insecureSkipVerify: (f = msg.getInsecureSkipVerify()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    waitTime: (f = msg.getWaitTime()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    serviceDiscovery: (f = msg.getServiceDiscovery()) && proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.toObject(includeInstance, f),
    httpAddress: jspb.Message.getFieldWithDefault(msg, 13, ""),
    dnsAddress: jspb.Message.getFieldWithDefault(msg, 14, ""),
    dnsPollingInterval: (f = msg.getDnsPollingInterval()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.ConsulConfiguration;
  return proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setDatacenter(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setCaFile(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setCaPath(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setCertFile(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setKeyFile(value);
      break;
    case 10:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setInsecureSkipVerify(value);
      break;
    case 11:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setWaitTime(value);
      break;
    case 12:
      var value = new proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.deserializeBinaryFromReader);
      msg.setServiceDiscovery(value);
      break;
    case 13:
      var value = /** @type {string} */ (reader.readString());
      msg.setHttpAddress(value);
      break;
    case 14:
      var value = /** @type {string} */ (reader.readString());
      msg.setDnsAddress(value);
      break;
    case 15:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setDnsPollingInterval(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDatacenter();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getCaFile();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getCaPath();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getCertFile();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getKeyFile();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getInsecureSkipVerify();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getWaitTime();
  if (f != null) {
    writer.writeMessage(
      11,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getServiceDiscovery();
  if (f != null) {
    writer.writeMessage(
      12,
      f,
      proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.serializeBinaryToWriter
    );
  }
  f = message.getHttpAddress();
  if (f.length > 0) {
    writer.writeString(
      13,
      f
    );
  }
  f = message.getDnsAddress();
  if (f.length > 0) {
    writer.writeString(
      14,
      f
    );
  }
  f = message.getDnsPollingInterval();
  if (f != null) {
    writer.writeMessage(
      15,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.repeatedFields_, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.displayName = 'proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    dataCentersList: jspb.Message.getRepeatedField(msg, 1)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions;
  return proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addDataCenters(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDataCentersList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string data_centers = 1;
 * @return {!Array<string>}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.prototype.getDataCentersList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/** @param {!Array<string>} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.prototype.setDataCentersList = function(value) {
  jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.prototype.addDataCenters = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions.prototype.clearDataCentersList = function() {
  this.setDataCentersList([]);
};


/**
 * optional string address = 1;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setAddress = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string datacenter = 2;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getDatacenter = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setDatacenter = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string username = 3;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setUsername = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string password = 4;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setPassword = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string token = 5;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string ca_file = 6;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getCaFile = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setCaFile = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string ca_path = 7;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getCaPath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setCaPath = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string cert_file = 8;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getCertFile = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setCertFile = function(value) {
  jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string key_file = 9;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getKeyFile = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setKeyFile = function(value) {
  jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional google.protobuf.BoolValue insecure_skip_verify = 10;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getInsecureSkipVerify = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 10));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setInsecureSkipVerify = function(value) {
  jspb.Message.setWrapperField(this, 10, value);
};


proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.clearInsecureSkipVerify = function() {
  this.setInsecureSkipVerify(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.hasInsecureSkipVerify = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional google.protobuf.Duration wait_time = 11;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getWaitTime = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 11));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setWaitTime = function(value) {
  jspb.Message.setWrapperField(this, 11, value);
};


proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.clearWaitTime = function() {
  this.setWaitTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.hasWaitTime = function() {
  return jspb.Message.getField(this, 11) != null;
};


/**
 * optional ServiceDiscoveryOptions service_discovery = 12;
 * @return {?proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getServiceDiscovery = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions, 12));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.ServiceDiscoveryOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setServiceDiscovery = function(value) {
  jspb.Message.setWrapperField(this, 12, value);
};


proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.clearServiceDiscovery = function() {
  this.setServiceDiscovery(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.hasServiceDiscovery = function() {
  return jspb.Message.getField(this, 12) != null;
};


/**
 * optional string http_address = 13;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getHttpAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 13, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setHttpAddress = function(value) {
  jspb.Message.setProto3StringField(this, 13, value);
};


/**
 * optional string dns_address = 14;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getDnsAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 14, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setDnsAddress = function(value) {
  jspb.Message.setProto3StringField(this, 14, value);
};


/**
 * optional google.protobuf.Duration dns_polling_interval = 15;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.getDnsPollingInterval = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 15));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.setDnsPollingInterval = function(value) {
  jspb.Message.setWrapperField(this, 15, value);
};


proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.clearDnsPollingInterval = function() {
  this.setDnsPollingInterval(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.ConsulConfiguration.prototype.hasDnsPollingInterval = function() {
  return jspb.Message.getField(this, 15) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.repeatedFields_, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.displayName = 'proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.repeatedFields_ = [22];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.toObject = function(includeInstance, msg) {
  var f, obj = {
    usetlstagging: jspb.Message.getFieldWithDefault(msg, 16, false),
    tlstagname: jspb.Message.getFieldWithDefault(msg, 17, ""),
    rootca: (f = msg.getRootca()) && github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.toObject(includeInstance, f),
    splittlsservices: jspb.Message.getFieldWithDefault(msg, 19, false),
    consistencymode: jspb.Message.getFieldWithDefault(msg, 20, 0),
    queryOptions: (f = msg.getQueryOptions()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_consul_query_options_pb.QueryOptions.toObject(includeInstance, f),
    serviceTagsAllowlistList: jspb.Message.getRepeatedField(msg, 22)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration;
  return proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 16:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUsetlstagging(value);
      break;
    case 17:
      var value = /** @type {string} */ (reader.readString());
      msg.setTlstagname(value);
      break;
    case 18:
      var value = new github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef;
      reader.readMessage(value,github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.deserializeBinaryFromReader);
      msg.setRootca(value);
      break;
    case 19:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSplittlsservices(value);
      break;
    case 20:
      var value = /** @type {!proto.consul.options.gloo.solo.io.ConsulConsistencyModes} */ (reader.readEnum());
      msg.setConsistencymode(value);
      break;
    case 21:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_consul_query_options_pb.QueryOptions;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_consul_query_options_pb.QueryOptions.deserializeBinaryFromReader);
      msg.setQueryOptions(value);
      break;
    case 22:
      var value = /** @type {string} */ (reader.readString());
      msg.addServiceTagsAllowlist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUsetlstagging();
  if (f) {
    writer.writeBool(
      16,
      f
    );
  }
  f = message.getTlstagname();
  if (f.length > 0) {
    writer.writeString(
      17,
      f
    );
  }
  f = message.getRootca();
  if (f != null) {
    writer.writeMessage(
      18,
      f,
      github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.serializeBinaryToWriter
    );
  }
  f = message.getSplittlsservices();
  if (f) {
    writer.writeBool(
      19,
      f
    );
  }
  f = message.getConsistencymode();
  if (f !== 0.0) {
    writer.writeEnum(
      20,
      f
    );
  }
  f = message.getQueryOptions();
  if (f != null) {
    writer.writeMessage(
      21,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_consul_query_options_pb.QueryOptions.serializeBinaryToWriter
    );
  }
  f = message.getServiceTagsAllowlistList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      22,
      f
    );
  }
};


/**
 * optional bool useTlsTagging = 16;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.getUsetlstagging = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 16, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.setUsetlstagging = function(value) {
  jspb.Message.setProto3BooleanField(this, 16, value);
};


/**
 * optional string tlsTagName = 17;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.getTlstagname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 17, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.setTlstagname = function(value) {
  jspb.Message.setProto3StringField(this, 17, value);
};


/**
 * optional core.solo.io.ResourceRef rootCa = 18;
 * @return {?proto.core.solo.io.ResourceRef}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.getRootca = function() {
  return /** @type{?proto.core.solo.io.ResourceRef} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef, 18));
};


/** @param {?proto.core.solo.io.ResourceRef|undefined} value */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.setRootca = function(value) {
  jspb.Message.setWrapperField(this, 18, value);
};


proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.clearRootca = function() {
  this.setRootca(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.hasRootca = function() {
  return jspb.Message.getField(this, 18) != null;
};


/**
 * optional bool splitTlsServices = 19;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.getSplittlsservices = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 19, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.setSplittlsservices = function(value) {
  jspb.Message.setProto3BooleanField(this, 19, value);
};


/**
 * optional consul.options.gloo.solo.io.ConsulConsistencyModes consistencyMode = 20;
 * @return {!proto.consul.options.gloo.solo.io.ConsulConsistencyModes}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.getConsistencymode = function() {
  return /** @type {!proto.consul.options.gloo.solo.io.ConsulConsistencyModes} */ (jspb.Message.getFieldWithDefault(this, 20, 0));
};


/** @param {!proto.consul.options.gloo.solo.io.ConsulConsistencyModes} value */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.setConsistencymode = function(value) {
  jspb.Message.setProto3EnumField(this, 20, value);
};


/**
 * optional consul.options.gloo.solo.io.QueryOptions query_options = 21;
 * @return {?proto.consul.options.gloo.solo.io.QueryOptions}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.getQueryOptions = function() {
  return /** @type{?proto.consul.options.gloo.solo.io.QueryOptions} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_consul_query_options_pb.QueryOptions, 21));
};


/** @param {?proto.consul.options.gloo.solo.io.QueryOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.setQueryOptions = function(value) {
  jspb.Message.setWrapperField(this, 21, value);
};


proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.clearQueryOptions = function() {
  this.setQueryOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.hasQueryOptions = function() {
  return jspb.Message.getField(this, 21) != null;
};


/**
 * repeated string service_tags_allowlist = 22;
 * @return {!Array<string>}
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.getServiceTagsAllowlistList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 22));
};


/** @param {!Array<string>} value */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.setServiceTagsAllowlistList = function(value) {
  jspb.Message.setField(this, 22, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.addServiceTagsAllowlist = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 22, value, opt_index);
};


proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration.prototype.clearServiceTagsAllowlistList = function() {
  this.setServiceTagsAllowlistList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.displayName = 'proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.toObject = function(includeInstance, msg) {
  var f, obj = {
    rateLimits: (f = msg.getRateLimits()) && proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration;
  return proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.deserializeBinaryFromReader);
      msg.setRateLimits(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRateLimits();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.serializeBinaryToWriter
    );
  }
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.displayName = 'proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.toObject = function(includeInstance, msg) {
  var f, obj = {
    qps: +jspb.Message.getFieldWithDefault(msg, 1, 0.0),
    burst: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits;
  return proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setQps(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setBurst(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getQps();
  if (f !== 0.0) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = message.getBurst();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
};


/**
 * optional float QPS = 1;
 * @return {number}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.prototype.getQps = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 1, 0.0));
};


/** @param {number} value */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.prototype.setQps = function(value) {
  jspb.Message.setProto3FloatField(this, 1, value);
};


/**
 * optional uint32 burst = 2;
 * @return {number}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.prototype.getBurst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits.prototype.setBurst = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional RateLimits rate_limits = 1;
 * @return {?proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.prototype.getRateLimits = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits, 1));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.RateLimits|undefined} value */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.prototype.setRateLimits = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.prototype.clearRateLimits = function() {
  this.setRateLimits(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration.prototype.hasRateLimits = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.ObservabilityOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.displayName = 'proto.gloo.solo.io.SettingsSpec.ObservabilityOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    grafanaintegration: (f = msg.getGrafanaintegration()) && proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.toObject(includeInstance, f),
    configstatusmetriclabelsMap: (f = msg.getConfigstatusmetriclabelsMap()) ? f.toObject(includeInstance, proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.toObject) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.ObservabilityOptions;
  return proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration;
      reader.readMessage(value,proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.deserializeBinaryFromReader);
      msg.setGrafanaintegration(value);
      break;
    case 2:
      var value = msg.getConfigstatusmetriclabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.deserializeBinaryFromReader, "");
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGrafanaintegration();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.serializeBinaryToWriter
    );
  }
  f = message.getConfigstatusmetriclabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(2, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.serializeBinaryToWriter);
  }
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.displayName = 'proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.toObject = function(includeInstance, msg) {
  var f, obj = {
    defaultDashboardFolderId: (f = msg.getDefaultDashboardFolderId()) && google_protobuf_wrappers_pb.UInt32Value.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration;
  return proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.UInt32Value;
      reader.readMessage(value,google_protobuf_wrappers_pb.UInt32Value.deserializeBinaryFromReader);
      msg.setDefaultDashboardFolderId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDefaultDashboardFolderId();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.UInt32Value.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.UInt32Value default_dashboard_folder_id = 1;
 * @return {?proto.google.protobuf.UInt32Value}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.prototype.getDefaultDashboardFolderId = function() {
  return /** @type{?proto.google.protobuf.UInt32Value} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.UInt32Value, 1));
};


/** @param {?proto.google.protobuf.UInt32Value|undefined} value */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.prototype.setDefaultDashboardFolderId = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.prototype.clearDefaultDashboardFolderId = function() {
  this.setDefaultDashboardFolderId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration.prototype.hasDefaultDashboardFolderId = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.displayName = 'proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.toObject = function(includeInstance, msg) {
  var f, obj = {
    labeltopathMap: (f = msg.getLabeltopathMap()) ? f.toObject(includeInstance, undefined) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels;
  return proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = msg.getLabeltopathMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLabeltopathMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(1, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * map<string, string> labelToPath = 1;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.prototype.getLabeltopathMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 1, opt_noLazyCreate,
      null));
};


proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels.prototype.clearLabeltopathMap = function() {
  this.getLabeltopathMap().clear();
};


/**
 * optional GrafanaIntegration grafanaIntegration = 1;
 * @return {?proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.prototype.getGrafanaintegration = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration, 1));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.GrafanaIntegration|undefined} value */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.prototype.setGrafanaintegration = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.prototype.clearGrafanaintegration = function() {
  this.setGrafanaintegration(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.prototype.hasGrafanaintegration = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * map<string, MetricLabels> configStatusMetricLabels = 2;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels>}
 */
proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.prototype.getConfigstatusmetriclabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels>} */ (
      jspb.Message.getMapField(this, 2, opt_noLazyCreate,
      proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.MetricLabels));
};


proto.gloo.solo.io.SettingsSpec.ObservabilityOptions.prototype.clearConfigstatusmetriclabelsMap = function() {
  this.getConfigstatusmetriclabelsMap().clear();
};


/**
 * optional string discovery_namespace = 1;
 * @return {string}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getDiscoveryNamespace = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsSpec.prototype.setDiscoveryNamespace = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string watch_namespaces = 2;
 * @return {!Array<string>}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getWatchNamespacesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/** @param {!Array<string>} value */
proto.gloo.solo.io.SettingsSpec.prototype.setWatchNamespacesList = function(value) {
  jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.gloo.solo.io.SettingsSpec.prototype.addWatchNamespaces = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearWatchNamespacesList = function() {
  this.setWatchNamespacesList([]);
};


/**
 * optional KubernetesCrds kubernetes_config_source = 4;
 * @return {?proto.gloo.solo.io.SettingsSpec.KubernetesCrds}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getKubernetesConfigSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.KubernetesCrds} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.KubernetesCrds, 4));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.KubernetesCrds|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setKubernetesConfigSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 4, proto.gloo.solo.io.SettingsSpec.oneofGroups_[0], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearKubernetesConfigSource = function() {
  this.setKubernetesConfigSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasKubernetesConfigSource = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional Directory directory_config_source = 5;
 * @return {?proto.gloo.solo.io.SettingsSpec.Directory}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getDirectoryConfigSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.Directory} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.Directory, 5));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.Directory|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setDirectoryConfigSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 5, proto.gloo.solo.io.SettingsSpec.oneofGroups_[0], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearDirectoryConfigSource = function() {
  this.setDirectoryConfigSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasDirectoryConfigSource = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional ConsulKv consul_kv_source = 21;
 * @return {?proto.gloo.solo.io.SettingsSpec.ConsulKv}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getConsulKvSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.ConsulKv} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.ConsulKv, 21));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.ConsulKv|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setConsulKvSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 21, proto.gloo.solo.io.SettingsSpec.oneofGroups_[0], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearConsulKvSource = function() {
  this.setConsulKvSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasConsulKvSource = function() {
  return jspb.Message.getField(this, 21) != null;
};


/**
 * optional KubernetesSecrets kubernetes_secret_source = 6;
 * @return {?proto.gloo.solo.io.SettingsSpec.KubernetesSecrets}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getKubernetesSecretSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.KubernetesSecrets} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.KubernetesSecrets, 6));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.KubernetesSecrets|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setKubernetesSecretSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 6, proto.gloo.solo.io.SettingsSpec.oneofGroups_[1], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearKubernetesSecretSource = function() {
  this.setKubernetesSecretSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasKubernetesSecretSource = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional VaultSecrets vault_secret_source = 7;
 * @return {?proto.gloo.solo.io.SettingsSpec.VaultSecrets}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getVaultSecretSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.VaultSecrets} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.VaultSecrets, 7));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.VaultSecrets|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setVaultSecretSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 7, proto.gloo.solo.io.SettingsSpec.oneofGroups_[1], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearVaultSecretSource = function() {
  this.setVaultSecretSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasVaultSecretSource = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional Directory directory_secret_source = 8;
 * @return {?proto.gloo.solo.io.SettingsSpec.Directory}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getDirectorySecretSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.Directory} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.Directory, 8));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.Directory|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setDirectorySecretSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 8, proto.gloo.solo.io.SettingsSpec.oneofGroups_[1], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearDirectorySecretSource = function() {
  this.setDirectorySecretSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasDirectorySecretSource = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional KubernetesConfigmaps kubernetes_artifact_source = 9;
 * @return {?proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getKubernetesArtifactSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps, 9));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.KubernetesConfigmaps|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setKubernetesArtifactSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 9, proto.gloo.solo.io.SettingsSpec.oneofGroups_[2], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearKubernetesArtifactSource = function() {
  this.setKubernetesArtifactSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasKubernetesArtifactSource = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional Directory directory_artifact_source = 10;
 * @return {?proto.gloo.solo.io.SettingsSpec.Directory}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getDirectoryArtifactSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.Directory} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.Directory, 10));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.Directory|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setDirectoryArtifactSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 10, proto.gloo.solo.io.SettingsSpec.oneofGroups_[2], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearDirectoryArtifactSource = function() {
  this.setDirectoryArtifactSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasDirectoryArtifactSource = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional ConsulKv consul_kv_artifact_source = 23;
 * @return {?proto.gloo.solo.io.SettingsSpec.ConsulKv}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getConsulKvArtifactSource = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.ConsulKv} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.ConsulKv, 23));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.ConsulKv|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setConsulKvArtifactSource = function(value) {
  jspb.Message.setOneofWrapperField(this, 23, proto.gloo.solo.io.SettingsSpec.oneofGroups_[2], value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearConsulKvArtifactSource = function() {
  this.setConsulKvArtifactSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasConsulKvArtifactSource = function() {
  return jspb.Message.getField(this, 23) != null;
};


/**
 * optional google.protobuf.Duration refresh_rate = 12;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getRefreshRate = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 12));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setRefreshRate = function(value) {
  jspb.Message.setWrapperField(this, 12, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearRefreshRate = function() {
  this.setRefreshRate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasRefreshRate = function() {
  return jspb.Message.getField(this, 12) != null;
};


/**
 * optional bool dev_mode = 13;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getDevMode = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 13, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.SettingsSpec.prototype.setDevMode = function(value) {
  jspb.Message.setProto3BooleanField(this, 13, value);
};


/**
 * optional bool linkerd = 17;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getLinkerd = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 17, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.SettingsSpec.prototype.setLinkerd = function(value) {
  jspb.Message.setProto3BooleanField(this, 17, value);
};


/**
 * optional KnativeOptions knative = 18;
 * @return {?proto.gloo.solo.io.SettingsSpec.KnativeOptions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getKnative = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.KnativeOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.KnativeOptions, 18));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.KnativeOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setKnative = function(value) {
  jspb.Message.setWrapperField(this, 18, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearKnative = function() {
  this.setKnative(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasKnative = function() {
  return jspb.Message.getField(this, 18) != null;
};


/**
 * optional DiscoveryOptions discovery = 19;
 * @return {?proto.gloo.solo.io.SettingsSpec.DiscoveryOptions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getDiscovery = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.DiscoveryOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.DiscoveryOptions, 19));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.DiscoveryOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setDiscovery = function(value) {
  jspb.Message.setWrapperField(this, 19, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearDiscovery = function() {
  this.setDiscovery(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasDiscovery = function() {
  return jspb.Message.getField(this, 19) != null;
};


/**
 * optional GlooOptions gloo = 24;
 * @return {?proto.gloo.solo.io.GlooOptions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getGloo = function() {
  return /** @type{?proto.gloo.solo.io.GlooOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.GlooOptions, 24));
};


/** @param {?proto.gloo.solo.io.GlooOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setGloo = function(value) {
  jspb.Message.setWrapperField(this, 24, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearGloo = function() {
  this.setGloo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasGloo = function() {
  return jspb.Message.getField(this, 24) != null;
};


/**
 * optional GatewayOptions gateway = 25;
 * @return {?proto.gloo.solo.io.GatewayOptions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getGateway = function() {
  return /** @type{?proto.gloo.solo.io.GatewayOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.GatewayOptions, 25));
};


/** @param {?proto.gloo.solo.io.GatewayOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setGateway = function(value) {
  jspb.Message.setWrapperField(this, 25, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearGateway = function() {
  this.setGateway(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasGateway = function() {
  return jspb.Message.getField(this, 25) != null;
};


/**
 * optional ConsulConfiguration consul = 20;
 * @return {?proto.gloo.solo.io.SettingsSpec.ConsulConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getConsul = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.ConsulConfiguration} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.ConsulConfiguration, 20));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.ConsulConfiguration|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setConsul = function(value) {
  jspb.Message.setWrapperField(this, 20, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearConsul = function() {
  this.setConsul(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasConsul = function() {
  return jspb.Message.getField(this, 20) != null;
};


/**
 * optional ConsulUpstreamDiscoveryConfiguration consulDiscovery = 30;
 * @return {?proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getConsuldiscovery = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration, 30));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.ConsulUpstreamDiscoveryConfiguration|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setConsuldiscovery = function(value) {
  jspb.Message.setWrapperField(this, 30, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearConsuldiscovery = function() {
  this.setConsuldiscovery(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasConsuldiscovery = function() {
  return jspb.Message.getField(this, 30) != null;
};


/**
 * optional KubernetesConfiguration kubernetes = 22;
 * @return {?proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getKubernetes = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration, 22));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.KubernetesConfiguration|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setKubernetes = function(value) {
  jspb.Message.setWrapperField(this, 22, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearKubernetes = function() {
  this.setKubernetes(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasKubernetes = function() {
  return jspb.Message.getField(this, 22) != null;
};


/**
 * optional Extensions extensions = 16;
 * @return {?proto.gloo.solo.io.Extensions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getExtensions = function() {
  return /** @type{?proto.gloo.solo.io.Extensions} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_extensions_pb.Extensions, 16));
};


/** @param {?proto.gloo.solo.io.Extensions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setExtensions = function(value) {
  jspb.Message.setWrapperField(this, 16, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearExtensions = function() {
  this.setExtensions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasExtensions = function() {
  return jspb.Message.getField(this, 16) != null;
};


/**
 * optional ratelimit.options.gloo.solo.io.ServiceSettings ratelimit = 26;
 * @return {?proto.ratelimit.options.gloo.solo.io.ServiceSettings}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getRatelimit = function() {
  return /** @type{?proto.ratelimit.options.gloo.solo.io.ServiceSettings} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.ServiceSettings, 26));
};


/** @param {?proto.ratelimit.options.gloo.solo.io.ServiceSettings|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setRatelimit = function(value) {
  jspb.Message.setWrapperField(this, 26, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearRatelimit = function() {
  this.setRatelimit(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasRatelimit = function() {
  return jspb.Message.getField(this, 26) != null;
};


/**
 * optional ratelimit.options.gloo.solo.io.Settings ratelimit_server = 27;
 * @return {?proto.ratelimit.options.gloo.solo.io.Settings}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getRatelimitServer = function() {
  return /** @type{?proto.ratelimit.options.gloo.solo.io.Settings} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_ratelimit_ratelimit_pb.Settings, 27));
};


/** @param {?proto.ratelimit.options.gloo.solo.io.Settings|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setRatelimitServer = function(value) {
  jspb.Message.setWrapperField(this, 27, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearRatelimitServer = function() {
  this.setRatelimitServer(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasRatelimitServer = function() {
  return jspb.Message.getField(this, 27) != null;
};


/**
 * optional rbac.options.gloo.solo.io.Settings rbac = 28;
 * @return {?proto.rbac.options.gloo.solo.io.Settings}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getRbac = function() {
  return /** @type{?proto.rbac.options.gloo.solo.io.Settings} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_rbac_rbac_pb.Settings, 28));
};


/** @param {?proto.rbac.options.gloo.solo.io.Settings|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setRbac = function(value) {
  jspb.Message.setWrapperField(this, 28, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearRbac = function() {
  this.setRbac(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasRbac = function() {
  return jspb.Message.getField(this, 28) != null;
};


/**
 * optional enterprise.gloo.solo.io.Settings extauth = 29;
 * @return {?proto.enterprise.gloo.solo.io.Settings}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getExtauth = function() {
  return /** @type{?proto.enterprise.gloo.solo.io.Settings} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.Settings, 29));
};


/** @param {?proto.enterprise.gloo.solo.io.Settings|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setExtauth = function(value) {
  jspb.Message.setWrapperField(this, 29, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearExtauth = function() {
  this.setExtauth(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasExtauth = function() {
  return jspb.Message.getField(this, 29) != null;
};


/**
 * map<string, enterprise.gloo.solo.io.Settings> named_extauth = 33;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.enterprise.gloo.solo.io.Settings>}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getNamedExtauthMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.enterprise.gloo.solo.io.Settings>} */ (
      jspb.Message.getMapField(this, 33, opt_noLazyCreate,
      proto.enterprise.gloo.solo.io.Settings));
};


proto.gloo.solo.io.SettingsSpec.prototype.clearNamedExtauthMap = function() {
  this.getNamedExtauthMap().clear();
};


/**
 * optional caching.options.gloo.solo.io.Settings caching_server = 36;
 * @return {?proto.caching.options.gloo.solo.io.Settings}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getCachingServer = function() {
  return /** @type{?proto.caching.options.gloo.solo.io.Settings} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_enterprise_options_caching_caching_pb.Settings, 36));
};


/** @param {?proto.caching.options.gloo.solo.io.Settings|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setCachingServer = function(value) {
  jspb.Message.setWrapperField(this, 36, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearCachingServer = function() {
  this.setCachingServer(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasCachingServer = function() {
  return jspb.Message.getField(this, 36) != null;
};


/**
 * optional ObservabilityOptions observabilityOptions = 31;
 * @return {?proto.gloo.solo.io.SettingsSpec.ObservabilityOptions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getObservabilityoptions = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec.ObservabilityOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.SettingsSpec.ObservabilityOptions, 31));
};


/** @param {?proto.gloo.solo.io.SettingsSpec.ObservabilityOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setObservabilityoptions = function(value) {
  jspb.Message.setWrapperField(this, 31, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearObservabilityoptions = function() {
  this.setObservabilityoptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasObservabilityoptions = function() {
  return jspb.Message.getField(this, 31) != null;
};


/**
 * optional UpstreamOptions upstreamOptions = 32;
 * @return {?proto.gloo.solo.io.UpstreamOptions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getUpstreamoptions = function() {
  return /** @type{?proto.gloo.solo.io.UpstreamOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.UpstreamOptions, 32));
};


/** @param {?proto.gloo.solo.io.UpstreamOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setUpstreamoptions = function(value) {
  jspb.Message.setWrapperField(this, 32, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearUpstreamoptions = function() {
  this.setUpstreamoptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasUpstreamoptions = function() {
  return jspb.Message.getField(this, 32) != null;
};


/**
 * optional ConsoleOptions console_options = 35;
 * @return {?proto.gloo.solo.io.ConsoleOptions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getConsoleOptions = function() {
  return /** @type{?proto.gloo.solo.io.ConsoleOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.ConsoleOptions, 35));
};


/** @param {?proto.gloo.solo.io.ConsoleOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setConsoleOptions = function(value) {
  jspb.Message.setWrapperField(this, 35, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearConsoleOptions = function() {
  this.setConsoleOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasConsoleOptions = function() {
  return jspb.Message.getField(this, 35) != null;
};


/**
 * optional GraphqlOptions graphql_options = 37;
 * @return {?proto.gloo.solo.io.GraphqlOptions}
 */
proto.gloo.solo.io.SettingsSpec.prototype.getGraphqlOptions = function() {
  return /** @type{?proto.gloo.solo.io.GraphqlOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.GraphqlOptions, 37));
};


/** @param {?proto.gloo.solo.io.GraphqlOptions|undefined} value */
proto.gloo.solo.io.SettingsSpec.prototype.setGraphqlOptions = function(value) {
  jspb.Message.setWrapperField(this, 37, value);
};


proto.gloo.solo.io.SettingsSpec.prototype.clearGraphqlOptions = function() {
  this.setGraphqlOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsSpec.prototype.hasGraphqlOptions = function() {
  return jspb.Message.getField(this, 37) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.UpstreamOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.UpstreamOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.UpstreamOptions.displayName = 'proto.gloo.solo.io.UpstreamOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.UpstreamOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.UpstreamOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.UpstreamOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.UpstreamOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    sslParameters: (f = msg.getSslParameters()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_ssl_pb.SslParameters.toObject(includeInstance, f),
    globalAnnotationsMap: (f = msg.getGlobalAnnotationsMap()) ? f.toObject(includeInstance, undefined) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.UpstreamOptions}
 */
proto.gloo.solo.io.UpstreamOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.UpstreamOptions;
  return proto.gloo.solo.io.UpstreamOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.UpstreamOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.UpstreamOptions}
 */
proto.gloo.solo.io.UpstreamOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_ssl_pb.SslParameters;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_ssl_pb.SslParameters.deserializeBinaryFromReader);
      msg.setSslParameters(value);
      break;
    case 2:
      var value = msg.getGlobalAnnotationsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.UpstreamOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.UpstreamOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.UpstreamOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.UpstreamOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSslParameters();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_ssl_pb.SslParameters.serializeBinaryToWriter
    );
  }
  f = message.getGlobalAnnotationsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(2, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * optional SslParameters ssl_parameters = 1;
 * @return {?proto.gloo.solo.io.SslParameters}
 */
proto.gloo.solo.io.UpstreamOptions.prototype.getSslParameters = function() {
  return /** @type{?proto.gloo.solo.io.SslParameters} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_ssl_pb.SslParameters, 1));
};


/** @param {?proto.gloo.solo.io.SslParameters|undefined} value */
proto.gloo.solo.io.UpstreamOptions.prototype.setSslParameters = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.UpstreamOptions.prototype.clearSslParameters = function() {
  this.setSslParameters(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.UpstreamOptions.prototype.hasSslParameters = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * map<string, string> global_annotations = 2;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.gloo.solo.io.UpstreamOptions.prototype.getGlobalAnnotationsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 2, opt_noLazyCreate,
      null));
};


proto.gloo.solo.io.UpstreamOptions.prototype.clearGlobalAnnotationsMap = function() {
  this.getGlobalAnnotationsMap().clear();
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.GlooOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.GlooOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.GlooOptions.displayName = 'proto.gloo.solo.io.GlooOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.GlooOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.GlooOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.GlooOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GlooOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    xdsBindAddr: jspb.Message.getFieldWithDefault(msg, 1, ""),
    validationBindAddr: jspb.Message.getFieldWithDefault(msg, 2, ""),
    circuitBreakers: (f = msg.getCircuitBreakers()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_circuit_breaker_pb.CircuitBreakerConfig.toObject(includeInstance, f),
    endpointsWarmingTimeout: (f = msg.getEndpointsWarmingTimeout()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    awsOptions: (f = msg.getAwsOptions()) && proto.gloo.solo.io.GlooOptions.AWSOptions.toObject(includeInstance, f),
    invalidConfigPolicy: (f = msg.getInvalidConfigPolicy()) && proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.toObject(includeInstance, f),
    disableKubernetesDestinations: jspb.Message.getFieldWithDefault(msg, 7, false),
    disableGrpcWeb: (f = msg.getDisableGrpcWeb()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    disableProxyGarbageCollection: (f = msg.getDisableProxyGarbageCollection()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    regexMaxProgramSize: (f = msg.getRegexMaxProgramSize()) && google_protobuf_wrappers_pb.UInt32Value.toObject(includeInstance, f),
    restXdsBindAddr: jspb.Message.getFieldWithDefault(msg, 11, ""),
    enableRestEds: (f = msg.getEnableRestEds()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    failoverUpstreamDnsPollingInterval: (f = msg.getFailoverUpstreamDnsPollingInterval()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    removeUnusedFilters: (f = msg.getRemoveUnusedFilters()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    proxyDebugBindAddr: jspb.Message.getFieldWithDefault(msg, 15, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.GlooOptions}
 */
proto.gloo.solo.io.GlooOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.GlooOptions;
  return proto.gloo.solo.io.GlooOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.GlooOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.GlooOptions}
 */
proto.gloo.solo.io.GlooOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setXdsBindAddr(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setValidationBindAddr(value);
      break;
    case 3:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_circuit_breaker_pb.CircuitBreakerConfig;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_circuit_breaker_pb.CircuitBreakerConfig.deserializeBinaryFromReader);
      msg.setCircuitBreakers(value);
      break;
    case 4:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setEndpointsWarmingTimeout(value);
      break;
    case 5:
      var value = new proto.gloo.solo.io.GlooOptions.AWSOptions;
      reader.readMessage(value,proto.gloo.solo.io.GlooOptions.AWSOptions.deserializeBinaryFromReader);
      msg.setAwsOptions(value);
      break;
    case 6:
      var value = new proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy;
      reader.readMessage(value,proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.deserializeBinaryFromReader);
      msg.setInvalidConfigPolicy(value);
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDisableKubernetesDestinations(value);
      break;
    case 8:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setDisableGrpcWeb(value);
      break;
    case 9:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setDisableProxyGarbageCollection(value);
      break;
    case 10:
      var value = new google_protobuf_wrappers_pb.UInt32Value;
      reader.readMessage(value,google_protobuf_wrappers_pb.UInt32Value.deserializeBinaryFromReader);
      msg.setRegexMaxProgramSize(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setRestXdsBindAddr(value);
      break;
    case 12:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setEnableRestEds(value);
      break;
    case 13:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setFailoverUpstreamDnsPollingInterval(value);
      break;
    case 14:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setRemoveUnusedFilters(value);
      break;
    case 15:
      var value = /** @type {string} */ (reader.readString());
      msg.setProxyDebugBindAddr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.GlooOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.GlooOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.GlooOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GlooOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getXdsBindAddr();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getValidationBindAddr();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getCircuitBreakers();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_circuit_breaker_pb.CircuitBreakerConfig.serializeBinaryToWriter
    );
  }
  f = message.getEndpointsWarmingTimeout();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getAwsOptions();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.gloo.solo.io.GlooOptions.AWSOptions.serializeBinaryToWriter
    );
  }
  f = message.getInvalidConfigPolicy();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.serializeBinaryToWriter
    );
  }
  f = message.getDisableKubernetesDestinations();
  if (f) {
    writer.writeBool(
      7,
      f
    );
  }
  f = message.getDisableGrpcWeb();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getDisableProxyGarbageCollection();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getRegexMaxProgramSize();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      google_protobuf_wrappers_pb.UInt32Value.serializeBinaryToWriter
    );
  }
  f = message.getRestXdsBindAddr();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getEnableRestEds();
  if (f != null) {
    writer.writeMessage(
      12,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getFailoverUpstreamDnsPollingInterval();
  if (f != null) {
    writer.writeMessage(
      13,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getRemoveUnusedFilters();
  if (f != null) {
    writer.writeMessage(
      14,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getProxyDebugBindAddr();
  if (f.length > 0) {
    writer.writeString(
      15,
      f
    );
  }
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.GlooOptions.AWSOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.gloo.solo.io.GlooOptions.AWSOptions.oneofGroups_);
};
goog.inherits(proto.gloo.solo.io.GlooOptions.AWSOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.GlooOptions.AWSOptions.displayName = 'proto.gloo.solo.io.GlooOptions.AWSOptions';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.CredentialsFetcherCase = {
  CREDENTIALS_FETCHER_NOT_SET: 0,
  ENABLE_CREDENTIALS_DISCOVEY: 1,
  SERVICE_ACCOUNT_CREDENTIALS: 2
};

/**
 * @return {proto.gloo.solo.io.GlooOptions.AWSOptions.CredentialsFetcherCase}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.getCredentialsFetcherCase = function() {
  return /** @type {proto.gloo.solo.io.GlooOptions.AWSOptions.CredentialsFetcherCase} */(jspb.Message.computeOneofCase(this, proto.gloo.solo.io.GlooOptions.AWSOptions.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.GlooOptions.AWSOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.GlooOptions.AWSOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    enableCredentialsDiscovey: jspb.Message.getFieldWithDefault(msg, 1, false),
    serviceAccountCredentials: (f = msg.getServiceAccountCredentials()) && github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_extensions_aws_filter_pb.AWSLambdaConfig.ServiceAccountCredentials.toObject(includeInstance, f),
    propagateOriginalRouting: (f = msg.getPropagateOriginalRouting()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    credentialRefreshDelay: (f = msg.getCredentialRefreshDelay()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.GlooOptions.AWSOptions}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.GlooOptions.AWSOptions;
  return proto.gloo.solo.io.GlooOptions.AWSOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.GlooOptions.AWSOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.GlooOptions.AWSOptions}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setEnableCredentialsDiscovey(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_extensions_aws_filter_pb.AWSLambdaConfig.ServiceAccountCredentials;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_extensions_aws_filter_pb.AWSLambdaConfig.ServiceAccountCredentials.deserializeBinaryFromReader);
      msg.setServiceAccountCredentials(value);
      break;
    case 3:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setPropagateOriginalRouting(value);
      break;
    case 4:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setCredentialRefreshDelay(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.GlooOptions.AWSOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.GlooOptions.AWSOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {boolean} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeBool(
      1,
      f
    );
  }
  f = message.getServiceAccountCredentials();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_extensions_aws_filter_pb.AWSLambdaConfig.ServiceAccountCredentials.serializeBinaryToWriter
    );
  }
  f = message.getPropagateOriginalRouting();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getCredentialRefreshDelay();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
};


/**
 * optional bool enable_credentials_discovey = 1;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.getEnableCredentialsDiscovey = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 1, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.setEnableCredentialsDiscovey = function(value) {
  jspb.Message.setOneofField(this, 1, proto.gloo.solo.io.GlooOptions.AWSOptions.oneofGroups_[0], value);
};


proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.clearEnableCredentialsDiscovey = function() {
  jspb.Message.setOneofField(this, 1, proto.gloo.solo.io.GlooOptions.AWSOptions.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.hasEnableCredentialsDiscovey = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.ServiceAccountCredentials service_account_credentials = 2;
 * @return {?proto.envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.ServiceAccountCredentials}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.getServiceAccountCredentials = function() {
  return /** @type{?proto.envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.ServiceAccountCredentials} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_extensions_aws_filter_pb.AWSLambdaConfig.ServiceAccountCredentials, 2));
};


/** @param {?proto.envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig.ServiceAccountCredentials|undefined} value */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.setServiceAccountCredentials = function(value) {
  jspb.Message.setOneofWrapperField(this, 2, proto.gloo.solo.io.GlooOptions.AWSOptions.oneofGroups_[0], value);
};


proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.clearServiceAccountCredentials = function() {
  this.setServiceAccountCredentials(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.hasServiceAccountCredentials = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.BoolValue propagate_original_routing = 3;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.getPropagateOriginalRouting = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 3));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.setPropagateOriginalRouting = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.clearPropagateOriginalRouting = function() {
  this.setPropagateOriginalRouting(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.hasPropagateOriginalRouting = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.Duration credential_refresh_delay = 4;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.getCredentialRefreshDelay = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 4));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.setCredentialRefreshDelay = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.clearCredentialRefreshDelay = function() {
  this.setCredentialRefreshDelay(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.AWSOptions.prototype.hasCredentialRefreshDelay = function() {
  return jspb.Message.getField(this, 4) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.displayName = 'proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.toObject = function(includeInstance, msg) {
  var f, obj = {
    replaceInvalidRoutes: jspb.Message.getFieldWithDefault(msg, 1, false),
    invalidRouteResponseCode: jspb.Message.getFieldWithDefault(msg, 2, 0),
    invalidRouteResponseBody: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy}
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy;
  return proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy}
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setReplaceInvalidRoutes(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setInvalidRouteResponseCode(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setInvalidRouteResponseBody(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getReplaceInvalidRoutes();
  if (f) {
    writer.writeBool(
      1,
      f
    );
  }
  f = message.getInvalidRouteResponseCode();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
  f = message.getInvalidRouteResponseBody();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional bool replace_invalid_routes = 1;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.prototype.getReplaceInvalidRoutes = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 1, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.prototype.setReplaceInvalidRoutes = function(value) {
  jspb.Message.setProto3BooleanField(this, 1, value);
};


/**
 * optional uint32 invalid_route_response_code = 2;
 * @return {number}
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.prototype.getInvalidRouteResponseCode = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.prototype.setInvalidRouteResponseCode = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional string invalid_route_response_body = 3;
 * @return {string}
 */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.prototype.getInvalidRouteResponseBody = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy.prototype.setInvalidRouteResponseBody = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string xds_bind_addr = 1;
 * @return {string}
 */
proto.gloo.solo.io.GlooOptions.prototype.getXdsBindAddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GlooOptions.prototype.setXdsBindAddr = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string validation_bind_addr = 2;
 * @return {string}
 */
proto.gloo.solo.io.GlooOptions.prototype.getValidationBindAddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GlooOptions.prototype.setValidationBindAddr = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional CircuitBreakerConfig circuit_breakers = 3;
 * @return {?proto.gloo.solo.io.CircuitBreakerConfig}
 */
proto.gloo.solo.io.GlooOptions.prototype.getCircuitBreakers = function() {
  return /** @type{?proto.gloo.solo.io.CircuitBreakerConfig} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_circuit_breaker_pb.CircuitBreakerConfig, 3));
};


/** @param {?proto.gloo.solo.io.CircuitBreakerConfig|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setCircuitBreakers = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearCircuitBreakers = function() {
  this.setCircuitBreakers(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasCircuitBreakers = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.Duration endpoints_warming_timeout = 4;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.GlooOptions.prototype.getEndpointsWarmingTimeout = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 4));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setEndpointsWarmingTimeout = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearEndpointsWarmingTimeout = function() {
  this.setEndpointsWarmingTimeout(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasEndpointsWarmingTimeout = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional AWSOptions aws_options = 5;
 * @return {?proto.gloo.solo.io.GlooOptions.AWSOptions}
 */
proto.gloo.solo.io.GlooOptions.prototype.getAwsOptions = function() {
  return /** @type{?proto.gloo.solo.io.GlooOptions.AWSOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.GlooOptions.AWSOptions, 5));
};


/** @param {?proto.gloo.solo.io.GlooOptions.AWSOptions|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setAwsOptions = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearAwsOptions = function() {
  this.setAwsOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasAwsOptions = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional InvalidConfigPolicy invalid_config_policy = 6;
 * @return {?proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy}
 */
proto.gloo.solo.io.GlooOptions.prototype.getInvalidConfigPolicy = function() {
  return /** @type{?proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy, 6));
};


/** @param {?proto.gloo.solo.io.GlooOptions.InvalidConfigPolicy|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setInvalidConfigPolicy = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearInvalidConfigPolicy = function() {
  this.setInvalidConfigPolicy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasInvalidConfigPolicy = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional bool disable_kubernetes_destinations = 7;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.getDisableKubernetesDestinations = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 7, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.GlooOptions.prototype.setDisableKubernetesDestinations = function(value) {
  jspb.Message.setProto3BooleanField(this, 7, value);
};


/**
 * optional google.protobuf.BoolValue disable_grpc_web = 8;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GlooOptions.prototype.getDisableGrpcWeb = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 8));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setDisableGrpcWeb = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearDisableGrpcWeb = function() {
  this.setDisableGrpcWeb(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasDisableGrpcWeb = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional google.protobuf.BoolValue disable_proxy_garbage_collection = 9;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GlooOptions.prototype.getDisableProxyGarbageCollection = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 9));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setDisableProxyGarbageCollection = function(value) {
  jspb.Message.setWrapperField(this, 9, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearDisableProxyGarbageCollection = function() {
  this.setDisableProxyGarbageCollection(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasDisableProxyGarbageCollection = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional google.protobuf.UInt32Value regex_max_program_size = 10;
 * @return {?proto.google.protobuf.UInt32Value}
 */
proto.gloo.solo.io.GlooOptions.prototype.getRegexMaxProgramSize = function() {
  return /** @type{?proto.google.protobuf.UInt32Value} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.UInt32Value, 10));
};


/** @param {?proto.google.protobuf.UInt32Value|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setRegexMaxProgramSize = function(value) {
  jspb.Message.setWrapperField(this, 10, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearRegexMaxProgramSize = function() {
  this.setRegexMaxProgramSize(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasRegexMaxProgramSize = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional string rest_xds_bind_addr = 11;
 * @return {string}
 */
proto.gloo.solo.io.GlooOptions.prototype.getRestXdsBindAddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GlooOptions.prototype.setRestXdsBindAddr = function(value) {
  jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional google.protobuf.BoolValue enable_rest_eds = 12;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GlooOptions.prototype.getEnableRestEds = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 12));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setEnableRestEds = function(value) {
  jspb.Message.setWrapperField(this, 12, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearEnableRestEds = function() {
  this.setEnableRestEds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasEnableRestEds = function() {
  return jspb.Message.getField(this, 12) != null;
};


/**
 * optional google.protobuf.Duration failover_upstream_dns_polling_interval = 13;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.GlooOptions.prototype.getFailoverUpstreamDnsPollingInterval = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 13));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setFailoverUpstreamDnsPollingInterval = function(value) {
  jspb.Message.setWrapperField(this, 13, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearFailoverUpstreamDnsPollingInterval = function() {
  this.setFailoverUpstreamDnsPollingInterval(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasFailoverUpstreamDnsPollingInterval = function() {
  return jspb.Message.getField(this, 13) != null;
};


/**
 * optional google.protobuf.BoolValue remove_unused_filters = 14;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GlooOptions.prototype.getRemoveUnusedFilters = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 14));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GlooOptions.prototype.setRemoveUnusedFilters = function(value) {
  jspb.Message.setWrapperField(this, 14, value);
};


proto.gloo.solo.io.GlooOptions.prototype.clearRemoveUnusedFilters = function() {
  this.setRemoveUnusedFilters(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GlooOptions.prototype.hasRemoveUnusedFilters = function() {
  return jspb.Message.getField(this, 14) != null;
};


/**
 * optional string proxy_debug_bind_addr = 15;
 * @return {string}
 */
proto.gloo.solo.io.GlooOptions.prototype.getProxyDebugBindAddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 15, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GlooOptions.prototype.setProxyDebugBindAddr = function(value) {
  jspb.Message.setProto3StringField(this, 15, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.VirtualServiceOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.VirtualServiceOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.VirtualServiceOptions.displayName = 'proto.gloo.solo.io.VirtualServiceOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.VirtualServiceOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.VirtualServiceOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.VirtualServiceOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.VirtualServiceOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    oneWayTls: (f = msg.getOneWayTls()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.VirtualServiceOptions}
 */
proto.gloo.solo.io.VirtualServiceOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.VirtualServiceOptions;
  return proto.gloo.solo.io.VirtualServiceOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.VirtualServiceOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.VirtualServiceOptions}
 */
proto.gloo.solo.io.VirtualServiceOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setOneWayTls(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.VirtualServiceOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.VirtualServiceOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.VirtualServiceOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.VirtualServiceOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getOneWayTls();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.BoolValue one_way_tls = 1;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.VirtualServiceOptions.prototype.getOneWayTls = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 1));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.VirtualServiceOptions.prototype.setOneWayTls = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.VirtualServiceOptions.prototype.clearOneWayTls = function() {
  this.setOneWayTls(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.VirtualServiceOptions.prototype.hasOneWayTls = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.GatewayOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.GatewayOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.GatewayOptions.displayName = 'proto.gloo.solo.io.GatewayOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.GatewayOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.GatewayOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.GatewayOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GatewayOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    validationServerAddr: jspb.Message.getFieldWithDefault(msg, 1, ""),
    validation: (f = msg.getValidation()) && proto.gloo.solo.io.GatewayOptions.ValidationOptions.toObject(includeInstance, f),
    readGatewaysFromAllNamespaces: jspb.Message.getFieldWithDefault(msg, 4, false),
    alwaysSortRouteTableRoutes: jspb.Message.getFieldWithDefault(msg, 5, false),
    compressedProxySpec: jspb.Message.getFieldWithDefault(msg, 6, false),
    virtualServiceOptions: (f = msg.getVirtualServiceOptions()) && proto.gloo.solo.io.VirtualServiceOptions.toObject(includeInstance, f),
    persistProxySpec: (f = msg.getPersistProxySpec()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    enableGatewayController: (f = msg.getEnableGatewayController()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    isolateVirtualHostsBySslConfig: (f = msg.getIsolateVirtualHostsBySslConfig()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.GatewayOptions}
 */
proto.gloo.solo.io.GatewayOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.GatewayOptions;
  return proto.gloo.solo.io.GatewayOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.GatewayOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.GatewayOptions}
 */
proto.gloo.solo.io.GatewayOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setValidationServerAddr(value);
      break;
    case 3:
      var value = new proto.gloo.solo.io.GatewayOptions.ValidationOptions;
      reader.readMessage(value,proto.gloo.solo.io.GatewayOptions.ValidationOptions.deserializeBinaryFromReader);
      msg.setValidation(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setReadGatewaysFromAllNamespaces(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setAlwaysSortRouteTableRoutes(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setCompressedProxySpec(value);
      break;
    case 7:
      var value = new proto.gloo.solo.io.VirtualServiceOptions;
      reader.readMessage(value,proto.gloo.solo.io.VirtualServiceOptions.deserializeBinaryFromReader);
      msg.setVirtualServiceOptions(value);
      break;
    case 8:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setPersistProxySpec(value);
      break;
    case 9:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setEnableGatewayController(value);
      break;
    case 10:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setIsolateVirtualHostsBySslConfig(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.GatewayOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.GatewayOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.GatewayOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GatewayOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValidationServerAddr();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getValidation();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.gloo.solo.io.GatewayOptions.ValidationOptions.serializeBinaryToWriter
    );
  }
  f = message.getReadGatewaysFromAllNamespaces();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getAlwaysSortRouteTableRoutes();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = message.getCompressedProxySpec();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
  f = message.getVirtualServiceOptions();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      proto.gloo.solo.io.VirtualServiceOptions.serializeBinaryToWriter
    );
  }
  f = message.getPersistProxySpec();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getEnableGatewayController();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getIsolateVirtualHostsBySslConfig();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.GatewayOptions.ValidationOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.GatewayOptions.ValidationOptions.displayName = 'proto.gloo.solo.io.GatewayOptions.ValidationOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.GatewayOptions.ValidationOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.GatewayOptions.ValidationOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    proxyValidationServerAddr: jspb.Message.getFieldWithDefault(msg, 2, ""),
    validationWebhookTlsCert: jspb.Message.getFieldWithDefault(msg, 3, ""),
    validationWebhookTlsKey: jspb.Message.getFieldWithDefault(msg, 4, ""),
    ignoreGlooValidationFailure: jspb.Message.getFieldWithDefault(msg, 5, false),
    alwaysAccept: (f = msg.getAlwaysAccept()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    allowWarnings: (f = msg.getAllowWarnings()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    warnRouteShortCircuiting: (f = msg.getWarnRouteShortCircuiting()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    disableTransformationValidation: (f = msg.getDisableTransformationValidation()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    validationServerGrpcMaxSizeBytes: (f = msg.getValidationServerGrpcMaxSizeBytes()) && google_protobuf_wrappers_pb.Int32Value.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.GatewayOptions.ValidationOptions}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.GatewayOptions.ValidationOptions;
  return proto.gloo.solo.io.GatewayOptions.ValidationOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.GatewayOptions.ValidationOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.GatewayOptions.ValidationOptions}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setProxyValidationServerAddr(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setValidationWebhookTlsCert(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setValidationWebhookTlsKey(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIgnoreGlooValidationFailure(value);
      break;
    case 6:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setAlwaysAccept(value);
      break;
    case 7:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setAllowWarnings(value);
      break;
    case 8:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setWarnRouteShortCircuiting(value);
      break;
    case 9:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setDisableTransformationValidation(value);
      break;
    case 11:
      var value = new google_protobuf_wrappers_pb.Int32Value;
      reader.readMessage(value,google_protobuf_wrappers_pb.Int32Value.deserializeBinaryFromReader);
      msg.setValidationServerGrpcMaxSizeBytes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.GatewayOptions.ValidationOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.GatewayOptions.ValidationOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProxyValidationServerAddr();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getValidationWebhookTlsCert();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getValidationWebhookTlsKey();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getIgnoreGlooValidationFailure();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = message.getAlwaysAccept();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getAllowWarnings();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getWarnRouteShortCircuiting();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getDisableTransformationValidation();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getValidationServerGrpcMaxSizeBytes();
  if (f != null) {
    writer.writeMessage(
      11,
      f,
      google_protobuf_wrappers_pb.Int32Value.serializeBinaryToWriter
    );
  }
};


/**
 * optional string proxy_validation_server_addr = 2;
 * @return {string}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getProxyValidationServerAddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setProxyValidationServerAddr = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string validation_webhook_tls_cert = 3;
 * @return {string}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getValidationWebhookTlsCert = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setValidationWebhookTlsCert = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string validation_webhook_tls_key = 4;
 * @return {string}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getValidationWebhookTlsKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setValidationWebhookTlsKey = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional bool ignore_gloo_validation_failure = 5;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getIgnoreGlooValidationFailure = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 5, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setIgnoreGlooValidationFailure = function(value) {
  jspb.Message.setProto3BooleanField(this, 5, value);
};


/**
 * optional google.protobuf.BoolValue always_accept = 6;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getAlwaysAccept = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 6));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setAlwaysAccept = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.clearAlwaysAccept = function() {
  this.setAlwaysAccept(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.hasAlwaysAccept = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional google.protobuf.BoolValue allow_warnings = 7;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getAllowWarnings = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 7));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setAllowWarnings = function(value) {
  jspb.Message.setWrapperField(this, 7, value);
};


proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.clearAllowWarnings = function() {
  this.setAllowWarnings(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.hasAllowWarnings = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional google.protobuf.BoolValue warn_route_short_circuiting = 8;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getWarnRouteShortCircuiting = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 8));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setWarnRouteShortCircuiting = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.clearWarnRouteShortCircuiting = function() {
  this.setWarnRouteShortCircuiting(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.hasWarnRouteShortCircuiting = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional google.protobuf.BoolValue disable_transformation_validation = 9;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getDisableTransformationValidation = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 9));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setDisableTransformationValidation = function(value) {
  jspb.Message.setWrapperField(this, 9, value);
};


proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.clearDisableTransformationValidation = function() {
  this.setDisableTransformationValidation(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.hasDisableTransformationValidation = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional google.protobuf.Int32Value validation_server_grpc_max_size_bytes = 11;
 * @return {?proto.google.protobuf.Int32Value}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.getValidationServerGrpcMaxSizeBytes = function() {
  return /** @type{?proto.google.protobuf.Int32Value} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.Int32Value, 11));
};


/** @param {?proto.google.protobuf.Int32Value|undefined} value */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.setValidationServerGrpcMaxSizeBytes = function(value) {
  jspb.Message.setWrapperField(this, 11, value);
};


proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.clearValidationServerGrpcMaxSizeBytes = function() {
  this.setValidationServerGrpcMaxSizeBytes(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.ValidationOptions.prototype.hasValidationServerGrpcMaxSizeBytes = function() {
  return jspb.Message.getField(this, 11) != null;
};


/**
 * optional string validation_server_addr = 1;
 * @return {string}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getValidationServerAddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.GatewayOptions.prototype.setValidationServerAddr = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional ValidationOptions validation = 3;
 * @return {?proto.gloo.solo.io.GatewayOptions.ValidationOptions}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getValidation = function() {
  return /** @type{?proto.gloo.solo.io.GatewayOptions.ValidationOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.GatewayOptions.ValidationOptions, 3));
};


/** @param {?proto.gloo.solo.io.GatewayOptions.ValidationOptions|undefined} value */
proto.gloo.solo.io.GatewayOptions.prototype.setValidation = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.gloo.solo.io.GatewayOptions.prototype.clearValidation = function() {
  this.setValidation(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.prototype.hasValidation = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional bool read_gateways_from_all_namespaces = 4;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getReadGatewaysFromAllNamespaces = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 4, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.GatewayOptions.prototype.setReadGatewaysFromAllNamespaces = function(value) {
  jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional bool always_sort_route_table_routes = 5;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getAlwaysSortRouteTableRoutes = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 5, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.GatewayOptions.prototype.setAlwaysSortRouteTableRoutes = function(value) {
  jspb.Message.setProto3BooleanField(this, 5, value);
};


/**
 * optional bool compressed_proxy_spec = 6;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getCompressedProxySpec = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 6, false));
};


/** @param {boolean} value */
proto.gloo.solo.io.GatewayOptions.prototype.setCompressedProxySpec = function(value) {
  jspb.Message.setProto3BooleanField(this, 6, value);
};


/**
 * optional VirtualServiceOptions virtual_service_options = 7;
 * @return {?proto.gloo.solo.io.VirtualServiceOptions}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getVirtualServiceOptions = function() {
  return /** @type{?proto.gloo.solo.io.VirtualServiceOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.VirtualServiceOptions, 7));
};


/** @param {?proto.gloo.solo.io.VirtualServiceOptions|undefined} value */
proto.gloo.solo.io.GatewayOptions.prototype.setVirtualServiceOptions = function(value) {
  jspb.Message.setWrapperField(this, 7, value);
};


proto.gloo.solo.io.GatewayOptions.prototype.clearVirtualServiceOptions = function() {
  this.setVirtualServiceOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.prototype.hasVirtualServiceOptions = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional google.protobuf.BoolValue persist_proxy_spec = 8;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getPersistProxySpec = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 8));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GatewayOptions.prototype.setPersistProxySpec = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


proto.gloo.solo.io.GatewayOptions.prototype.clearPersistProxySpec = function() {
  this.setPersistProxySpec(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.prototype.hasPersistProxySpec = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional google.protobuf.BoolValue enable_gateway_controller = 9;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getEnableGatewayController = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 9));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GatewayOptions.prototype.setEnableGatewayController = function(value) {
  jspb.Message.setWrapperField(this, 9, value);
};


proto.gloo.solo.io.GatewayOptions.prototype.clearEnableGatewayController = function() {
  this.setEnableGatewayController(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.prototype.hasEnableGatewayController = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional google.protobuf.BoolValue isolate_virtual_hosts_by_ssl_config = 10;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GatewayOptions.prototype.getIsolateVirtualHostsBySslConfig = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 10));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GatewayOptions.prototype.setIsolateVirtualHostsBySslConfig = function(value) {
  jspb.Message.setWrapperField(this, 10, value);
};


proto.gloo.solo.io.GatewayOptions.prototype.clearIsolateVirtualHostsBySslConfig = function() {
  this.setIsolateVirtualHostsBySslConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GatewayOptions.prototype.hasIsolateVirtualHostsBySslConfig = function() {
  return jspb.Message.getField(this, 10) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.ConsoleOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.ConsoleOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.ConsoleOptions.displayName = 'proto.gloo.solo.io.ConsoleOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.ConsoleOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.ConsoleOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.ConsoleOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.ConsoleOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    readOnly: (f = msg.getReadOnly()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    apiExplorerEnabled: (f = msg.getApiExplorerEnabled()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.ConsoleOptions}
 */
proto.gloo.solo.io.ConsoleOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.ConsoleOptions;
  return proto.gloo.solo.io.ConsoleOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.ConsoleOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.ConsoleOptions}
 */
proto.gloo.solo.io.ConsoleOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setReadOnly(value);
      break;
    case 2:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setApiExplorerEnabled(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.ConsoleOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.ConsoleOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.ConsoleOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.ConsoleOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getReadOnly();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getApiExplorerEnabled();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.BoolValue read_only = 1;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.ConsoleOptions.prototype.getReadOnly = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 1));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.ConsoleOptions.prototype.setReadOnly = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.ConsoleOptions.prototype.clearReadOnly = function() {
  this.setReadOnly(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.ConsoleOptions.prototype.hasReadOnly = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.BoolValue api_explorer_enabled = 2;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.ConsoleOptions.prototype.getApiExplorerEnabled = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 2));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.ConsoleOptions.prototype.setApiExplorerEnabled = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.gloo.solo.io.ConsoleOptions.prototype.clearApiExplorerEnabled = function() {
  this.setApiExplorerEnabled(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.ConsoleOptions.prototype.hasApiExplorerEnabled = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.GraphqlOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.GraphqlOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.GraphqlOptions.displayName = 'proto.gloo.solo.io.GraphqlOptions';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.GraphqlOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.GraphqlOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.GraphqlOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GraphqlOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    schemaChangeValidationOptions: (f = msg.getSchemaChangeValidationOptions()) && proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.GraphqlOptions}
 */
proto.gloo.solo.io.GraphqlOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.GraphqlOptions;
  return proto.gloo.solo.io.GraphqlOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.GraphqlOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.GraphqlOptions}
 */
proto.gloo.solo.io.GraphqlOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions;
      reader.readMessage(value,proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.deserializeBinaryFromReader);
      msg.setSchemaChangeValidationOptions(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.GraphqlOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.GraphqlOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.GraphqlOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GraphqlOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSchemaChangeValidationOptions();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.serializeBinaryToWriter
    );
  }
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.repeatedFields_, null);
};
goog.inherits(proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.displayName = 'proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    rejectBreakingChanges: (f = msg.getRejectBreakingChanges()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    processingRulesList: jspb.Message.getRepeatedField(msg, 2)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions}
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions;
  return proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions}
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setRejectBreakingChanges(value);
      break;
    case 2:
      var value = /** @type {!Array<!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule>} */ (reader.readPackedEnum());
      msg.setProcessingRulesList(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRejectBreakingChanges();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getProcessingRulesList();
  if (f.length > 0) {
    writer.writePackedEnum(
      2,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule = {
  RULE_UNSPECIFIED: 0,
  RULE_DANGEROUS_TO_BREAKING: 1,
  RULE_DEPRECATED_FIELD_REMOVAL_DANGEROUS: 2,
  RULE_IGNORE_DESCRIPTION_CHANGES: 3,
  RULE_IGNORE_UNREACHABLE: 4
};

/**
 * optional google.protobuf.BoolValue reject_breaking_changes = 1;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.getRejectBreakingChanges = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 1));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.setRejectBreakingChanges = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.clearRejectBreakingChanges = function() {
  this.setRejectBreakingChanges(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.hasRejectBreakingChanges = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated ProcessingRule processing_rules = 2;
 * @return {!Array<!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule>}
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.getProcessingRulesList = function() {
  return /** @type {!Array<!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule>} */ (jspb.Message.getRepeatedField(this, 2));
};


/** @param {!Array<!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule>} value */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.setProcessingRulesList = function(value) {
  jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {!proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule} value
 * @param {number=} opt_index
 */
proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.addProcessingRules = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.prototype.clearProcessingRulesList = function() {
  this.setProcessingRulesList([]);
};


/**
 * optional SchemaChangeValidationOptions schema_change_validation_options = 1;
 * @return {?proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions}
 */
proto.gloo.solo.io.GraphqlOptions.prototype.getSchemaChangeValidationOptions = function() {
  return /** @type{?proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions, 1));
};


/** @param {?proto.gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions|undefined} value */
proto.gloo.solo.io.GraphqlOptions.prototype.setSchemaChangeValidationOptions = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.GraphqlOptions.prototype.clearSchemaChangeValidationOptions = function() {
  this.setSchemaChangeValidationOptions(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.GraphqlOptions.prototype.hasSchemaChangeValidationOptions = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsStatus = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsStatus, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsStatus.displayName = 'proto.gloo.solo.io.SettingsStatus';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsStatus.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsStatus.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsStatus} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsStatus.toObject = function(includeInstance, msg) {
  var f, obj = {
    state: jspb.Message.getFieldWithDefault(msg, 1, 0),
    reason: jspb.Message.getFieldWithDefault(msg, 2, ""),
    reportedBy: jspb.Message.getFieldWithDefault(msg, 3, ""),
    subresourceStatusesMap: (f = msg.getSubresourceStatusesMap()) ? f.toObject(includeInstance, proto.gloo.solo.io.SettingsStatus.toObject) : [],
    details: (f = msg.getDetails()) && google_protobuf_struct_pb.Struct.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsStatus}
 */
proto.gloo.solo.io.SettingsStatus.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsStatus;
  return proto.gloo.solo.io.SettingsStatus.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsStatus} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsStatus}
 */
proto.gloo.solo.io.SettingsStatus.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.gloo.solo.io.SettingsStatus.State} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setReason(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setReportedBy(value);
      break;
    case 4:
      var value = msg.getSubresourceStatusesMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.gloo.solo.io.SettingsStatus.deserializeBinaryFromReader, "");
         });
      break;
    case 5:
      var value = new google_protobuf_struct_pb.Struct;
      reader.readMessage(value,google_protobuf_struct_pb.Struct.deserializeBinaryFromReader);
      msg.setDetails(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsStatus.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsStatus.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsStatus} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsStatus.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getReason();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getReportedBy();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getSubresourceStatusesMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.gloo.solo.io.SettingsStatus.serializeBinaryToWriter);
  }
  f = message.getDetails();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_struct_pb.Struct.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.gloo.solo.io.SettingsStatus.State = {
  PENDING: 0,
  ACCEPTED: 1,
  REJECTED: 2,
  WARNING: 3
};

/**
 * optional State state = 1;
 * @return {!proto.gloo.solo.io.SettingsStatus.State}
 */
proto.gloo.solo.io.SettingsStatus.prototype.getState = function() {
  return /** @type {!proto.gloo.solo.io.SettingsStatus.State} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.gloo.solo.io.SettingsStatus.State} value */
proto.gloo.solo.io.SettingsStatus.prototype.setState = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string reason = 2;
 * @return {string}
 */
proto.gloo.solo.io.SettingsStatus.prototype.getReason = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsStatus.prototype.setReason = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string reported_by = 3;
 * @return {string}
 */
proto.gloo.solo.io.SettingsStatus.prototype.getReportedBy = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.gloo.solo.io.SettingsStatus.prototype.setReportedBy = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * map<string, SettingsStatus> subresource_statuses = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.gloo.solo.io.SettingsStatus>}
 */
proto.gloo.solo.io.SettingsStatus.prototype.getSubresourceStatusesMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.gloo.solo.io.SettingsStatus>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      proto.gloo.solo.io.SettingsStatus));
};


proto.gloo.solo.io.SettingsStatus.prototype.clearSubresourceStatusesMap = function() {
  this.getSubresourceStatusesMap().clear();
};


/**
 * optional google.protobuf.Struct details = 5;
 * @return {?proto.google.protobuf.Struct}
 */
proto.gloo.solo.io.SettingsStatus.prototype.getDetails = function() {
  return /** @type{?proto.google.protobuf.Struct} */ (
    jspb.Message.getWrapperField(this, google_protobuf_struct_pb.Struct, 5));
};


/** @param {?proto.google.protobuf.Struct|undefined} value */
proto.gloo.solo.io.SettingsStatus.prototype.setDetails = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.gloo.solo.io.SettingsStatus.prototype.clearDetails = function() {
  this.setDetails(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.SettingsStatus.prototype.hasDetails = function() {
  return jspb.Message.getField(this, 5) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.gloo.solo.io.SettingsNamespacedStatuses = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.SettingsNamespacedStatuses, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.SettingsNamespacedStatuses.displayName = 'proto.gloo.solo.io.SettingsNamespacedStatuses';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.gloo.solo.io.SettingsNamespacedStatuses.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.SettingsNamespacedStatuses.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.SettingsNamespacedStatuses} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsNamespacedStatuses.toObject = function(includeInstance, msg) {
  var f, obj = {
    statusesMap: (f = msg.getStatusesMap()) ? f.toObject(includeInstance, proto.gloo.solo.io.SettingsStatus.toObject) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.gloo.solo.io.SettingsNamespacedStatuses}
 */
proto.gloo.solo.io.SettingsNamespacedStatuses.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.SettingsNamespacedStatuses;
  return proto.gloo.solo.io.SettingsNamespacedStatuses.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.SettingsNamespacedStatuses} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.SettingsNamespacedStatuses}
 */
proto.gloo.solo.io.SettingsNamespacedStatuses.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = msg.getStatusesMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.gloo.solo.io.SettingsStatus.deserializeBinaryFromReader, "");
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.gloo.solo.io.SettingsNamespacedStatuses.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.SettingsNamespacedStatuses.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.SettingsNamespacedStatuses} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.SettingsNamespacedStatuses.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStatusesMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(1, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.gloo.solo.io.SettingsStatus.serializeBinaryToWriter);
  }
};


/**
 * map<string, SettingsStatus> statuses = 1;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.gloo.solo.io.SettingsStatus>}
 */
proto.gloo.solo.io.SettingsNamespacedStatuses.prototype.getStatusesMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.gloo.solo.io.SettingsStatus>} */ (
      jspb.Message.getMapField(this, 1, opt_noLazyCreate,
      proto.gloo.solo.io.SettingsStatus));
};


proto.gloo.solo.io.SettingsNamespacedStatuses.prototype.clearStatusesMap = function() {
  this.getStatusesMap().clear();
};


goog.object.extend(exports, proto.gloo.solo.io);
