/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "venkat.twitter.twitter";

export interface UserFeed {
  index: string;
  feed: string[];
}

const baseUserFeed: object = { index: "", feed: "" };

export const UserFeed = {
  encode(message: UserFeed, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    for (const v of message.feed) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UserFeed {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUserFeed } as UserFeed;
    message.feed = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.feed.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserFeed {
    const message = { ...baseUserFeed } as UserFeed;
    message.feed = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.feed !== undefined && object.feed !== null) {
      for (const e of object.feed) {
        message.feed.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: UserFeed): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    if (message.feed) {
      obj.feed = message.feed.map((e) => e);
    } else {
      obj.feed = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<UserFeed>): UserFeed {
    const message = { ...baseUserFeed } as UserFeed;
    message.feed = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.feed !== undefined && object.feed !== null) {
      for (const e of object.feed) {
        message.feed.push(e);
      }
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
