// GENERATED CODE -- DO NOT EDIT!

// package: kurtosis_portal_daemon
// file: portal_client.proto

import * as portal_client_pb from "./portal_client_pb";
import * as portal_ping_pb from "./portal_ping_pb";
import * as grpc from "@grpc/grpc-js";

interface IKurtosisPortalClientService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  ping: grpc.MethodDefinition<portal_ping_pb.PortalPing, portal_ping_pb.PortalPong>;
  switchContext: grpc.MethodDefinition<portal_client_pb.SwitchContextArgs, portal_client_pb.SwitchContextResponse>;
  forwardPort: grpc.MethodDefinition<portal_client_pb.ForwardPortArgs, portal_client_pb.ForwardPortResponse>;
}

export const KurtosisPortalClientService: IKurtosisPortalClientService;

export interface IKurtosisPortalClientServer extends grpc.UntypedServiceImplementation {
  ping: grpc.handleUnaryCall<portal_ping_pb.PortalPing, portal_ping_pb.PortalPong>;
  switchContext: grpc.handleUnaryCall<portal_client_pb.SwitchContextArgs, portal_client_pb.SwitchContextResponse>;
  forwardPort: grpc.handleUnaryCall<portal_client_pb.ForwardPortArgs, portal_client_pb.ForwardPortResponse>;
}

export class KurtosisPortalClientClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  ping(argument: portal_ping_pb.PortalPing, callback: grpc.requestCallback<portal_ping_pb.PortalPong>): grpc.ClientUnaryCall;
  ping(argument: portal_ping_pb.PortalPing, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<portal_ping_pb.PortalPong>): grpc.ClientUnaryCall;
  ping(argument: portal_ping_pb.PortalPing, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<portal_ping_pb.PortalPong>): grpc.ClientUnaryCall;
  switchContext(argument: portal_client_pb.SwitchContextArgs, callback: grpc.requestCallback<portal_client_pb.SwitchContextResponse>): grpc.ClientUnaryCall;
  switchContext(argument: portal_client_pb.SwitchContextArgs, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<portal_client_pb.SwitchContextResponse>): grpc.ClientUnaryCall;
  switchContext(argument: portal_client_pb.SwitchContextArgs, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<portal_client_pb.SwitchContextResponse>): grpc.ClientUnaryCall;
  forwardPort(argument: portal_client_pb.ForwardPortArgs, callback: grpc.requestCallback<portal_client_pb.ForwardPortResponse>): grpc.ClientUnaryCall;
  forwardPort(argument: portal_client_pb.ForwardPortArgs, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<portal_client_pb.ForwardPortResponse>): grpc.ClientUnaryCall;
  forwardPort(argument: portal_client_pb.ForwardPortArgs, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<portal_client_pb.ForwardPortResponse>): grpc.ClientUnaryCall;
}
