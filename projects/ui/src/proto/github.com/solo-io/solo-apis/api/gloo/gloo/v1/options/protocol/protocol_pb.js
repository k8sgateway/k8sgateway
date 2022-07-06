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

var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
var extproto_ext_pb = require('../../../../../../../../../extproto/ext_pb.js');
goog.exportSymbol('proto.protocol.options.gloo.solo.io.Http1ProtocolOptions', null, global);
goog.exportSymbol('proto.protocol.options.gloo.solo.io.Http2ProtocolOptions', null, global);
goog.exportSymbol('proto.protocol.options.gloo.solo.io.HttpProtocolOptions', null, global);
goog.exportSymbol('proto.protocol.options.gloo.solo.io.HttpProtocolOptions.HeadersWithUnderscoresAction', null, global);

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
proto.protocol.options.gloo.solo.io.HttpProtocolOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protocol.options.gloo.solo.io.HttpProtocolOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.protocol.options.gloo.solo.io.HttpProtocolOptions.displayName = 'proto.protocol.options.gloo.solo.io.HttpProtocolOptions';
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
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.protocol.options.gloo.solo.io.HttpProtocolOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    idleTimeout: (f = msg.getIdleTimeout()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    maxHeadersCount: jspb.Message.getFieldWithDefault(msg, 2, 0),
    maxStreamDuration: (f = msg.getMaxStreamDuration()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    headersWithUnderscoresAction: jspb.Message.getFieldWithDefault(msg, 4, 0)
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
 * @return {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protocol.options.gloo.solo.io.HttpProtocolOptions;
  return proto.protocol.options.gloo.solo.io.HttpProtocolOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setIdleTimeout(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setMaxHeadersCount(value);
      break;
    case 3:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setMaxStreamDuration(value);
      break;
    case 4:
      var value = /** @type {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions.HeadersWithUnderscoresAction} */ (reader.readEnum());
      msg.setHeadersWithUnderscoresAction(value);
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
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protocol.options.gloo.solo.io.HttpProtocolOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIdleTimeout();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getMaxHeadersCount();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
  f = message.getMaxStreamDuration();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getHeadersWithUnderscoresAction();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.HeadersWithUnderscoresAction = {
  ALLOW: 0,
  REJECT_REQUEST: 1,
  DROP_HEADER: 2
};

/**
 * optional google.protobuf.Duration idle_timeout = 1;
 * @return {?proto.google.protobuf.Duration}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.getIdleTimeout = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 1));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.setIdleTimeout = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.clearIdleTimeout = function() {
  this.setIdleTimeout(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.hasIdleTimeout = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional uint32 max_headers_count = 2;
 * @return {number}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.getMaxHeadersCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.setMaxHeadersCount = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional google.protobuf.Duration max_stream_duration = 3;
 * @return {?proto.google.protobuf.Duration}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.getMaxStreamDuration = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 3));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.setMaxStreamDuration = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.clearMaxStreamDuration = function() {
  this.setMaxStreamDuration(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.hasMaxStreamDuration = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional HeadersWithUnderscoresAction headers_with_underscores_action = 4;
 * @return {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions.HeadersWithUnderscoresAction}
 */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.getHeadersWithUnderscoresAction = function() {
  return /** @type {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions.HeadersWithUnderscoresAction} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {!proto.protocol.options.gloo.solo.io.HttpProtocolOptions.HeadersWithUnderscoresAction} value */
proto.protocol.options.gloo.solo.io.HttpProtocolOptions.prototype.setHeadersWithUnderscoresAction = function(value) {
  jspb.Message.setProto3EnumField(this, 4, value);
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
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.oneofGroups_);
};
goog.inherits(proto.protocol.options.gloo.solo.io.Http1ProtocolOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.displayName = 'proto.protocol.options.gloo.solo.io.Http1ProtocolOptions';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.oneofGroups_ = [[22,31]];

/**
 * @enum {number}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.HeaderFormatCase = {
  HEADER_FORMAT_NOT_SET: 0,
  PROPER_CASE_HEADER_KEY_FORMAT: 22,
  PRESERVE_CASE_HEADER_KEY_FORMAT: 31
};

/**
 * @return {proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.HeaderFormatCase}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.getHeaderFormatCase = function() {
  return /** @type {proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.HeaderFormatCase} */(jspb.Message.computeOneofCase(this, proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.oneofGroups_[0]));
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
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protocol.options.gloo.solo.io.Http1ProtocolOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    enableTrailers: jspb.Message.getFieldWithDefault(msg, 1, false),
    properCaseHeaderKeyFormat: jspb.Message.getFieldWithDefault(msg, 22, false),
    preserveCaseHeaderKeyFormat: jspb.Message.getFieldWithDefault(msg, 31, false),
    overrideStreamErrorOnInvalidHttpMessage: (f = msg.getOverrideStreamErrorOnInvalidHttpMessage()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f)
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
 * @return {!proto.protocol.options.gloo.solo.io.Http1ProtocolOptions}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protocol.options.gloo.solo.io.Http1ProtocolOptions;
  return proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protocol.options.gloo.solo.io.Http1ProtocolOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protocol.options.gloo.solo.io.Http1ProtocolOptions}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setEnableTrailers(value);
      break;
    case 22:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setProperCaseHeaderKeyFormat(value);
      break;
    case 31:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setPreserveCaseHeaderKeyFormat(value);
      break;
    case 2:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setOverrideStreamErrorOnInvalidHttpMessage(value);
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
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protocol.options.gloo.solo.io.Http1ProtocolOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEnableTrailers();
  if (f) {
    writer.writeBool(
      1,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 22));
  if (f != null) {
    writer.writeBool(
      22,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 31));
  if (f != null) {
    writer.writeBool(
      31,
      f
    );
  }
  f = message.getOverrideStreamErrorOnInvalidHttpMessage();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
};


/**
 * optional bool enable_trailers = 1;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.getEnableTrailers = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 1, false));
};


/** @param {boolean} value */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.setEnableTrailers = function(value) {
  jspb.Message.setProto3BooleanField(this, 1, value);
};


/**
 * optional bool proper_case_header_key_format = 22;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.getProperCaseHeaderKeyFormat = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 22, false));
};


/** @param {boolean} value */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.setProperCaseHeaderKeyFormat = function(value) {
  jspb.Message.setOneofField(this, 22, proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.oneofGroups_[0], value);
};


proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.clearProperCaseHeaderKeyFormat = function() {
  jspb.Message.setOneofField(this, 22, proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.hasProperCaseHeaderKeyFormat = function() {
  return jspb.Message.getField(this, 22) != null;
};


/**
 * optional bool preserve_case_header_key_format = 31;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.getPreserveCaseHeaderKeyFormat = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 31, false));
};


/** @param {boolean} value */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.setPreserveCaseHeaderKeyFormat = function(value) {
  jspb.Message.setOneofField(this, 31, proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.oneofGroups_[0], value);
};


proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.clearPreserveCaseHeaderKeyFormat = function() {
  jspb.Message.setOneofField(this, 31, proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.hasPreserveCaseHeaderKeyFormat = function() {
  return jspb.Message.getField(this, 31) != null;
};


/**
 * optional google.protobuf.BoolValue override_stream_error_on_invalid_http_message = 2;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.getOverrideStreamErrorOnInvalidHttpMessage = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 2));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.setOverrideStreamErrorOnInvalidHttpMessage = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.clearOverrideStreamErrorOnInvalidHttpMessage = function() {
  this.setOverrideStreamErrorOnInvalidHttpMessage(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.Http1ProtocolOptions.prototype.hasOverrideStreamErrorOnInvalidHttpMessage = function() {
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
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protocol.options.gloo.solo.io.Http2ProtocolOptions, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.displayName = 'proto.protocol.options.gloo.solo.io.Http2ProtocolOptions';
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
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.toObject = function(opt_includeInstance) {
  return proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protocol.options.gloo.solo.io.Http2ProtocolOptions} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.toObject = function(includeInstance, msg) {
  var f, obj = {
    maxConcurrentStreams: (f = msg.getMaxConcurrentStreams()) && google_protobuf_wrappers_pb.UInt32Value.toObject(includeInstance, f),
    initialStreamWindowSize: (f = msg.getInitialStreamWindowSize()) && google_protobuf_wrappers_pb.UInt32Value.toObject(includeInstance, f),
    initialConnectionWindowSize: (f = msg.getInitialConnectionWindowSize()) && google_protobuf_wrappers_pb.UInt32Value.toObject(includeInstance, f),
    overrideStreamErrorOnInvalidHttpMessage: (f = msg.getOverrideStreamErrorOnInvalidHttpMessage()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f)
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
 * @return {!proto.protocol.options.gloo.solo.io.Http2ProtocolOptions}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protocol.options.gloo.solo.io.Http2ProtocolOptions;
  return proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protocol.options.gloo.solo.io.Http2ProtocolOptions} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protocol.options.gloo.solo.io.Http2ProtocolOptions}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = new google_protobuf_wrappers_pb.UInt32Value;
      reader.readMessage(value,google_protobuf_wrappers_pb.UInt32Value.deserializeBinaryFromReader);
      msg.setMaxConcurrentStreams(value);
      break;
    case 3:
      var value = new google_protobuf_wrappers_pb.UInt32Value;
      reader.readMessage(value,google_protobuf_wrappers_pb.UInt32Value.deserializeBinaryFromReader);
      msg.setInitialStreamWindowSize(value);
      break;
    case 4:
      var value = new google_protobuf_wrappers_pb.UInt32Value;
      reader.readMessage(value,google_protobuf_wrappers_pb.UInt32Value.deserializeBinaryFromReader);
      msg.setInitialConnectionWindowSize(value);
      break;
    case 14:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setOverrideStreamErrorOnInvalidHttpMessage(value);
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
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protocol.options.gloo.solo.io.Http2ProtocolOptions} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMaxConcurrentStreams();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_wrappers_pb.UInt32Value.serializeBinaryToWriter
    );
  }
  f = message.getInitialStreamWindowSize();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_wrappers_pb.UInt32Value.serializeBinaryToWriter
    );
  }
  f = message.getInitialConnectionWindowSize();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_wrappers_pb.UInt32Value.serializeBinaryToWriter
    );
  }
  f = message.getOverrideStreamErrorOnInvalidHttpMessage();
  if (f != null) {
    writer.writeMessage(
      14,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.UInt32Value max_concurrent_streams = 2;
 * @return {?proto.google.protobuf.UInt32Value}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.getMaxConcurrentStreams = function() {
  return /** @type{?proto.google.protobuf.UInt32Value} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.UInt32Value, 2));
};


/** @param {?proto.google.protobuf.UInt32Value|undefined} value */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.setMaxConcurrentStreams = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.clearMaxConcurrentStreams = function() {
  this.setMaxConcurrentStreams(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.hasMaxConcurrentStreams = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.UInt32Value initial_stream_window_size = 3;
 * @return {?proto.google.protobuf.UInt32Value}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.getInitialStreamWindowSize = function() {
  return /** @type{?proto.google.protobuf.UInt32Value} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.UInt32Value, 3));
};


/** @param {?proto.google.protobuf.UInt32Value|undefined} value */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.setInitialStreamWindowSize = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.clearInitialStreamWindowSize = function() {
  this.setInitialStreamWindowSize(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.hasInitialStreamWindowSize = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.UInt32Value initial_connection_window_size = 4;
 * @return {?proto.google.protobuf.UInt32Value}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.getInitialConnectionWindowSize = function() {
  return /** @type{?proto.google.protobuf.UInt32Value} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.UInt32Value, 4));
};


/** @param {?proto.google.protobuf.UInt32Value|undefined} value */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.setInitialConnectionWindowSize = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.clearInitialConnectionWindowSize = function() {
  this.setInitialConnectionWindowSize(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.hasInitialConnectionWindowSize = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.BoolValue override_stream_error_on_invalid_http_message = 14;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.getOverrideStreamErrorOnInvalidHttpMessage = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 14));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.setOverrideStreamErrorOnInvalidHttpMessage = function(value) {
  jspb.Message.setWrapperField(this, 14, value);
};


proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.clearOverrideStreamErrorOnInvalidHttpMessage = function() {
  this.setOverrideStreamErrorOnInvalidHttpMessage(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.protocol.options.gloo.solo.io.Http2ProtocolOptions.prototype.hasOverrideStreamErrorOnInvalidHttpMessage = function() {
  return jspb.Message.getField(this, 14) != null;
};


goog.object.extend(exports, proto.protocol.options.gloo.solo.io);
