import * as grpcWeb from 'grpc-web';

import * as user_pb from './user_pb';


export class UserServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getUser(
    request: user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.GetUserResponse) => void
  ): grpcWeb.ClientReadableStream<user_pb.GetUserResponse>;

  registerUser(
    request: user_pb.RegisterUserRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.RegisterUserResponse) => void
  ): grpcWeb.ClientReadableStream<user_pb.RegisterUserResponse>;

  deleteUser(
    request: user_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.DeleteUserResponse) => void
  ): grpcWeb.ClientReadableStream<user_pb.DeleteUserResponse>;

}

export class UserServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getUser(
    request: user_pb.GetUserRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.GetUserResponse>;

  registerUser(
    request: user_pb.RegisterUserRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.RegisterUserResponse>;

  deleteUser(
    request: user_pb.DeleteUserRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.DeleteUserResponse>;

}

