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

var gogoproto_gogo_pb = require('../../../../../../gogo/protobuf/gogoproto/gogo_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
goog.exportSymbol('proto.gloo.solo.io.ConnectionConfig', null, global);
goog.exportSymbol('proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive', null, global);

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
proto.gloo.solo.io.ConnectionConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.ConnectionConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.ConnectionConfig.displayName = 'proto.gloo.solo.io.ConnectionConfig';
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
proto.gloo.solo.io.ConnectionConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.ConnectionConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.ConnectionConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.ConnectionConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    maxRequestsPerConnection: jspb.Message.getFieldWithDefault(msg, 1, 0),
    connectTimeout: (f = msg.getConnectTimeout()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    tcpKeepalive: (f = msg.getTcpKeepalive()) && proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.toObject(includeInstance, f)
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
 * @return {!proto.gloo.solo.io.ConnectionConfig}
 */
proto.gloo.solo.io.ConnectionConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.ConnectionConfig;
  return proto.gloo.solo.io.ConnectionConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.ConnectionConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.ConnectionConfig}
 */
proto.gloo.solo.io.ConnectionConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setMaxRequestsPerConnection(value);
      break;
    case 2:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setConnectTimeout(value);
      break;
    case 3:
      var value = new proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive;
      reader.readMessage(value,proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.deserializeBinaryFromReader);
      msg.setTcpKeepalive(value);
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
proto.gloo.solo.io.ConnectionConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.ConnectionConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.ConnectionConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.ConnectionConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMaxRequestsPerConnection();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getConnectTimeout();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getTcpKeepalive();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.serializeBinaryToWriter
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
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.displayName = 'proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive';
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
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.toObject = function(includeInstance, msg) {
  var f, obj = {
    keepaliveProbes: jspb.Message.getFieldWithDefault(msg, 1, 0),
    keepaliveTime: (f = msg.getKeepaliveTime()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    keepaliveInterval: (f = msg.getKeepaliveInterval()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f)
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
 * @return {!proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive}
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive;
  return proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive}
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setKeepaliveProbes(value);
      break;
    case 2:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setKeepaliveTime(value);
      break;
    case 3:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setKeepaliveInterval(value);
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
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKeepaliveProbes();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getKeepaliveTime();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getKeepaliveInterval();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
};


/**
 * optional uint32 keepalive_probes = 1;
 * @return {number}
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.getKeepaliveProbes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.setKeepaliveProbes = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional google.protobuf.Duration keepalive_time = 2;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.getKeepaliveTime = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 2));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.setKeepaliveTime = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.clearKeepaliveTime = function() {
  this.setKeepaliveTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.hasKeepaliveTime = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.Duration keepalive_interval = 3;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.getKeepaliveInterval = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 3));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.setKeepaliveInterval = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.clearKeepaliveInterval = function() {
  this.setKeepaliveInterval(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive.prototype.hasKeepaliveInterval = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional uint32 max_requests_per_connection = 1;
 * @return {number}
 */
proto.gloo.solo.io.ConnectionConfig.prototype.getMaxRequestsPerConnection = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.gloo.solo.io.ConnectionConfig.prototype.setMaxRequestsPerConnection = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional google.protobuf.Duration connect_timeout = 2;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.ConnectionConfig.prototype.getConnectTimeout = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 2));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.ConnectionConfig.prototype.setConnectTimeout = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.gloo.solo.io.ConnectionConfig.prototype.clearConnectTimeout = function() {
  this.setConnectTimeout(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.ConnectionConfig.prototype.hasConnectTimeout = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional TcpKeepAlive tcp_keepalive = 3;
 * @return {?proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive}
 */
proto.gloo.solo.io.ConnectionConfig.prototype.getTcpKeepalive = function() {
  return /** @type{?proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive, 3));
};


/** @param {?proto.gloo.solo.io.ConnectionConfig.TcpKeepAlive|undefined} value */
proto.gloo.solo.io.ConnectionConfig.prototype.setTcpKeepalive = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.gloo.solo.io.ConnectionConfig.prototype.clearTcpKeepalive = function() {
  this.setTcpKeepalive(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.ConnectionConfig.prototype.hasTcpKeepalive = function() {
  return jspb.Message.getField(this, 3) != null;
};


goog.object.extend(exports, proto.gloo.solo.io);
