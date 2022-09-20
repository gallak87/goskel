/**
 * @fileoverview gRPC-Web generated client stub for server
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.server = require('./user_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.server.UserServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.server.UserServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.server.GetUserRequest,
 *   !proto.server.GetUserResponse>}
 */
const methodDescriptor_UserService_GetUser = new grpc.web.MethodDescriptor(
  '/server.UserService/GetUser',
  grpc.web.MethodType.UNARY,
  proto.server.GetUserRequest,
  proto.server.GetUserResponse,
  /**
   * @param {!proto.server.GetUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.GetUserResponse.deserializeBinary
);


/**
 * @param {!proto.server.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.server.GetUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.server.GetUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.server.UserServiceClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/server.UserService/GetUser',
      request,
      metadata || {},
      methodDescriptor_UserService_GetUser,
      callback);
};


/**
 * @param {!proto.server.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.server.GetUserResponse>}
 *     Promise that resolves to the response
 */
proto.server.UserServicePromiseClient.prototype.getUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/server.UserService/GetUser',
      request,
      metadata || {},
      methodDescriptor_UserService_GetUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.server.RegisterUserRequest,
 *   !proto.server.RegisterUserResponse>}
 */
const methodDescriptor_UserService_RegisterUser = new grpc.web.MethodDescriptor(
  '/server.UserService/RegisterUser',
  grpc.web.MethodType.UNARY,
  proto.server.RegisterUserRequest,
  proto.server.RegisterUserResponse,
  /**
   * @param {!proto.server.RegisterUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.RegisterUserResponse.deserializeBinary
);


/**
 * @param {!proto.server.RegisterUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.server.RegisterUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.server.RegisterUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.server.UserServiceClient.prototype.registerUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/server.UserService/RegisterUser',
      request,
      metadata || {},
      methodDescriptor_UserService_RegisterUser,
      callback);
};


/**
 * @param {!proto.server.RegisterUserRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.server.RegisterUserResponse>}
 *     Promise that resolves to the response
 */
proto.server.UserServicePromiseClient.prototype.registerUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/server.UserService/RegisterUser',
      request,
      metadata || {},
      methodDescriptor_UserService_RegisterUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.server.DeleteUserRequest,
 *   !proto.server.DeleteUserResponse>}
 */
const methodDescriptor_UserService_DeleteUser = new grpc.web.MethodDescriptor(
  '/server.UserService/DeleteUser',
  grpc.web.MethodType.UNARY,
  proto.server.DeleteUserRequest,
  proto.server.DeleteUserResponse,
  /**
   * @param {!proto.server.DeleteUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.DeleteUserResponse.deserializeBinary
);


/**
 * @param {!proto.server.DeleteUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.server.DeleteUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.server.DeleteUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.server.UserServiceClient.prototype.deleteUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/server.UserService/DeleteUser',
      request,
      metadata || {},
      methodDescriptor_UserService_DeleteUser,
      callback);
};


/**
 * @param {!proto.server.DeleteUserRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.server.DeleteUserResponse>}
 *     Promise that resolves to the response
 */
proto.server.UserServicePromiseClient.prototype.deleteUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/server.UserService/DeleteUser',
      request,
      metadata || {},
      methodDescriptor_UserService_DeleteUser);
};


module.exports = proto.server;

