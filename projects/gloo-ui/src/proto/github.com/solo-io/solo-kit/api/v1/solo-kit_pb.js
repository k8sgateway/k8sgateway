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

var google_protobuf_descriptor_pb = require('google-protobuf/google/protobuf/descriptor_pb.js');
goog.exportSymbol('proto.core.solo.io.Resource', null, global);
goog.exportSymbol('proto.core.solo.io.resource', null, global);
goog.exportSymbol('proto.core.solo.io.skipHashing', null, global);

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
proto.core.solo.io.Resource = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.solo.io.Resource, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.core.solo.io.Resource.displayName = 'proto.core.solo.io.Resource';
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
proto.core.solo.io.Resource.prototype.toObject = function(opt_includeInstance) {
  return proto.core.solo.io.Resource.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.solo.io.Resource} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.Resource.toObject = function(includeInstance, msg) {
  var f, obj = {
    shortName: jspb.Message.getFieldWithDefault(msg, 1, ""),
    pluralName: jspb.Message.getFieldWithDefault(msg, 2, ""),
    clusterScoped: jspb.Message.getFieldWithDefault(msg, 3, false),
    skipDocsGen: jspb.Message.getFieldWithDefault(msg, 4, false)
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
 * @return {!proto.core.solo.io.Resource}
 */
proto.core.solo.io.Resource.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.solo.io.Resource;
  return proto.core.solo.io.Resource.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.solo.io.Resource} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.solo.io.Resource}
 */
proto.core.solo.io.Resource.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setShortName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPluralName(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setClusterScoped(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSkipDocsGen(value);
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
proto.core.solo.io.Resource.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.solo.io.Resource.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.solo.io.Resource} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.Resource.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getShortName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPluralName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getClusterScoped();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getSkipDocsGen();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
};


/**
 * optional string short_name = 1;
 * @return {string}
 */
proto.core.solo.io.Resource.prototype.getShortName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.core.solo.io.Resource.prototype.setShortName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string plural_name = 2;
 * @return {string}
 */
proto.core.solo.io.Resource.prototype.getPluralName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.core.solo.io.Resource.prototype.setPluralName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bool cluster_scoped = 3;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.core.solo.io.Resource.prototype.getClusterScoped = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 3, false));
};


/** @param {boolean} value */
proto.core.solo.io.Resource.prototype.setClusterScoped = function(value) {
  jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional bool skip_docs_gen = 4;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.core.solo.io.Resource.prototype.getSkipDocsGen = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 4, false));
};


/** @param {boolean} value */
proto.core.solo.io.Resource.prototype.setSkipDocsGen = function(value) {
  jspb.Message.setProto3BooleanField(this, 4, value);
};



/**
 * A tuple of {field number, class constructor} for the extension
 * field named `resource`.
 * @type {!jspb.ExtensionFieldInfo<!proto.core.solo.io.Resource>}
 */
proto.core.solo.io.resource = new jspb.ExtensionFieldInfo(
    10000,
    {resource: 0},
    proto.core.solo.io.Resource,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         proto.core.solo.io.Resource.toObject),
    0);

google_protobuf_descriptor_pb.MessageOptions.extensionsBinary[10000] = new jspb.ExtensionFieldBinaryInfo(
    proto.core.solo.io.resource,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    proto.core.solo.io.Resource.serializeBinaryToWriter,
    proto.core.solo.io.Resource.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.MessageOptions.extensions[10000] = proto.core.solo.io.resource;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `skipHashing`.
 * @type {!jspb.ExtensionFieldInfo<boolean>}
 */
proto.core.solo.io.skipHashing = new jspb.ExtensionFieldInfo(
    10000,
    {skipHashing: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[10000] = new jspb.ExtensionFieldBinaryInfo(
    proto.core.solo.io.skipHashing,
    jspb.BinaryReader.prototype.readBool,
    jspb.BinaryWriter.prototype.writeBool,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[10000] = proto.core.solo.io.skipHashing;

goog.object.extend(exports, proto.core.solo.io);
