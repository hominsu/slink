// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file slink/service/v1/site.proto (package slink.service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message slink.service.v1.PingReply
 */
export const PingReply = proto3.makeMessageType(
  "slink.service.v1.PingReply",
  () => [
    { no: 1, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message slink.service.v1.CountReply
 */
export const CountReply = proto3.makeMessageType(
  "slink.service.v1.CountReply",
  () => [
    { no: 1, name: "count", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);
