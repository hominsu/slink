// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file slink/service/v1/slink.proto (package slink.service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message slink.service.v1.CreateShortLinkRequest
 */
export declare class CreateShortLinkRequest extends Message<CreateShortLinkRequest> {
  /**
   * @generated from field: string link = 1;
   */
  link: string;

  constructor(data?: PartialMessage<CreateShortLinkRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "slink.service.v1.CreateShortLinkRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateShortLinkRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateShortLinkRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateShortLinkRequest;

  static equals(a: CreateShortLinkRequest | PlainMessage<CreateShortLinkRequest> | undefined, b: CreateShortLinkRequest | PlainMessage<CreateShortLinkRequest> | undefined): boolean;
}

/**
 * @generated from message slink.service.v1.CreateShortLinkReply
 */
export declare class CreateShortLinkReply extends Message<CreateShortLinkReply> {
  /**
   * @generated from field: string key = 1;
   */
  key: string;

  constructor(data?: PartialMessage<CreateShortLinkReply>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "slink.service.v1.CreateShortLinkReply";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateShortLinkReply;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateShortLinkReply;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateShortLinkReply;

  static equals(a: CreateShortLinkReply | PlainMessage<CreateShortLinkReply> | undefined, b: CreateShortLinkReply | PlainMessage<CreateShortLinkReply> | undefined): boolean;
}
