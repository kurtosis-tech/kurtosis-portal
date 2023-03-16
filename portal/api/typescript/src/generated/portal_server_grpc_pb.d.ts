// GENERATED CODE -- DO NOT EDIT!

// package: kurtosis_portal_daemon
// file: portal_server.proto

import * as portal_server_pb from "./portal_server_pb";
import * as portal_ping_pb from "./portal_ping_pb";
import * as grpc from "@grpc/grpc-js";

interface IKurtosisPortalServerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  ping: grpc.MethodDefinition<portal_ping_pb.PortalPing, portal_ping_pb.PortalPong>;
}

export const KurtosisPortalServerService: IKurtosisPortalServerService;

export interface IKurtosisPortalServerServer extends grpc.UntypedServiceImplementation {
  ping: grpc.handleUnaryCall<portal_ping_pb.PortalPing, portal_ping_pb.PortalPong>;
}

export class KurtosisPortalServerClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  ping(argument: portal_ping_pb.PortalPing, callback: grpc.requestCallback<portal_ping_pb.PortalPong>): grpc.ClientUnaryCall;
  ping(argument: portal_ping_pb.PortalPing, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<portal_ping_pb.PortalPong>): grpc.ClientUnaryCall;
  ping(argument: portal_ping_pb.PortalPing, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<portal_ping_pb.PortalPong>): grpc.ClientUnaryCall;
}
