// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=js+dts"
// @generated from file slink/service/v1/site.proto (package slink.service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { CountReply, PingReply } from "./site_pb.js";

/**
 * @generated from service slink.service.v1.SiteService
 */
export declare const SiteService: {
  readonly typeName: "slink.service.v1.SiteService",
  readonly methods: {
    /**
     * @generated from rpc slink.service.v1.SiteService.Ping
     */
    readonly ping: {
      readonly name: "Ping",
      readonly I: typeof Empty,
      readonly O: typeof PingReply,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc slink.service.v1.SiteService.Count
     */
    readonly count: {
      readonly name: "Count",
      readonly I: typeof Empty,
      readonly O: typeof CountReply,
      readonly kind: MethodKind.Unary,
    },
  }
};

