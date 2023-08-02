import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as portal_remote_endpoint_pb from './portal_remote_endpoint_pb';
import * as portal_ping_pb from './portal_ping_pb';


export class GetRemoteEndpointsArgs extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRemoteEndpointsArgs.AsObject;
  static toObject(includeInstance: boolean, msg: GetRemoteEndpointsArgs): GetRemoteEndpointsArgs.AsObject;
  static serializeBinaryToWriter(message: GetRemoteEndpointsArgs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRemoteEndpointsArgs;
  static deserializeBinaryFromReader(message: GetRemoteEndpointsArgs, reader: jspb.BinaryReader): GetRemoteEndpointsArgs;
}

export namespace GetRemoteEndpointsArgs {
  export type AsObject = {
  }
}

export class RemoteEndpoint extends jspb.Message {
  getRemoteEndpointType(): portal_remote_endpoint_pb.RemoteEndpointType;
  setRemoteEndpointType(value: portal_remote_endpoint_pb.RemoteEndpointType): RemoteEndpoint;

  getRemoteHost(): string;
  setRemoteHost(value: string): RemoteEndpoint;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoteEndpoint.AsObject;
  static toObject(includeInstance: boolean, msg: RemoteEndpoint): RemoteEndpoint.AsObject;
  static serializeBinaryToWriter(message: RemoteEndpoint, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoteEndpoint;
  static deserializeBinaryFromReader(message: RemoteEndpoint, reader: jspb.BinaryReader): RemoteEndpoint;
}

export namespace RemoteEndpoint {
  export type AsObject = {
    remoteEndpointType: portal_remote_endpoint_pb.RemoteEndpointType,
    remoteHost: string,
  }
}

export class GetRemoteEndpointsResponse extends jspb.Message {
  getRemoteEndpointsList(): Array<RemoteEndpoint>;
  setRemoteEndpointsList(value: Array<RemoteEndpoint>): GetRemoteEndpointsResponse;
  clearRemoteEndpointsList(): GetRemoteEndpointsResponse;
  addRemoteEndpoints(value?: RemoteEndpoint, index?: number): RemoteEndpoint;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRemoteEndpointsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetRemoteEndpointsResponse): GetRemoteEndpointsResponse.AsObject;
  static serializeBinaryToWriter(message: GetRemoteEndpointsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRemoteEndpointsResponse;
  static deserializeBinaryFromReader(message: GetRemoteEndpointsResponse, reader: jspb.BinaryReader): GetRemoteEndpointsResponse;
}

export namespace GetRemoteEndpointsResponse {
  export type AsObject = {
    remoteEndpointsList: Array<RemoteEndpoint.AsObject>,
  }
}

