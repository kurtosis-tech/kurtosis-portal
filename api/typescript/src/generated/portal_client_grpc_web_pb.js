/**
 * @fileoverview gRPC-Web generated client stub for kurtosis_portal_daemon
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.19.1
// source: portal_client.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')

var portal_remote_endpoint_pb = require('./portal_remote_endpoint_pb.js')

var portal_ping_pb = require('./portal_ping_pb.js')
const proto = {};
proto.kurtosis_portal_daemon = require('./portal_client_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.kurtosis_portal_daemon.KurtosisPortalClientClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.kurtosis_portal_daemon.KurtosisPortalClientPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kurtosis_portal_daemon.PortalPing,
 *   !proto.kurtosis_portal_daemon.PortalPong>}
 */
const methodDescriptor_KurtosisPortalClient_Ping = new grpc.web.MethodDescriptor(
  '/kurtosis_portal_daemon.KurtosisPortalClient/Ping',
  grpc.web.MethodType.UNARY,
  portal_ping_pb.PortalPing,
  portal_ping_pb.PortalPong,
  /**
   * @param {!proto.kurtosis_portal_daemon.PortalPing} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  portal_ping_pb.PortalPong.deserializeBinary
);


/**
 * @param {!proto.kurtosis_portal_daemon.PortalPing} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.kurtosis_portal_daemon.PortalPong)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kurtosis_portal_daemon.PortalPong>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kurtosis_portal_daemon.KurtosisPortalClientClient.prototype.ping =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kurtosis_portal_daemon.KurtosisPortalClient/Ping',
      request,
      metadata || {},
      methodDescriptor_KurtosisPortalClient_Ping,
      callback);
};


/**
 * @param {!proto.kurtosis_portal_daemon.PortalPing} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kurtosis_portal_daemon.PortalPong>}
 *     Promise that resolves to the response
 */
proto.kurtosis_portal_daemon.KurtosisPortalClientPromiseClient.prototype.ping =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kurtosis_portal_daemon.KurtosisPortalClient/Ping',
      request,
      metadata || {},
      methodDescriptor_KurtosisPortalClient_Ping);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kurtosis_portal_daemon.SwitchContextArgs,
 *   !proto.kurtosis_portal_daemon.SwitchContextResponse>}
 */
const methodDescriptor_KurtosisPortalClient_SwitchContext = new grpc.web.MethodDescriptor(
  '/kurtosis_portal_daemon.KurtosisPortalClient/SwitchContext',
  grpc.web.MethodType.UNARY,
  proto.kurtosis_portal_daemon.SwitchContextArgs,
  proto.kurtosis_portal_daemon.SwitchContextResponse,
  /**
   * @param {!proto.kurtosis_portal_daemon.SwitchContextArgs} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kurtosis_portal_daemon.SwitchContextResponse.deserializeBinary
);


/**
 * @param {!proto.kurtosis_portal_daemon.SwitchContextArgs} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.kurtosis_portal_daemon.SwitchContextResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kurtosis_portal_daemon.SwitchContextResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kurtosis_portal_daemon.KurtosisPortalClientClient.prototype.switchContext =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kurtosis_portal_daemon.KurtosisPortalClient/SwitchContext',
      request,
      metadata || {},
      methodDescriptor_KurtosisPortalClient_SwitchContext,
      callback);
};


/**
 * @param {!proto.kurtosis_portal_daemon.SwitchContextArgs} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kurtosis_portal_daemon.SwitchContextResponse>}
 *     Promise that resolves to the response
 */
proto.kurtosis_portal_daemon.KurtosisPortalClientPromiseClient.prototype.switchContext =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kurtosis_portal_daemon.KurtosisPortalClient/SwitchContext',
      request,
      metadata || {},
      methodDescriptor_KurtosisPortalClient_SwitchContext);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.kurtosis_portal_daemon.ForwardPortArgs,
 *   !proto.kurtosis_portal_daemon.ForwardPortResponse>}
 */
const methodDescriptor_KurtosisPortalClient_ForwardPort = new grpc.web.MethodDescriptor(
  '/kurtosis_portal_daemon.KurtosisPortalClient/ForwardPort',
  grpc.web.MethodType.UNARY,
  proto.kurtosis_portal_daemon.ForwardPortArgs,
  proto.kurtosis_portal_daemon.ForwardPortResponse,
  /**
   * @param {!proto.kurtosis_portal_daemon.ForwardPortArgs} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.kurtosis_portal_daemon.ForwardPortResponse.deserializeBinary
);


/**
 * @param {!proto.kurtosis_portal_daemon.ForwardPortArgs} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.kurtosis_portal_daemon.ForwardPortResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.kurtosis_portal_daemon.ForwardPortResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.kurtosis_portal_daemon.KurtosisPortalClientClient.prototype.forwardPort =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/kurtosis_portal_daemon.KurtosisPortalClient/ForwardPort',
      request,
      metadata || {},
      methodDescriptor_KurtosisPortalClient_ForwardPort,
      callback);
};


/**
 * @param {!proto.kurtosis_portal_daemon.ForwardPortArgs} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.kurtosis_portal_daemon.ForwardPortResponse>}
 *     Promise that resolves to the response
 */
proto.kurtosis_portal_daemon.KurtosisPortalClientPromiseClient.prototype.forwardPort =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/kurtosis_portal_daemon.KurtosisPortalClient/ForwardPort',
      request,
      metadata || {},
      methodDescriptor_KurtosisPortalClient_ForwardPort);
};


module.exports = proto.kurtosis_portal_daemon;

