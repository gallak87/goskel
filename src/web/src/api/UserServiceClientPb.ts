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


export class UserServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetUser = new grpcWeb.AbstractClientBase.MethodInfo(
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
    callback: (err: grpcWeb.Error,
               response: user_pb.GetUserResponse) => void): grpcWeb.ClientReadableStream<user_pb.GetUserResponse>;

  getUser(
    request: user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: user_pb.GetUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/server.UserService/GetUser',
        request,
        metadata || {},
        this.methodInfoGetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/server.UserService/GetUser',
    request,
    metadata || {},
    this.methodInfoGetUser);
  }

  methodInfoRegisterUser = new grpcWeb.AbstractClientBase.MethodInfo(
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
    callback: (err: grpcWeb.Error,
               response: user_pb.RegisterUserResponse) => void): grpcWeb.ClientReadableStream<user_pb.RegisterUserResponse>;

  registerUser(
    request: user_pb.RegisterUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: user_pb.RegisterUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/server.UserService/RegisterUser',
        request,
        metadata || {},
        this.methodInfoRegisterUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/server.UserService/RegisterUser',
    request,
    metadata || {},
    this.methodInfoRegisterUser);
  }

  methodInfoDeleteUser = new grpcWeb.AbstractClientBase.MethodInfo(
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
    callback: (err: grpcWeb.Error,
               response: user_pb.DeleteUserResponse) => void): grpcWeb.ClientReadableStream<user_pb.DeleteUserResponse>;

  deleteUser(
    request: user_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: user_pb.DeleteUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/server.UserService/DeleteUser',
        request,
        metadata || {},
        this.methodInfoDeleteUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/server.UserService/DeleteUser',
    request,
    metadata || {},
    this.methodInfoDeleteUser);
  }

}

