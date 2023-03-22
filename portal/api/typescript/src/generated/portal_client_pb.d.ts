import * as jspb from 'google-protobuf'

import * as portal_ping_pb from './portal_ping_pb';


export class SwitchContextArgs extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SwitchContextArgs.AsObject;
  static toObject(includeInstance: boolean, msg: SwitchContextArgs): SwitchContextArgs.AsObject;
  static serializeBinaryToWriter(message: SwitchContextArgs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SwitchContextArgs;
  static deserializeBinaryFromReader(message: SwitchContextArgs, reader: jspb.BinaryReader): SwitchContextArgs;
}

export namespace SwitchContextArgs {
  export type AsObject = {
  }
}

export class SwitchContextResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SwitchContextResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SwitchContextResponse): SwitchContextResponse.AsObject;
  static serializeBinaryToWriter(message: SwitchContextResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SwitchContextResponse;
  static deserializeBinaryFromReader(message: SwitchContextResponse, reader: jspb.BinaryReader): SwitchContextResponse;
}

export namespace SwitchContextResponse {
  export type AsObject = {
  }
}

export class ForwardPortArgs extends jspb.Message {
  getLocalPortNumber(): number;
  setLocalPortNumber(value: number): ForwardPortArgs;

  getRemotePortNumber(): number;
  setRemotePortNumber(value: number): ForwardPortArgs;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ForwardPortArgs.AsObject;
  static toObject(includeInstance: boolean, msg: ForwardPortArgs): ForwardPortArgs.AsObject;
  static serializeBinaryToWriter(message: ForwardPortArgs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ForwardPortArgs;
  static deserializeBinaryFromReader(message: ForwardPortArgs, reader: jspb.BinaryReader): ForwardPortArgs;
}

export namespace ForwardPortArgs {
  export type AsObject = {
    localPortNumber: number,
    remotePortNumber: number,
  }
}

export class ForwardPortResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ForwardPortResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ForwardPortResponse): ForwardPortResponse.AsObject;
  static serializeBinaryToWriter(message: ForwardPortResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ForwardPortResponse;
  static deserializeBinaryFromReader(message: ForwardPortResponse, reader: jspb.BinaryReader): ForwardPortResponse;
}

export namespace ForwardPortResponse {
  export type AsObject = {
  }
}

