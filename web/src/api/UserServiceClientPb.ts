/* eslint-disable */
/**
 * @fileoverview gRPC-Web generated client stub for server
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as user_pb from './user_pb';


export class UserClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetUser = new grpcWeb.MethodDescriptor(
    '/server.User/GetUser',
    grpcWeb.MethodType.UNARY,
    user_pb.GetUserRequest,
    user_pb.GetUserResponse,
    (request: user_pb.GetUserRequest) => {
      return request.serializeBinary();
    },
    user_pb.GetUserResponse.deserializeBinary
  );

  getUser(
    request: user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_pb.GetUserResponse>;

  getUser(
    request: user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.GetUserResponse) => void): grpcWeb.ClientReadableStream<user_pb.GetUserResponse>;

  getUser(
    request: user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_pb.GetUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/server.User/GetUser',
        request,
        metadata || {},
        this.methodDescriptorGetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/server.User/GetUser',
    request,
    metadata || {},
    this.methodDescriptorGetUser);
  }

  methodDescriptorRegisterUser = new grpcWeb.MethodDescriptor(
    '/server.User/RegisterUser',
    grpcWeb.MethodType.UNARY,
    user_pb.RegisterUserRequest,
    user_pb.RegisterUserResponse,
    (request: user_pb.RegisterUserRequest) => {
      return request.serializeBinary();
    },
    user_pb.RegisterUserResponse.deserializeBinary
  );

  registerUser(
    request: user_pb.RegisterUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_pb.RegisterUserResponse>;

  registerUser(
    request: user_pb.RegisterUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.RegisterUserResponse) => void): grpcWeb.ClientReadableStream<user_pb.RegisterUserResponse>;

  registerUser(
    request: user_pb.RegisterUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_pb.RegisterUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/server.User/RegisterUser',
        request,
        metadata || {},
        this.methodDescriptorRegisterUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/server.User/RegisterUser',
    request,
    metadata || {},
    this.methodDescriptorRegisterUser);
  }

  methodDescriptorDeleteUser = new grpcWeb.MethodDescriptor(
    '/server.User/DeleteUser',
    grpcWeb.MethodType.UNARY,
    user_pb.DeleteUserRequest,
    user_pb.DeleteUserResponse,
    (request: user_pb.DeleteUserRequest) => {
      return request.serializeBinary();
    },
    user_pb.DeleteUserResponse.deserializeBinary
  );

  deleteUser(
    request: user_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_pb.DeleteUserResponse>;

  deleteUser(
    request: user_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.DeleteUserResponse) => void): grpcWeb.ClientReadableStream<user_pb.DeleteUserResponse>;

  deleteUser(
    request: user_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_pb.DeleteUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/server.User/DeleteUser',
        request,
        metadata || {},
        this.methodDescriptorDeleteUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/server.User/DeleteUser',
    request,
    metadata || {},
    this.methodDescriptorDeleteUser);
  }

}