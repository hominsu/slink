// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=js+dts"
// @generated from file slink/service/v1/admin.proto (package slink.service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { DeleteShortLinkRequest, ListShortLinkReply, ListShortLinkRequest } from "./admin_pb.js";

/**
 * @generated from service slink.service.v1.AdminService
 */
export declare const AdminService: {
  readonly typeName: "slink.service.v1.AdminService",
  readonly methods: {
    /**
     * @generated from rpc slink.service.v1.AdminService.FlushShortLink
     */
    readonly flushShortLink: {
      readonly name: "FlushShortLink",
      readonly I: typeof Empty,
      readonly O: typeof Empty,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc slink.service.v1.AdminService.ListShortLinks
     */
    readonly listShortLinks: {
      readonly name: "ListShortLinks",
      readonly I: typeof ListShortLinkRequest,
      readonly O: typeof ListShortLinkReply,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc slink.service.v1.AdminService.DeleteShortLink
     */
    readonly deleteShortLink: {
      readonly name: "DeleteShortLink",
      readonly I: typeof DeleteShortLinkRequest,
      readonly O: typeof Empty,
      readonly kind: MethodKind.Unary,
    },
  }
};

