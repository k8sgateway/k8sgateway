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

var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
var envoy_annotations_deprecation_pb = require('../../../../envoy/annotations/deprecation_pb.js');
var udpa_annotations_migrate_pb = require('../../../../udpa/annotations/migrate_pb.js');
var udpa_annotations_status_pb = require('../../../../udpa/annotations/status_pb.js');
var udpa_annotations_versioning_pb = require('../../../../udpa/annotations/versioning_pb.js');
var validate_validate_pb = require('../../../../validate/validate_pb.js');
var solo$kit_api_v1_ref_pb = require('../../../../solo-kit/api/v1/ref_pb.js');
var gogoproto_gogo_pb = require('../../../../gogoproto/gogo_pb.js');
goog.exportSymbol('proto.envoy.config.trace.v3.ZipkinConfig', null, global);
goog.exportSymbol('proto.envoy.config.trace.v3.ZipkinConfig.CollectorEndpointVersion', null, global);

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
proto.envoy.config.trace.v3.ZipkinConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.envoy.config.trace.v3.ZipkinConfig.oneofGroups_);
};
goog.inherits(proto.envoy.config.trace.v3.ZipkinConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.config.trace.v3.ZipkinConfig.displayName = 'proto.envoy.config.trace.v3.ZipkinConfig';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.envoy.config.trace.v3.ZipkinConfig.oneofGroups_ = [[1,6]];

/**
 * @enum {number}
 */
proto.envoy.config.trace.v3.ZipkinConfig.CollectorClusterCase = {
  COLLECTOR_CLUSTER_NOT_SET: 0,
  COLLECTOR_UPSTREAM_REF: 1,
  CLUSTER_NAME: 6
};

/**
 * @return {proto.envoy.config.trace.v3.ZipkinConfig.CollectorClusterCase}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.getCollectorClusterCase = function() {
  return /** @type {proto.envoy.config.trace.v3.ZipkinConfig.CollectorClusterCase} */(jspb.Message.computeOneofCase(this, proto.envoy.config.trace.v3.ZipkinConfig.oneofGroups_[0]));
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
proto.envoy.config.trace.v3.ZipkinConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.config.trace.v3.ZipkinConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.config.trace.v3.ZipkinConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.trace.v3.ZipkinConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    collectorUpstreamRef: (f = msg.getCollectorUpstreamRef()) && solo$kit_api_v1_ref_pb.ResourceRef.toObject(includeInstance, f),
    clusterName: jspb.Message.getFieldWithDefault(msg, 6, ""),
    collectorEndpoint: jspb.Message.getFieldWithDefault(msg, 2, ""),
    traceId128bit: jspb.Message.getFieldWithDefault(msg, 3, false),
    sharedSpanContext: (f = msg.getSharedSpanContext()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    collectorEndpointVersion: jspb.Message.getFieldWithDefault(msg, 5, 0)
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
 * @return {!proto.envoy.config.trace.v3.ZipkinConfig}
 */
proto.envoy.config.trace.v3.ZipkinConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.config.trace.v3.ZipkinConfig;
  return proto.envoy.config.trace.v3.ZipkinConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.config.trace.v3.ZipkinConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.config.trace.v3.ZipkinConfig}
 */
proto.envoy.config.trace.v3.ZipkinConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new solo$kit_api_v1_ref_pb.ResourceRef;
      reader.readMessage(value,solo$kit_api_v1_ref_pb.ResourceRef.deserializeBinaryFromReader);
      msg.setCollectorUpstreamRef(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setClusterName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setCollectorEndpoint(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setTraceId128bit(value);
      break;
    case 4:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setSharedSpanContext(value);
      break;
    case 5:
      var value = /** @type {!proto.envoy.config.trace.v3.ZipkinConfig.CollectorEndpointVersion} */ (reader.readEnum());
      msg.setCollectorEndpointVersion(value);
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
proto.envoy.config.trace.v3.ZipkinConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.config.trace.v3.ZipkinConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.config.trace.v3.ZipkinConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.trace.v3.ZipkinConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCollectorUpstreamRef();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      solo$kit_api_v1_ref_pb.ResourceRef.serializeBinaryToWriter
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 6));
  if (f != null) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getCollectorEndpoint();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getTraceId128bit();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getSharedSpanContext();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getCollectorEndpointVersion();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.envoy.config.trace.v3.ZipkinConfig.CollectorEndpointVersion = {
  DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE: 0,
  HTTP_JSON: 1,
  HTTP_PROTO: 2
};

/**
 * optional core.solo.io.ResourceRef collector_upstream_ref = 1;
 * @return {?proto.core.solo.io.ResourceRef}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.getCollectorUpstreamRef = function() {
  return /** @type{?proto.core.solo.io.ResourceRef} */ (
    jspb.Message.getWrapperField(this, solo$kit_api_v1_ref_pb.ResourceRef, 1));
};


/** @param {?proto.core.solo.io.ResourceRef|undefined} value */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.setCollectorUpstreamRef = function(value) {
  jspb.Message.setOneofWrapperField(this, 1, proto.envoy.config.trace.v3.ZipkinConfig.oneofGroups_[0], value);
};


proto.envoy.config.trace.v3.ZipkinConfig.prototype.clearCollectorUpstreamRef = function() {
  this.setCollectorUpstreamRef(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.hasCollectorUpstreamRef = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string cluster_name = 6;
 * @return {string}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.getClusterName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.setClusterName = function(value) {
  jspb.Message.setOneofField(this, 6, proto.envoy.config.trace.v3.ZipkinConfig.oneofGroups_[0], value);
};


proto.envoy.config.trace.v3.ZipkinConfig.prototype.clearClusterName = function() {
  jspb.Message.setOneofField(this, 6, proto.envoy.config.trace.v3.ZipkinConfig.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.hasClusterName = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional string collector_endpoint = 2;
 * @return {string}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.getCollectorEndpoint = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.setCollectorEndpoint = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bool trace_id_128bit = 3;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.getTraceId128bit = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 3, false));
};


/** @param {boolean} value */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.setTraceId128bit = function(value) {
  jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional google.protobuf.BoolValue shared_span_context = 4;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.getSharedSpanContext = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 4));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.setSharedSpanContext = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.envoy.config.trace.v3.ZipkinConfig.prototype.clearSharedSpanContext = function() {
  this.setSharedSpanContext(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.hasSharedSpanContext = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional CollectorEndpointVersion collector_endpoint_version = 5;
 * @return {!proto.envoy.config.trace.v3.ZipkinConfig.CollectorEndpointVersion}
 */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.getCollectorEndpointVersion = function() {
  return /** @type {!proto.envoy.config.trace.v3.ZipkinConfig.CollectorEndpointVersion} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {!proto.envoy.config.trace.v3.ZipkinConfig.CollectorEndpointVersion} value */
proto.envoy.config.trace.v3.ZipkinConfig.prototype.setCollectorEndpointVersion = function(value) {
  jspb.Message.setProto3EnumField(this, 5, value);
};


goog.object.extend(exports, proto.envoy.config.trace.v3);
