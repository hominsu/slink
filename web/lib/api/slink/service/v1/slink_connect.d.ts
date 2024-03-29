// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=js+dts"
// @generated from file slink/service/v1/slink.proto (package slink.service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateShortLinkReply, CreateShortLinkRequest } from "./slink_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service slink.service.v1.ShortLinkService
 */
export declare const ShortLinkService: {
  readonly typeName: "slink.service.v1.ShortLinkService",
  readonly methods: {
    /**
     * @generated from rpc slink.service.v1.ShortLinkService.CreateShortLink
     */
    readonly createShortLink: {
      readonly name: "CreateShortLink",
      readonly I: typeof CreateShortLinkRequest,
      readonly O: typeof CreateShortLinkReply,
      readonly kind: MethodKind.Unary,
    },
  }
};

