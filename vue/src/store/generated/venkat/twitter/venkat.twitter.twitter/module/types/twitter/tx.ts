/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "venkat.twitter.twitter";

export interface MsgAddTweet {
  creator: string;
  content: string;
}

export interface MsgAddTweetResponse {}

export interface MsgAddFollower {
  creator: string;
  followerId: string;
}

export interface MsgAddFollowerResponse {}

export interface MsgFetchFeed {
  creator: string;
}

export interface MsgFetchFeedResponse {
  followerTweets: string[];
}

const baseMsgAddTweet: object = { creator: "", content: "" };

export const MsgAddTweet = {
  encode(message: MsgAddTweet, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.content !== "") {
      writer.uint32(18).string(message.content);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddTweet {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddTweet } as MsgAddTweet;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.content = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddTweet {
    const message = { ...baseMsgAddTweet } as MsgAddTweet;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = String(object.content);
    } else {
      message.content = "";
    }
    return message;
  },

  toJSON(message: MsgAddTweet): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.content !== undefined && (obj.content = message.content);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddTweet>): MsgAddTweet {
    const message = { ...baseMsgAddTweet } as MsgAddTweet;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = "";
    }
    return message;
  },
};

const baseMsgAddTweetResponse: object = {};

export const MsgAddTweetResponse = {
  encode(_: MsgAddTweetResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddTweetResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddTweetResponse } as MsgAddTweetResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddTweetResponse {
    const message = { ...baseMsgAddTweetResponse } as MsgAddTweetResponse;
    return message;
  },

  toJSON(_: MsgAddTweetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddTweetResponse>): MsgAddTweetResponse {
    const message = { ...baseMsgAddTweetResponse } as MsgAddTweetResponse;
    return message;
  },
};

const baseMsgAddFollower: object = { creator: "", followerId: "" };

export const MsgAddFollower = {
  encode(message: MsgAddFollower, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.followerId !== "") {
      writer.uint32(18).string(message.followerId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddFollower {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddFollower } as MsgAddFollower;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.followerId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddFollower {
    const message = { ...baseMsgAddFollower } as MsgAddFollower;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.followerId !== undefined && object.followerId !== null) {
      message.followerId = String(object.followerId);
    } else {
      message.followerId = "";
    }
    return message;
  },

  toJSON(message: MsgAddFollower): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.followerId !== undefined && (obj.followerId = message.followerId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddFollower>): MsgAddFollower {
    const message = { ...baseMsgAddFollower } as MsgAddFollower;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.followerId !== undefined && object.followerId !== null) {
      message.followerId = object.followerId;
    } else {
      message.followerId = "";
    }
    return message;
  },
};

const baseMsgAddFollowerResponse: object = {};

export const MsgAddFollowerResponse = {
  encode(_: MsgAddFollowerResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddFollowerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddFollowerResponse } as MsgAddFollowerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddFollowerResponse {
    const message = { ...baseMsgAddFollowerResponse } as MsgAddFollowerResponse;
    return message;
  },

  toJSON(_: MsgAddFollowerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddFollowerResponse>): MsgAddFollowerResponse {
    const message = { ...baseMsgAddFollowerResponse } as MsgAddFollowerResponse;
    return message;
  },
};

const baseMsgFetchFeed: object = { creator: "" };

export const MsgFetchFeed = {
  encode(message: MsgFetchFeed, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFetchFeed {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFetchFeed } as MsgFetchFeed;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFetchFeed {
    const message = { ...baseMsgFetchFeed } as MsgFetchFeed;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgFetchFeed): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgFetchFeed>): MsgFetchFeed {
    const message = { ...baseMsgFetchFeed } as MsgFetchFeed;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgFetchFeedResponse: object = { followerTweets: "" };

export const MsgFetchFeedResponse = {
  encode(
    message: MsgFetchFeedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.followerTweets) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFetchFeedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFetchFeedResponse } as MsgFetchFeedResponse;
    message.followerTweets = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.followerTweets.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFetchFeedResponse {
    const message = { ...baseMsgFetchFeedResponse } as MsgFetchFeedResponse;
    message.followerTweets = [];
    if (object.followerTweets !== undefined && object.followerTweets !== null) {
      for (const e of object.followerTweets) {
        message.followerTweets.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgFetchFeedResponse): unknown {
    const obj: any = {};
    if (message.followerTweets) {
      obj.followerTweets = message.followerTweets.map((e) => e);
    } else {
      obj.followerTweets = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgFetchFeedResponse>): MsgFetchFeedResponse {
    const message = { ...baseMsgFetchFeedResponse } as MsgFetchFeedResponse;
    message.followerTweets = [];
    if (object.followerTweets !== undefined && object.followerTweets !== null) {
      for (const e of object.followerTweets) {
        message.followerTweets.push(e);
      }
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddTweet(request: MsgAddTweet): Promise<MsgAddTweetResponse>;
  AddFollower(request: MsgAddFollower): Promise<MsgAddFollowerResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  FetchFeed(request: MsgFetchFeed): Promise<MsgFetchFeedResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AddTweet(request: MsgAddTweet): Promise<MsgAddTweetResponse> {
    const data = MsgAddTweet.encode(request).finish();
    const promise = this.rpc.request(
      "venkat.twitter.twitter.Msg",
      "AddTweet",
      data
    );
    return promise.then((data) => MsgAddTweetResponse.decode(new Reader(data)));
  }

  AddFollower(request: MsgAddFollower): Promise<MsgAddFollowerResponse> {
    const data = MsgAddFollower.encode(request).finish();
    const promise = this.rpc.request(
      "venkat.twitter.twitter.Msg",
      "AddFollower",
      data
    );
    return promise.then((data) =>
      MsgAddFollowerResponse.decode(new Reader(data))
    );
  }

  FetchFeed(request: MsgFetchFeed): Promise<MsgFetchFeedResponse> {
    const data = MsgFetchFeed.encode(request).finish();
    const promise = this.rpc.request(
      "venkat.twitter.twitter.Msg",
      "FetchFeed",
      data
    );
    return promise.then((data) =>
      MsgFetchFeedResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
