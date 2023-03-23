// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var portal_ping_pb = require('./portal_ping_pb.js');

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
};

exports.KurtosisPortalServerClient = grpc.makeGenericClientConstructor(KurtosisPortalServerService);
