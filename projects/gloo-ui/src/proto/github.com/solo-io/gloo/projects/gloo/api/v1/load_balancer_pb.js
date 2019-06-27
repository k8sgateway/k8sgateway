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

var gogoproto_gogo_pb = require('../../../../../../../gogoproto/gogo_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
goog.exportSymbol('proto.gloo.solo.io.LoadBalancerConfig', null, global);
goog.exportSymbol('proto.gloo.solo.io.LoadBalancerConfig.LeastRequest', null, global);
goog.exportSymbol('proto.gloo.solo.io.LoadBalancerConfig.Random', null, global);
goog.exportSymbol('proto.gloo.solo.io.LoadBalancerConfig.RoundRobin', null, global);

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
proto.gloo.solo.io.LoadBalancerConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.gloo.solo.io.LoadBalancerConfig.oneofGroups_);
};
goog.inherits(proto.gloo.solo.io.LoadBalancerConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.LoadBalancerConfig.displayName = 'proto.gloo.solo.io.LoadBalancerConfig';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.gloo.solo.io.LoadBalancerConfig.oneofGroups_ = [[3,4,5]];

/**
 * @enum {number}
 */
proto.gloo.solo.io.LoadBalancerConfig.TypeCase = {
  TYPE_NOT_SET: 0,
  ROUND_ROBIN: 3,
  LEAST_REQUEST: 4,
  RANDOM: 5
};

/**
 * @return {proto.gloo.solo.io.LoadBalancerConfig.TypeCase}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.getTypeCase = function() {
  return /** @type {proto.gloo.solo.io.LoadBalancerConfig.TypeCase} */(jspb.Message.computeOneofCase(this, proto.gloo.solo.io.LoadBalancerConfig.oneofGroups_[0]));
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
proto.gloo.solo.io.LoadBalancerConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.LoadBalancerConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.LoadBalancerConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.LoadBalancerConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    healthyPanicThreshold: (f = msg.getHealthyPanicThreshold()) && google_protobuf_wrappers_pb.DoubleValue.toObject(includeInstance, f),
    updateMergeWindow: (f = msg.getUpdateMergeWindow()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    roundRobin: (f = msg.getRoundRobin()) && proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.toObject(includeInstance, f),
    leastRequest: (f = msg.getLeastRequest()) && proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.toObject(includeInstance, f),
    random: (f = msg.getRandom()) && proto.gloo.solo.io.LoadBalancerConfig.Random.toObject(includeInstance, f)
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
 * @return {!proto.gloo.solo.io.LoadBalancerConfig}
 */
proto.gloo.solo.io.LoadBalancerConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.LoadBalancerConfig;
  return proto.gloo.solo.io.LoadBalancerConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.LoadBalancerConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.LoadBalancerConfig}
 */
proto.gloo.solo.io.LoadBalancerConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.DoubleValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.DoubleValue.deserializeBinaryFromReader);
      msg.setHealthyPanicThreshold(value);
      break;
    case 2:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setUpdateMergeWindow(value);
      break;
    case 3:
      var value = new proto.gloo.solo.io.LoadBalancerConfig.RoundRobin;
      reader.readMessage(value,proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.deserializeBinaryFromReader);
      msg.setRoundRobin(value);
      break;
    case 4:
      var value = new proto.gloo.solo.io.LoadBalancerConfig.LeastRequest;
      reader.readMessage(value,proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.deserializeBinaryFromReader);
      msg.setLeastRequest(value);
      break;
    case 5:
      var value = new proto.gloo.solo.io.LoadBalancerConfig.Random;
      reader.readMessage(value,proto.gloo.solo.io.LoadBalancerConfig.Random.deserializeBinaryFromReader);
      msg.setRandom(value);
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
proto.gloo.solo.io.LoadBalancerConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.LoadBalancerConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.LoadBalancerConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.LoadBalancerConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHealthyPanicThreshold();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.DoubleValue.serializeBinaryToWriter
    );
  }
  f = message.getUpdateMergeWindow();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getRoundRobin();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.serializeBinaryToWriter
    );
  }
  f = message.getLeastRequest();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.serializeBinaryToWriter
    );
  }
  f = message.getRandom();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.gloo.solo.io.LoadBalancerConfig.Random.serializeBinaryToWriter
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
proto.gloo.solo.io.LoadBalancerConfig.RoundRobin = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.LoadBalancerConfig.RoundRobin, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.displayName = 'proto.gloo.solo.io.LoadBalancerConfig.RoundRobin';
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
proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.RoundRobin} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.toObject = function(includeInstance, msg) {
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
 * @return {!proto.gloo.solo.io.LoadBalancerConfig.RoundRobin}
 */
proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.LoadBalancerConfig.RoundRobin;
  return proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.RoundRobin} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.LoadBalancerConfig.RoundRobin}
 */
proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.deserializeBinaryFromReader = function(msg, reader) {
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
proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.RoundRobin} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.LoadBalancerConfig.RoundRobin.serializeBinaryToWriter = function(message, writer) {
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
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.LoadBalancerConfig.LeastRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.displayName = 'proto.gloo.solo.io.LoadBalancerConfig.LeastRequest';
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
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.LeastRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    choiceCount: jspb.Message.getFieldWithDefault(msg, 1, 0)
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
 * @return {!proto.gloo.solo.io.LoadBalancerConfig.LeastRequest}
 */
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.LoadBalancerConfig.LeastRequest;
  return proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.LeastRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.LoadBalancerConfig.LeastRequest}
 */
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setChoiceCount(value);
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
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.LeastRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getChoiceCount();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
};


/**
 * optional uint32 choice_count = 1;
 * @return {number}
 */
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.prototype.getChoiceCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.gloo.solo.io.LoadBalancerConfig.LeastRequest.prototype.setChoiceCount = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
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
proto.gloo.solo.io.LoadBalancerConfig.Random = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.LoadBalancerConfig.Random, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.LoadBalancerConfig.Random.displayName = 'proto.gloo.solo.io.LoadBalancerConfig.Random';
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
proto.gloo.solo.io.LoadBalancerConfig.Random.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.LoadBalancerConfig.Random.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.Random} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.LoadBalancerConfig.Random.toObject = function(includeInstance, msg) {
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
 * @return {!proto.gloo.solo.io.LoadBalancerConfig.Random}
 */
proto.gloo.solo.io.LoadBalancerConfig.Random.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.LoadBalancerConfig.Random;
  return proto.gloo.solo.io.LoadBalancerConfig.Random.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.Random} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.LoadBalancerConfig.Random}
 */
proto.gloo.solo.io.LoadBalancerConfig.Random.deserializeBinaryFromReader = function(msg, reader) {
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
proto.gloo.solo.io.LoadBalancerConfig.Random.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.LoadBalancerConfig.Random.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.LoadBalancerConfig.Random} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.LoadBalancerConfig.Random.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


/**
 * optional google.protobuf.DoubleValue healthy_panic_threshold = 1;
 * @return {?proto.google.protobuf.DoubleValue}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.getHealthyPanicThreshold = function() {
  return /** @type{?proto.google.protobuf.DoubleValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.DoubleValue, 1));
};


/** @param {?proto.google.protobuf.DoubleValue|undefined} value */
proto.gloo.solo.io.LoadBalancerConfig.prototype.setHealthyPanicThreshold = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.gloo.solo.io.LoadBalancerConfig.prototype.clearHealthyPanicThreshold = function() {
  this.setHealthyPanicThreshold(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.hasHealthyPanicThreshold = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.Duration update_merge_window = 2;
 * @return {?proto.google.protobuf.Duration}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.getUpdateMergeWindow = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 2));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.gloo.solo.io.LoadBalancerConfig.prototype.setUpdateMergeWindow = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.gloo.solo.io.LoadBalancerConfig.prototype.clearUpdateMergeWindow = function() {
  this.setUpdateMergeWindow(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.hasUpdateMergeWindow = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional RoundRobin round_robin = 3;
 * @return {?proto.gloo.solo.io.LoadBalancerConfig.RoundRobin}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.getRoundRobin = function() {
  return /** @type{?proto.gloo.solo.io.LoadBalancerConfig.RoundRobin} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.LoadBalancerConfig.RoundRobin, 3));
};


/** @param {?proto.gloo.solo.io.LoadBalancerConfig.RoundRobin|undefined} value */
proto.gloo.solo.io.LoadBalancerConfig.prototype.setRoundRobin = function(value) {
  jspb.Message.setOneofWrapperField(this, 3, proto.gloo.solo.io.LoadBalancerConfig.oneofGroups_[0], value);
};


proto.gloo.solo.io.LoadBalancerConfig.prototype.clearRoundRobin = function() {
  this.setRoundRobin(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.hasRoundRobin = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional LeastRequest least_request = 4;
 * @return {?proto.gloo.solo.io.LoadBalancerConfig.LeastRequest}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.getLeastRequest = function() {
  return /** @type{?proto.gloo.solo.io.LoadBalancerConfig.LeastRequest} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.LoadBalancerConfig.LeastRequest, 4));
};


/** @param {?proto.gloo.solo.io.LoadBalancerConfig.LeastRequest|undefined} value */
proto.gloo.solo.io.LoadBalancerConfig.prototype.setLeastRequest = function(value) {
  jspb.Message.setOneofWrapperField(this, 4, proto.gloo.solo.io.LoadBalancerConfig.oneofGroups_[0], value);
};


proto.gloo.solo.io.LoadBalancerConfig.prototype.clearLeastRequest = function() {
  this.setLeastRequest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.hasLeastRequest = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional Random random = 5;
 * @return {?proto.gloo.solo.io.LoadBalancerConfig.Random}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.getRandom = function() {
  return /** @type{?proto.gloo.solo.io.LoadBalancerConfig.Random} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.LoadBalancerConfig.Random, 5));
};


/** @param {?proto.gloo.solo.io.LoadBalancerConfig.Random|undefined} value */
proto.gloo.solo.io.LoadBalancerConfig.prototype.setRandom = function(value) {
  jspb.Message.setOneofWrapperField(this, 5, proto.gloo.solo.io.LoadBalancerConfig.oneofGroups_[0], value);
};


proto.gloo.solo.io.LoadBalancerConfig.prototype.clearRandom = function() {
  this.setRandom(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.LoadBalancerConfig.prototype.hasRandom = function() {
  return jspb.Message.getField(this, 5) != null;
};


goog.object.extend(exports, proto.gloo.solo.io);
