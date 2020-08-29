/* eslint-disable */

// package: server
// file: user.proto

import * as user_pb from "./user_pb";
import {grpc} from "@improbable-eng/grpc-web";

type UserServiceGetUser = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof user_pb.GetUserRequest;
  readonly responseType: typeof user_pb.GetUserResponse;
};

type UserServiceRegisterUser = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof user_pb.RegisterUserRequest;
  readonly responseType: typeof user_pb.RegisterUserResponse;
};

type UserServiceDeleteUser = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof user_pb.DeleteUserRequest;
  readonly responseType: typeof user_pb.DeleteUserResponse;
};

export class UserService {
  static readonly serviceName: string;
  static readonly GetUser: UserServiceGetUser;
  static readonly RegisterUser: UserServiceRegisterUser;
  static readonly DeleteUser: UserServiceDeleteUser;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class UserServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getUser(
    requestMessage: user_pb.GetUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: user_pb.GetUserResponse|null) => void
  ): UnaryResponse;
  getUser(
    requestMessage: user_pb.GetUserRequest,
    callback: (error: ServiceError|null, responseMessage: user_pb.GetUserResponse|null) => void
  ): UnaryResponse;
  registerUser(
    requestMessage: user_pb.RegisterUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: user_pb.RegisterUserResponse|null) => void
  ): UnaryResponse;
  registerUser(
    requestMessage: user_pb.RegisterUserRequest,
    callback: (error: ServiceError|null, responseMessage: user_pb.RegisterUserResponse|null) => void
  ): UnaryResponse;
  deleteUser(
    requestMessage: user_pb.DeleteUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: user_pb.DeleteUserResponse|null) => void
  ): UnaryResponse;
  deleteUser(
    requestMessage: user_pb.DeleteUserRequest,
    callback: (error: ServiceError|null, responseMessage: user_pb.DeleteUserResponse|null) => void
  ): UnaryResponse;
}


