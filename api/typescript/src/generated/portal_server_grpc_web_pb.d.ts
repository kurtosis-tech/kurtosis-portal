import * as grpcWeb from 'grpc-web';

import * as portal_server_pb from './portal_server_pb';
import * as portal_ping_pb from './portal_ping_pb';


export class KurtosisPortalServerClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  ping(
    request: portal_ping_pb.PortalPing,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: portal_ping_pb.PortalPong) => void
  ): grpcWeb.ClientReadableStream<portal_ping_pb.PortalPong>;

  getRemoteEndpoints(
    request: portal_server_pb.GetRemoteEndpointsArgs,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: portal_server_pb.GetRemoteEndpointsResponse) => void
  ): grpcWeb.ClientReadableStream<portal_server_pb.GetRemoteEndpointsResponse>;

}

export class KurtosisPortalServerPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  ping(
    request: portal_ping_pb.PortalPing,
    metadata?: grpcWeb.Metadata
  ): Promise<portal_ping_pb.PortalPong>;

  getRemoteEndpoints(
    request: portal_server_pb.GetRemoteEndpointsArgs,
    metadata?: grpcWeb.Metadata
  ): Promise<portal_server_pb.GetRemoteEndpointsResponse>;

}

