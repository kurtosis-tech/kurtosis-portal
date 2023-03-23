import * as grpcWeb from 'grpc-web';

import * as portal_client_pb from './portal_client_pb';
import * as portal_ping_pb from './portal_ping_pb';


export class KurtosisPortalClientClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  ping(
    request: portal_ping_pb.PortalPing,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: portal_ping_pb.PortalPong) => void
  ): grpcWeb.ClientReadableStream<portal_ping_pb.PortalPong>;

  switchContext(
    request: portal_client_pb.SwitchContextArgs,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: portal_client_pb.SwitchContextResponse) => void
  ): grpcWeb.ClientReadableStream<portal_client_pb.SwitchContextResponse>;

  forwardPort(
    request: portal_client_pb.ForwardPortArgs,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: portal_client_pb.ForwardPortResponse) => void
  ): grpcWeb.ClientReadableStream<portal_client_pb.ForwardPortResponse>;

}

export class KurtosisPortalClientPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  ping(
    request: portal_ping_pb.PortalPing,
    metadata?: grpcWeb.Metadata
  ): Promise<portal_ping_pb.PortalPong>;

  switchContext(
    request: portal_client_pb.SwitchContextArgs,
    metadata?: grpcWeb.Metadata
  ): Promise<portal_client_pb.SwitchContextResponse>;

  forwardPort(
    request: portal_client_pb.ForwardPortArgs,
    metadata?: grpcWeb.Metadata
  ): Promise<portal_client_pb.ForwardPortResponse>;

}

