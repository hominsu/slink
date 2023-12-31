// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file slink/service/v1/admin.proto (package slink.service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message slink.service.v1.ListShortLinkRequest
 */
export declare class ListShortLinkRequest extends Message<ListShortLinkRequest> {
  /**
   * @generated from field: int32 page_size = 1;
   */
  pageSize: number;

  /**
   * @generated from field: string page_token = 2;
   */
  pageToken: string;

  constructor(data?: PartialMessage<ListShortLinkRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "slink.service.v1.ListShortLinkRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListShortLinkRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListShortLinkRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListShortLinkRequest;

  static equals(a: ListShortLinkRequest | PlainMessage<ListShortLinkRequest> | undefined, b: ListShortLinkRequest | PlainMessage<ListShortLinkRequest> | undefined): boolean;
}

/**
 * @generated from message slink.service.v1.ListShortLinkReply
 */
export declare class ListShortLinkReply extends Message<ListShortLinkReply> {
  /**
   * @generated from field: repeated string keys = 1;
   */
  keys: string[];

  /**
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;

  constructor(data?: PartialMessage<ListShortLinkReply>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "slink.service.v1.ListShortLinkReply";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListShortLinkReply;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListShortLinkReply;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListShortLinkReply;

  static equals(a: ListShortLinkReply | PlainMessage<ListShortLinkReply> | undefined, b: ListShortLinkReply | PlainMessage<ListShortLinkReply> | undefined): boolean;
}

/**
 * @generated from message slink.service.v1.DeleteShortLinkRequest
 */
export declare class DeleteShortLinkRequest extends Message<DeleteShortLinkRequest> {
  /**
   * @generated from field: string key = 1;
   */
  key: string;

  constructor(data?: PartialMessage<DeleteShortLinkRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "slink.service.v1.DeleteShortLinkRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteShortLinkRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteShortLinkRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteShortLinkRequest;

  static equals(a: DeleteShortLinkRequest | PlainMessage<DeleteShortLinkRequest> | undefined, b: DeleteShortLinkRequest | PlainMessage<DeleteShortLinkRequest> | undefined): boolean;
}

