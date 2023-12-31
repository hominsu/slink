// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file slink/service/v1/user.proto (package slink.service.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message slink.service.v1.LoginRequest
 */
export const LoginRequest = proto3.makeMessageType(
  "slink.service.v1.LoginRequest",
  () => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message slink.service.v1.LoginReply
 */
export const LoginReply = proto3.makeMessageType(
  "slink.service.v1.LoginReply",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expires_at", kind: "message", T: Timestamp },
  ],
);
