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
var github_com_solo$io_gloo_projects_gloo_api_v1_extensions_pb = require('../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/extensions_pb.js');
var github_com_solo$io_solo$kit_api_v1_metadata_pb = require('../../../../../../../github.com/solo-io/solo-kit/api/v1/metadata_pb.js');
goog.exportSymbol('proto.gloo.solo.io.AwsSecret', null, global);
goog.exportSymbol('proto.gloo.solo.io.AzureSecret', null, global);
goog.exportSymbol('proto.gloo.solo.io.Secret', null, global);
goog.exportSymbol('proto.gloo.solo.io.TlsSecret', null, global);

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
proto.gloo.solo.io.Secret = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.gloo.solo.io.Secret.oneofGroups_);
};
goog.inherits(proto.gloo.solo.io.Secret, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.Secret.displayName = 'proto.gloo.solo.io.Secret';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.gloo.solo.io.Secret.oneofGroups_ = [[1,2,3,4]];

/**
 * @enum {number}
 */
proto.gloo.solo.io.Secret.KindCase = {
  KIND_NOT_SET: 0,
  AWS: 1,
  AZURE: 2,
  TLS: 3,
  EXTENSION: 4
};

/**
 * @return {proto.gloo.solo.io.Secret.KindCase}
 */
proto.gloo.solo.io.Secret.prototype.getKindCase = function() {
  return /** @type {proto.gloo.solo.io.Secret.KindCase} */(jspb.Message.computeOneofCase(this, proto.gloo.solo.io.Secret.oneofGroups_[0]));
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
proto.gloo.solo.io.Secret.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.Secret.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.Secret} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.Secret.toObject = function(includeInstance, msg) {
  var f, obj = {
    aws: (f = msg.getAws()) && proto.gloo.solo.io.AwsSecret.toObject(includeInstance, f),
    azure: (f = msg.getAzure()) && proto.gloo.solo.io.AzureSecret.toObject(includeInstance, f),
    tls: (f = msg.getTls()) && proto.gloo.solo.io.TlsSecret.toObject(includeInstance, f),
    extension: (f = msg.getExtension$()) && github_com_solo$io_gloo_projects_gloo_api_v1_extensions_pb.Extension.toObject(includeInstance, f),
    metadata: (f = msg.getMetadata()) && github_com_solo$io_solo$kit_api_v1_metadata_pb.Metadata.toObject(includeInstance, f)
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
 * @return {!proto.gloo.solo.io.Secret}
 */
proto.gloo.solo.io.Secret.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.Secret;
  return proto.gloo.solo.io.Secret.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.Secret} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.Secret}
 */
proto.gloo.solo.io.Secret.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.gloo.solo.io.AwsSecret;
      reader.readMessage(value,proto.gloo.solo.io.AwsSecret.deserializeBinaryFromReader);
      msg.setAws(value);
      break;
    case 2:
      var value = new proto.gloo.solo.io.AzureSecret;
      reader.readMessage(value,proto.gloo.solo.io.AzureSecret.deserializeBinaryFromReader);
      msg.setAzure(value);
      break;
    case 3:
      var value = new proto.gloo.solo.io.TlsSecret;
      reader.readMessage(value,proto.gloo.solo.io.TlsSecret.deserializeBinaryFromReader);
      msg.setTls(value);
      break;
    case 4:
      var value = new github_com_solo$io_gloo_projects_gloo_api_v1_extensions_pb.Extension;
      reader.readMessage(value,github_com_solo$io_gloo_projects_gloo_api_v1_extensions_pb.Extension.deserializeBinaryFromReader);
      msg.setExtension$(value);
      break;
    case 7:
      var value = new github_com_solo$io_solo$kit_api_v1_metadata_pb.Metadata;
      reader.readMessage(value,github_com_solo$io_solo$kit_api_v1_metadata_pb.Metadata.deserializeBinaryFromReader);
      msg.setMetadata(value);
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
proto.gloo.solo.io.Secret.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.Secret.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.Secret} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.Secret.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAws();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.gloo.solo.io.AwsSecret.serializeBinaryToWriter
    );
  }
  f = message.getAzure();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.gloo.solo.io.AzureSecret.serializeBinaryToWriter
    );
  }
  f = message.getTls();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.gloo.solo.io.TlsSecret.serializeBinaryToWriter
    );
  }
  f = message.getExtension$();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      github_com_solo$io_gloo_projects_gloo_api_v1_extensions_pb.Extension.serializeBinaryToWriter
    );
  }
  f = message.getMetadata();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      github_com_solo$io_solo$kit_api_v1_metadata_pb.Metadata.serializeBinaryToWriter
    );
  }
};


/**
 * optional AwsSecret aws = 1;
 * @return {?proto.gloo.solo.io.AwsSecret}
 */
proto.gloo.solo.io.Secret.prototype.getAws = function() {
  return /** @type{?proto.gloo.solo.io.AwsSecret} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.AwsSecret, 1));
};


/** @param {?proto.gloo.solo.io.AwsSecret|undefined} value */
proto.gloo.solo.io.Secret.prototype.setAws = function(value) {
  jspb.Message.setOneofWrapperField(this, 1, proto.gloo.solo.io.Secret.oneofGroups_[0], value);
};


proto.gloo.solo.io.Secret.prototype.clearAws = function() {
  this.setAws(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.Secret.prototype.hasAws = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional AzureSecret azure = 2;
 * @return {?proto.gloo.solo.io.AzureSecret}
 */
proto.gloo.solo.io.Secret.prototype.getAzure = function() {
  return /** @type{?proto.gloo.solo.io.AzureSecret} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.AzureSecret, 2));
};


/** @param {?proto.gloo.solo.io.AzureSecret|undefined} value */
proto.gloo.solo.io.Secret.prototype.setAzure = function(value) {
  jspb.Message.setOneofWrapperField(this, 2, proto.gloo.solo.io.Secret.oneofGroups_[0], value);
};


proto.gloo.solo.io.Secret.prototype.clearAzure = function() {
  this.setAzure(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.Secret.prototype.hasAzure = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional TlsSecret tls = 3;
 * @return {?proto.gloo.solo.io.TlsSecret}
 */
proto.gloo.solo.io.Secret.prototype.getTls = function() {
  return /** @type{?proto.gloo.solo.io.TlsSecret} */ (
    jspb.Message.getWrapperField(this, proto.gloo.solo.io.TlsSecret, 3));
};


/** @param {?proto.gloo.solo.io.TlsSecret|undefined} value */
proto.gloo.solo.io.Secret.prototype.setTls = function(value) {
  jspb.Message.setOneofWrapperField(this, 3, proto.gloo.solo.io.Secret.oneofGroups_[0], value);
};


proto.gloo.solo.io.Secret.prototype.clearTls = function() {
  this.setTls(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.Secret.prototype.hasTls = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional Extension extension = 4;
 * @return {?proto.gloo.solo.io.Extension}
 */
proto.gloo.solo.io.Secret.prototype.getExtension$ = function() {
  return /** @type{?proto.gloo.solo.io.Extension} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_gloo_projects_gloo_api_v1_extensions_pb.Extension, 4));
};


/** @param {?proto.gloo.solo.io.Extension|undefined} value */
proto.gloo.solo.io.Secret.prototype.setExtension$ = function(value) {
  jspb.Message.setOneofWrapperField(this, 4, proto.gloo.solo.io.Secret.oneofGroups_[0], value);
};


proto.gloo.solo.io.Secret.prototype.clearExtension$ = function() {
  this.setExtension$(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.Secret.prototype.hasExtension$ = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional core.solo.io.Metadata metadata = 7;
 * @return {?proto.core.solo.io.Metadata}
 */
proto.gloo.solo.io.Secret.prototype.getMetadata = function() {
  return /** @type{?proto.core.solo.io.Metadata} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$kit_api_v1_metadata_pb.Metadata, 7));
};


/** @param {?proto.core.solo.io.Metadata|undefined} value */
proto.gloo.solo.io.Secret.prototype.setMetadata = function(value) {
  jspb.Message.setWrapperField(this, 7, value);
};


proto.gloo.solo.io.Secret.prototype.clearMetadata = function() {
  this.setMetadata(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.gloo.solo.io.Secret.prototype.hasMetadata = function() {
  return jspb.Message.getField(this, 7) != null;
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
proto.gloo.solo.io.AwsSecret = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.AwsSecret, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.AwsSecret.displayName = 'proto.gloo.solo.io.AwsSecret';
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
proto.gloo.solo.io.AwsSecret.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.AwsSecret.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.AwsSecret} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.AwsSecret.toObject = function(includeInstance, msg) {
  var f, obj = {
    accessKey: jspb.Message.getFieldWithDefault(msg, 1, ""),
    secretKey: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.gloo.solo.io.AwsSecret}
 */
proto.gloo.solo.io.AwsSecret.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.AwsSecret;
  return proto.gloo.solo.io.AwsSecret.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.AwsSecret} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.AwsSecret}
 */
proto.gloo.solo.io.AwsSecret.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccessKey(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setSecretKey(value);
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
proto.gloo.solo.io.AwsSecret.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.AwsSecret.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.AwsSecret} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.AwsSecret.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAccessKey();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSecretKey();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string access_key = 1;
 * @return {string}
 */
proto.gloo.solo.io.AwsSecret.prototype.getAccessKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.AwsSecret.prototype.setAccessKey = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string secret_key = 2;
 * @return {string}
 */
proto.gloo.solo.io.AwsSecret.prototype.getSecretKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.gloo.solo.io.AwsSecret.prototype.setSecretKey = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
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
proto.gloo.solo.io.AzureSecret = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.AzureSecret, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.AzureSecret.displayName = 'proto.gloo.solo.io.AzureSecret';
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
proto.gloo.solo.io.AzureSecret.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.AzureSecret.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.AzureSecret} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.AzureSecret.toObject = function(includeInstance, msg) {
  var f, obj = {
    apiKeysMap: (f = msg.getApiKeysMap()) ? f.toObject(includeInstance, undefined) : []
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
 * @return {!proto.gloo.solo.io.AzureSecret}
 */
proto.gloo.solo.io.AzureSecret.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.AzureSecret;
  return proto.gloo.solo.io.AzureSecret.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.AzureSecret} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.AzureSecret}
 */
proto.gloo.solo.io.AzureSecret.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = msg.getApiKeysMap();
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
proto.gloo.solo.io.AzureSecret.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.AzureSecret.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.AzureSecret} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.AzureSecret.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApiKeysMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(1, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * map<string, string> api_keys = 1;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.gloo.solo.io.AzureSecret.prototype.getApiKeysMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 1, opt_noLazyCreate,
      null));
};


proto.gloo.solo.io.AzureSecret.prototype.clearApiKeysMap = function() {
  this.getApiKeysMap().clear();
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
proto.gloo.solo.io.TlsSecret = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.gloo.solo.io.TlsSecret, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.gloo.solo.io.TlsSecret.displayName = 'proto.gloo.solo.io.TlsSecret';
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
proto.gloo.solo.io.TlsSecret.prototype.toObject = function(opt_includeInstance) {
  return proto.gloo.solo.io.TlsSecret.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.gloo.solo.io.TlsSecret} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.TlsSecret.toObject = function(includeInstance, msg) {
  var f, obj = {
    certChain: jspb.Message.getFieldWithDefault(msg, 1, ""),
    privateKey: jspb.Message.getFieldWithDefault(msg, 2, ""),
    rootCa: jspb.Message.getFieldWithDefault(msg, 3, "")
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
 * @return {!proto.gloo.solo.io.TlsSecret}
 */
proto.gloo.solo.io.TlsSecret.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.gloo.solo.io.TlsSecret;
  return proto.gloo.solo.io.TlsSecret.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.gloo.solo.io.TlsSecret} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.gloo.solo.io.TlsSecret}
 */
proto.gloo.solo.io.TlsSecret.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setCertChain(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPrivateKey(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRootCa(value);
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
proto.gloo.solo.io.TlsSecret.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.gloo.solo.io.TlsSecret.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.gloo.solo.io.TlsSecret} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.gloo.solo.io.TlsSecret.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCertChain();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPrivateKey();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRootCa();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string cert_chain = 1;
 * @return {string}
 */
proto.gloo.solo.io.TlsSecret.prototype.getCertChain = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.gloo.solo.io.TlsSecret.prototype.setCertChain = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string private_key = 2;
 * @return {string}
 */
proto.gloo.solo.io.TlsSecret.prototype.getPrivateKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.gloo.solo.io.TlsSecret.prototype.setPrivateKey = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string root_ca = 3;
 * @return {string}
 */
proto.gloo.solo.io.TlsSecret.prototype.getRootCa = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.gloo.solo.io.TlsSecret.prototype.setRootCa = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


goog.object.extend(exports, proto.gloo.solo.io);
