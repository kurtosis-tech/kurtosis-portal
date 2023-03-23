import * as jspb from 'google-protobuf'



export class PortalPing extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PortalPing.AsObject;
  static toObject(includeInstance: boolean, msg: PortalPing): PortalPing.AsObject;
  static serializeBinaryToWriter(message: PortalPing, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PortalPing;
  static deserializeBinaryFromReader(message: PortalPing, reader: jspb.BinaryReader): PortalPing;
}

export namespace PortalPing {
  export type AsObject = {
  }
}

export class PortalPong extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PortalPong.AsObject;
  static toObject(includeInstance: boolean, msg: PortalPong): PortalPong.AsObject;
  static serializeBinaryToWriter(message: PortalPong, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PortalPong;
  static deserializeBinaryFromReader(message: PortalPong, reader: jspb.BinaryReader): PortalPong;
}

export namespace PortalPong {
  export type AsObject = {
  }
}

