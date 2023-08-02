// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var portal_client_pb = require('./portal_client_pb.js');
var google_api_annotations_pb = require('./google/api/annotations_pb.js');
var portal_remote_endpoint_pb = require('./portal_remote_endpoint_pb.js');
var portal_ping_pb = require('./portal_ping_pb.js');

function serialize_kurtosis_portal_daemon_ForwardPortArgs(arg) {
  if (!(arg instanceof portal_client_pb.ForwardPortArgs)) {
    throw new Error('Expected argument of type kurtosis_portal_daemon.ForwardPortArgs');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_portal_daemon_ForwardPortArgs(buffer_arg) {
  return portal_client_pb.ForwardPortArgs.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kurtosis_portal_daemon_ForwardPortResponse(arg) {
  if (!(arg instanceof portal_client_pb.ForwardPortResponse)) {
    throw new Error('Expected argument of type kurtosis_portal_daemon.ForwardPortResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_portal_daemon_ForwardPortResponse(buffer_arg) {
  return portal_client_pb.ForwardPortResponse.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_kurtosis_portal_daemon_SwitchContextArgs(arg) {
  if (!(arg instanceof portal_client_pb.SwitchContextArgs)) {
    throw new Error('Expected argument of type kurtosis_portal_daemon.SwitchContextArgs');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_portal_daemon_SwitchContextArgs(buffer_arg) {
  return portal_client_pb.SwitchContextArgs.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kurtosis_portal_daemon_SwitchContextResponse(arg) {
  if (!(arg instanceof portal_client_pb.SwitchContextResponse)) {
    throw new Error('Expected argument of type kurtosis_portal_daemon.SwitchContextResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_portal_daemon_SwitchContextResponse(buffer_arg) {
  return portal_client_pb.SwitchContextResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var KurtosisPortalClientService = exports.KurtosisPortalClientService = {
  // To check availability
ping: {
    path: '/kurtosis_portal_daemon.KurtosisPortalClient/Ping',
    requestStream: false,
    responseStream: false,
    requestType: portal_ping_pb.PortalPing,
    responseType: portal_ping_pb.PortalPong,
    requestSerialize: serialize_kurtosis_portal_daemon_PortalPing,
    requestDeserialize: deserialize_kurtosis_portal_daemon_PortalPing,
    responseSerialize: serialize_kurtosis_portal_daemon_PortalPong,
    responseDeserialize: deserialize_kurtosis_portal_daemon_PortalPong,
  },
  // SwitchContext switches the current context used by Kurtosis.
//
// If the new context is a dual-backend-context, it connects to it automatically using the connection information
// linked to the context Right now, it is expected that the remote environment is running a Kurtosis Portal Server
// on port 9720
switchContext: {
    path: '/kurtosis_portal_daemon.KurtosisPortalClient/SwitchContext',
    requestStream: false,
    responseStream: false,
    requestType: portal_client_pb.SwitchContextArgs,
    responseType: portal_client_pb.SwitchContextResponse,
    requestSerialize: serialize_kurtosis_portal_daemon_SwitchContextArgs,
    requestDeserialize: deserialize_kurtosis_portal_daemon_SwitchContextArgs,
    responseSerialize: serialize_kurtosis_portal_daemon_SwitchContextResponse,
    responseDeserialize: deserialize_kurtosis_portal_daemon_SwitchContextResponse,
  },
  // TODO: Raw endpoint to forward a port from server to client. This is very low level, in the future the portal
//  should accept higher level info, like (enclave_id, service_id, port_id) and automatically find the ephemeral
//  port number.
forwardPort: {
    path: '/kurtosis_portal_daemon.KurtosisPortalClient/ForwardPort',
    requestStream: false,
    responseStream: false,
    requestType: portal_client_pb.ForwardPortArgs,
    responseType: portal_client_pb.ForwardPortResponse,
    requestSerialize: serialize_kurtosis_portal_daemon_ForwardPortArgs,
    requestDeserialize: deserialize_kurtosis_portal_daemon_ForwardPortArgs,
    responseSerialize: serialize_kurtosis_portal_daemon_ForwardPortResponse,
    responseDeserialize: deserialize_kurtosis_portal_daemon_ForwardPortResponse,
  },
};

exports.KurtosisPortalClientClient = grpc.makeGenericClientConstructor(KurtosisPortalClientService);
