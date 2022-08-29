/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../twitter/params";
import { UserData } from "../twitter/user_data";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { UserFeed } from "../twitter/user_feed";

export const protobufPackage = "venkat.twitter.twitter";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetUserDataRequest {
  index: string;
}

export interface QueryGetUserDataResponse {
  userData: UserData | undefined;
}

export interface QueryAllUserDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserDataResponse {
  userData: UserData[];
  pagination: PageResponse | undefined;
}

export interface QueryGetUserFeedRequest {
  index: string;
}

export interface QueryGetUserFeedResponse {
  userFeed: UserFeed | undefined;
}

export interface QueryAllUserFeedRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserFeedResponse {
  userFeed: UserFeed[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetUserDataRequest: object = { index: "" };

export const QueryGetUserDataRequest = {
  encode(
    message: QueryGetUserDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUserDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUserDataRequest,
    } as QueryGetUserDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserDataRequest {
    const message = {
      ...baseQueryGetUserDataRequest,
    } as QueryGetUserDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetUserDataRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUserDataRequest>
  ): QueryGetUserDataRequest {
    const message = {
      ...baseQueryGetUserDataRequest,
    } as QueryGetUserDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetUserDataResponse: object = {};

export const QueryGetUserDataResponse = {
  encode(
    message: QueryGetUserDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.userData !== undefined) {
      UserData.encode(message.userData, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetUserDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUserDataResponse,
    } as QueryGetUserDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userData = UserData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserDataResponse {
    const message = {
      ...baseQueryGetUserDataResponse,
    } as QueryGetUserDataResponse;
    if (object.userData !== undefined && object.userData !== null) {
      message.userData = UserData.fromJSON(object.userData);
    } else {
      message.userData = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetUserDataResponse): unknown {
    const obj: any = {};
    message.userData !== undefined &&
      (obj.userData = message.userData
        ? UserData.toJSON(message.userData)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUserDataResponse>
  ): QueryGetUserDataResponse {
    const message = {
      ...baseQueryGetUserDataResponse,
    } as QueryGetUserDataResponse;
    if (object.userData !== undefined && object.userData !== null) {
      message.userData = UserData.fromPartial(object.userData);
    } else {
      message.userData = undefined;
    }
    return message;
  },
};

const baseQueryAllUserDataRequest: object = {};

export const QueryAllUserDataRequest = {
  encode(
    message: QueryAllUserDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllUserDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllUserDataRequest,
    } as QueryAllUserDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserDataRequest {
    const message = {
      ...baseQueryAllUserDataRequest,
    } as QueryAllUserDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllUserDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllUserDataRequest>
  ): QueryAllUserDataRequest {
    const message = {
      ...baseQueryAllUserDataRequest,
    } as QueryAllUserDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllUserDataResponse: object = {};

export const QueryAllUserDataResponse = {
  encode(
    message: QueryAllUserDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.userData) {
      UserData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllUserDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllUserDataResponse,
    } as QueryAllUserDataResponse;
    message.userData = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userData.push(UserData.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserDataResponse {
    const message = {
      ...baseQueryAllUserDataResponse,
    } as QueryAllUserDataResponse;
    message.userData = [];
    if (object.userData !== undefined && object.userData !== null) {
      for (const e of object.userData) {
        message.userData.push(UserData.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllUserDataResponse): unknown {
    const obj: any = {};
    if (message.userData) {
      obj.userData = message.userData.map((e) =>
        e ? UserData.toJSON(e) : undefined
      );
    } else {
      obj.userData = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllUserDataResponse>
  ): QueryAllUserDataResponse {
    const message = {
      ...baseQueryAllUserDataResponse,
    } as QueryAllUserDataResponse;
    message.userData = [];
    if (object.userData !== undefined && object.userData !== null) {
      for (const e of object.userData) {
        message.userData.push(UserData.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetUserFeedRequest: object = { index: "" };

export const QueryGetUserFeedRequest = {
  encode(
    message: QueryGetUserFeedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUserFeedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUserFeedRequest,
    } as QueryGetUserFeedRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserFeedRequest {
    const message = {
      ...baseQueryGetUserFeedRequest,
    } as QueryGetUserFeedRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetUserFeedRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUserFeedRequest>
  ): QueryGetUserFeedRequest {
    const message = {
      ...baseQueryGetUserFeedRequest,
    } as QueryGetUserFeedRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetUserFeedResponse: object = {};

export const QueryGetUserFeedResponse = {
  encode(
    message: QueryGetUserFeedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.userFeed !== undefined) {
      UserFeed.encode(message.userFeed, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetUserFeedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUserFeedResponse,
    } as QueryGetUserFeedResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userFeed = UserFeed.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserFeedResponse {
    const message = {
      ...baseQueryGetUserFeedResponse,
    } as QueryGetUserFeedResponse;
    if (object.userFeed !== undefined && object.userFeed !== null) {
      message.userFeed = UserFeed.fromJSON(object.userFeed);
    } else {
      message.userFeed = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetUserFeedResponse): unknown {
    const obj: any = {};
    message.userFeed !== undefined &&
      (obj.userFeed = message.userFeed
        ? UserFeed.toJSON(message.userFeed)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUserFeedResponse>
  ): QueryGetUserFeedResponse {
    const message = {
      ...baseQueryGetUserFeedResponse,
    } as QueryGetUserFeedResponse;
    if (object.userFeed !== undefined && object.userFeed !== null) {
      message.userFeed = UserFeed.fromPartial(object.userFeed);
    } else {
      message.userFeed = undefined;
    }
    return message;
  },
};

const baseQueryAllUserFeedRequest: object = {};

export const QueryAllUserFeedRequest = {
  encode(
    message: QueryAllUserFeedRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllUserFeedRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllUserFeedRequest,
    } as QueryAllUserFeedRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserFeedRequest {
    const message = {
      ...baseQueryAllUserFeedRequest,
    } as QueryAllUserFeedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllUserFeedRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllUserFeedRequest>
  ): QueryAllUserFeedRequest {
    const message = {
      ...baseQueryAllUserFeedRequest,
    } as QueryAllUserFeedRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllUserFeedResponse: object = {};

export const QueryAllUserFeedResponse = {
  encode(
    message: QueryAllUserFeedResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.userFeed) {
      UserFeed.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllUserFeedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllUserFeedResponse,
    } as QueryAllUserFeedResponse;
    message.userFeed = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userFeed.push(UserFeed.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserFeedResponse {
    const message = {
      ...baseQueryAllUserFeedResponse,
    } as QueryAllUserFeedResponse;
    message.userFeed = [];
    if (object.userFeed !== undefined && object.userFeed !== null) {
      for (const e of object.userFeed) {
        message.userFeed.push(UserFeed.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllUserFeedResponse): unknown {
    const obj: any = {};
    if (message.userFeed) {
      obj.userFeed = message.userFeed.map((e) =>
        e ? UserFeed.toJSON(e) : undefined
      );
    } else {
      obj.userFeed = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllUserFeedResponse>
  ): QueryAllUserFeedResponse {
    const message = {
      ...baseQueryAllUserFeedResponse,
    } as QueryAllUserFeedResponse;
    message.userFeed = [];
    if (object.userFeed !== undefined && object.userFeed !== null) {
      for (const e of object.userFeed) {
        message.userFeed.push(UserFeed.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a UserData by index. */
  UserData(request: QueryGetUserDataRequest): Promise<QueryGetUserDataResponse>;
  /** Queries a list of UserData items. */
  UserDataAll(
    request: QueryAllUserDataRequest
  ): Promise<QueryAllUserDataResponse>;
  /** Queries a UserFeed by index. */
  UserFeed(request: QueryGetUserFeedRequest): Promise<QueryGetUserFeedResponse>;
  /** Queries a list of UserFeed items. */
  UserFeedAll(
    request: QueryAllUserFeedRequest
  ): Promise<QueryAllUserFeedResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "venkat.twitter.twitter.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  UserData(
    request: QueryGetUserDataRequest
  ): Promise<QueryGetUserDataResponse> {
    const data = QueryGetUserDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "venkat.twitter.twitter.Query",
      "UserData",
      data
    );
    return promise.then((data) =>
      QueryGetUserDataResponse.decode(new Reader(data))
    );
  }

  UserDataAll(
    request: QueryAllUserDataRequest
  ): Promise<QueryAllUserDataResponse> {
    const data = QueryAllUserDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "venkat.twitter.twitter.Query",
      "UserDataAll",
      data
    );
    return promise.then((data) =>
      QueryAllUserDataResponse.decode(new Reader(data))
    );
  }

  UserFeed(
    request: QueryGetUserFeedRequest
  ): Promise<QueryGetUserFeedResponse> {
    const data = QueryGetUserFeedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "venkat.twitter.twitter.Query",
      "UserFeed",
      data
    );
    return promise.then((data) =>
      QueryGetUserFeedResponse.decode(new Reader(data))
    );
  }

  UserFeedAll(
    request: QueryAllUserFeedRequest
  ): Promise<QueryAllUserFeedResponse> {
    const data = QueryAllUserFeedRequest.encode(request).finish();
    const promise = this.rpc.request(
      "venkat.twitter.twitter.Query",
      "UserFeedAll",
      data
    );
    return promise.then((data) =>
      QueryAllUserFeedResponse.decode(new Reader(data))
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
