/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "venkat.twitter.twitter";

export interface UserData {
  index: string;
  tweets: string[];
  followers: string[];
}

const baseUserData: object = { index: "", tweets: "", followers: "" };

export const UserData = {
  encode(message: UserData, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    for (const v of message.tweets) {
      writer.uint32(18).string(v!);
    }
    for (const v of message.followers) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UserData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUserData } as UserData;
    message.tweets = [];
    message.followers = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.tweets.push(reader.string());
          break;
        case 3:
          message.followers.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserData {
    const message = { ...baseUserData } as UserData;
    message.tweets = [];
    message.followers = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.tweets !== undefined && object.tweets !== null) {
      for (const e of object.tweets) {
        message.tweets.push(String(e));
      }
    }
    if (object.followers !== undefined && object.followers !== null) {
      for (const e of object.followers) {
        message.followers.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: UserData): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    if (message.tweets) {
      obj.tweets = message.tweets.map((e) => e);
    } else {
      obj.tweets = [];
    }
    if (message.followers) {
      obj.followers = message.followers.map((e) => e);
    } else {
      obj.followers = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<UserData>): UserData {
    const message = { ...baseUserData } as UserData;
    message.tweets = [];
    message.followers = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.tweets !== undefined && object.tweets !== null) {
      for (const e of object.tweets) {
        message.tweets.push(e);
      }
    }
    if (object.followers !== undefined && object.followers !== null) {
      for (const e of object.followers) {
        message.followers.push(e);
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
