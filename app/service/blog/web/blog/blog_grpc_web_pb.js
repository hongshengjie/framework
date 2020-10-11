/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = require('./blog_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.BlogClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.BlogPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.IndexResp>}
 */
const methodDescriptor_Blog_Index = new grpc.web.MethodDescriptor(
  '/Blog/Index',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.IndexResp,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.IndexResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.IndexResp>}
 */
const methodInfo_Blog_Index = new grpc.web.AbstractClientBase.MethodInfo(
  proto.IndexResp,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.IndexResp.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.IndexResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.IndexResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BlogClient.prototype.index =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Blog/Index',
      request,
      metadata || {},
      methodDescriptor_Blog_Index,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.IndexResp>}
 *     Promise that resolves to the response
 */
proto.BlogPromiseClient.prototype.index =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Blog/Index',
      request,
      metadata || {},
      methodDescriptor_Blog_Index);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ListReq,
 *   !proto.ListResp>}
 */
const methodDescriptor_Blog_List = new grpc.web.MethodDescriptor(
  '/Blog/List',
  grpc.web.MethodType.UNARY,
  proto.ListReq,
  proto.ListResp,
  /**
   * @param {!proto.ListReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ListResp.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ListReq,
 *   !proto.ListResp>}
 */
const methodInfo_Blog_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ListResp,
  /**
   * @param {!proto.ListReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ListResp.deserializeBinary
);


/**
 * @param {!proto.ListReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ListResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ListResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BlogClient.prototype.list =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Blog/List',
      request,
      metadata || {},
      methodDescriptor_Blog_List,
      callback);
};


/**
 * @param {!proto.ListReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ListResp>}
 *     Promise that resolves to the response
 */
proto.BlogPromiseClient.prototype.list =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Blog/List',
      request,
      metadata || {},
      methodDescriptor_Blog_List);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.DetailReq,
 *   !proto.Post>}
 */
const methodDescriptor_Blog_Detail = new grpc.web.MethodDescriptor(
  '/Blog/Detail',
  grpc.web.MethodType.UNARY,
  proto.DetailReq,
  proto.Post,
  /**
   * @param {!proto.DetailReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Post.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DetailReq,
 *   !proto.Post>}
 */
const methodInfo_Blog_Detail = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Post,
  /**
   * @param {!proto.DetailReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Post.deserializeBinary
);


/**
 * @param {!proto.DetailReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Post)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Post>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BlogClient.prototype.detail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Blog/Detail',
      request,
      metadata || {},
      methodDescriptor_Blog_Detail,
      callback);
};


/**
 * @param {!proto.DetailReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Post>}
 *     Promise that resolves to the response
 */
proto.BlogPromiseClient.prototype.detail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Blog/Detail',
      request,
      metadata || {},
      methodDescriptor_Blog_Detail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Post,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_Blog_Edit = new grpc.web.MethodDescriptor(
  '/Blog/Edit',
  grpc.web.MethodType.UNARY,
  proto.Post,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.Post} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Post,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_Blog_Edit = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.Post} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.Post} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.BlogClient.prototype.edit =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Blog/Edit',
      request,
      metadata || {},
      methodDescriptor_Blog_Edit,
      callback);
};


/**
 * @param {!proto.Post} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.BlogPromiseClient.prototype.edit =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Blog/Edit',
      request,
      metadata || {},
      methodDescriptor_Blog_Edit);
};


module.exports = proto;

