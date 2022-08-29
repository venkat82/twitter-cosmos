/* eslint-disable */
import { Params } from "../twitter/params";
import { UserData } from "../twitter/user_data";
import { UserFeed } from "../twitter/user_feed";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "venkat.twitter.twitter";

/** GenesisState defines the twitter module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  userDataList: UserData[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  userFeedList: UserFeed[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.userDataList) {
      UserData.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.userFeedList) {
      UserFeed.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.userDataList = [];
    message.userFeedList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.userDataList.push(UserData.decode(reader, reader.uint32()));
          break;
        case 3:
          message.userFeedList.push(UserFeed.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.userDataList = [];
    message.userFeedList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.userDataList !== undefined && object.userDataList !== null) {
      for (const e of object.userDataList) {
        message.userDataList.push(UserData.fromJSON(e));
      }
    }
    if (object.userFeedList !== undefined && object.userFeedList !== null) {
      for (const e of object.userFeedList) {
        message.userFeedList.push(UserFeed.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.userDataList) {
      obj.userDataList = message.userDataList.map((e) =>
        e ? UserData.toJSON(e) : undefined
      );
    } else {
      obj.userDataList = [];
    }
    if (message.userFeedList) {
      obj.userFeedList = message.userFeedList.map((e) =>
        e ? UserFeed.toJSON(e) : undefined
      );
    } else {
      obj.userFeedList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.userDataList = [];
    message.userFeedList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.userDataList !== undefined && object.userDataList !== null) {
      for (const e of object.userDataList) {
        message.userDataList.push(UserData.fromPartial(e));
      }
    }
    if (object.userFeedList !== undefined && object.userFeedList !== null) {
      for (const e of object.userFeedList) {
        message.userFeedList.push(UserFeed.fromPartial(e));
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
