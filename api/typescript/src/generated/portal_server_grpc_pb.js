// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var portal_server_pb = require('./portal_server_pb.js');
var google_api_annotations_pb = require('./google/api/annotations_pb.js');
var portal_remote_endpoint_pb = require('./portal_remote_endpoint_pb.js');
var portal_ping_pb = require('./portal_ping_pb.js');

function serialize_kurtosis_portal_daemon_GetRemoteEndpointsArgs(arg) {
  if (!(arg instanceof portal_server_pb.GetRemoteEndpointsArgs)) {
    throw new Error('Expected argument of type kurtosis_portal_daemon.GetRemoteEndpointsArgs');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_portal_daemon_GetRemoteEndpointsArgs(buffer_arg) {
  return portal_server_pb.GetRemoteEndpointsArgs.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kurtosis_portal_daemon_GetRemoteEndpointsResponse(arg) {
  if (!(arg instanceof portal_server_pb.GetRemoteEndpointsResponse)) {
    throw new Error('Expected argument of type kurtosis_portal_daemon.GetRemoteEndpointsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_portal_daemon_GetRemoteEndpointsResponse(buffer_arg) {
  return portal_server_pb.GetRemoteEndpointsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kurtosis_portal_daemon_PortalPing(arg) {
  if (!(arg instanceof portal_ping_pb.PortalPing)) {
    throw new Error('Expected argument of type kurtosis_portal_daemon.PortalPing');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_portal_daemon_PortalPing(buffer_arg) {
  return portal_ping_pb.PortalPing.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kurtosis_portal_daemon_PortalPong(arg) {
  if (!(arg instanceof portal_ping_pb.PortalPong)) {
    throw new Error('Expected argument of type kurtosis_portal_daemon.PortalPong');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_portal_daemon_PortalPong(buffer_arg) {
  return portal_ping_pb.PortalPong.deserializeBinary(new Uint8Array(buffer_arg));
}


var KurtosisPortalServerService = exports.KurtosisPortalServerService = {
  // To check availability
ping: {
    path: '/kurtosis_portal_daemon.KurtosisPortalServer/Ping',
    requestStream: false,
    responseStream: false,
    requestType: portal_ping_pb.PortalPing,
    responseType: portal_ping_pb.PortalPong,
    requestSerialize: serialize_kurtosis_portal_daemon_PortalPing,
    requestDeserialize: deserialize_kurtosis_portal_daemon_PortalPing,
    responseSerialize: serialize_kurtosis_portal_daemon_PortalPong,
    responseDeserialize: deserialize_kurtosis_portal_daemon_PortalPong,
  },
  getRemoteEndpoints: {
    path: '/kurtosis_portal_daemon.KurtosisPortalServer/GetRemoteEndpoints',
    requestStream: false,
    responseStream: false,
    requestType: portal_server_pb.GetRemoteEndpointsArgs,
    responseType: portal_server_pb.GetRemoteEndpointsResponse,
    requestSerialize: serialize_kurtosis_portal_daemon_GetRemoteEndpointsArgs,
    requestDeserialize: deserialize_kurtosis_portal_daemon_GetRemoteEndpointsArgs,
    responseSerialize: serialize_kurtosis_portal_daemon_GetRemoteEndpointsResponse,
    responseDeserialize: deserialize_kurtosis_portal_daemon_GetRemoteEndpointsResponse,
  },
};

exports.KurtosisPortalServerClient = grpc.makeGenericClientConstructor(KurtosisPortalServerService);
