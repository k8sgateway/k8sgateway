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

var udpa_annotations_status_pb = require('../../../../udpa/annotations/status_pb.js');
var validate_validate_pb = require('../../../../validate/validate_pb.js');
var gogoproto_gogo_pb = require('../../../../gogoproto/gogo_pb.js');
goog.exportSymbol('proto.envoy.config.core.v3.ProxyProtocolConfig', null, global);
goog.exportSymbol('proto.envoy.config.core.v3.ProxyProtocolConfig.Version', null, global);

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
proto.envoy.config.core.v3.ProxyProtocolConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.config.core.v3.ProxyProtocolConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.config.core.v3.ProxyProtocolConfig.displayName = 'proto.envoy.config.core.v3.ProxyProtocolConfig';
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
proto.envoy.config.core.v3.ProxyProtocolConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.config.core.v3.ProxyProtocolConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.config.core.v3.ProxyProtocolConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.core.v3.ProxyProtocolConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    version: jspb.Message.getFieldWithDefault(msg, 1, 0)
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
 * @return {!proto.envoy.config.core.v3.ProxyProtocolConfig}
 */
proto.envoy.config.core.v3.ProxyProtocolConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.config.core.v3.ProxyProtocolConfig;
  return proto.envoy.config.core.v3.ProxyProtocolConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.config.core.v3.ProxyProtocolConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.config.core.v3.ProxyProtocolConfig}
 */
proto.envoy.config.core.v3.ProxyProtocolConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.envoy.config.core.v3.ProxyProtocolConfig.Version} */ (reader.readEnum());
      msg.setVersion(value);
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
proto.envoy.config.core.v3.ProxyProtocolConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.config.core.v3.ProxyProtocolConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.config.core.v3.ProxyProtocolConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.core.v3.ProxyProtocolConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVersion();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.envoy.config.core.v3.ProxyProtocolConfig.Version = {
  V1: 0,
  V2: 1
};

/**
 * optional Version version = 1;
 * @return {!proto.envoy.config.core.v3.ProxyProtocolConfig.Version}
 */
proto.envoy.config.core.v3.ProxyProtocolConfig.prototype.getVersion = function() {
  return /** @type {!proto.envoy.config.core.v3.ProxyProtocolConfig.Version} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.envoy.config.core.v3.ProxyProtocolConfig.Version} value */
proto.envoy.config.core.v3.ProxyProtocolConfig.prototype.setVersion = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


goog.object.extend(exports, proto.envoy.config.core.v3);
