// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var docencia_pb = require('./docencia_pb.js');

function serialize_coreGRPC_ChangePasswordReply(arg) {
  if (!(arg instanceof docencia_pb.ChangePasswordReply)) {
    throw new Error('Expected argument of type coreGRPC.ChangePasswordReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_coreGRPC_ChangePasswordReply(buffer_arg) {
  return docencia_pb.ChangePasswordReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_coreGRPC_ChangePasswordRequest(arg) {
  if (!(arg instanceof docencia_pb.ChangePasswordRequest)) {
    throw new Error('Expected argument of type coreGRPC.ChangePasswordRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_coreGRPC_ChangePasswordRequest(buffer_arg) {
  return docencia_pb.ChangePasswordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_coreGRPC_CodeReply(arg) {
  if (!(arg instanceof docencia_pb.CodeReply)) {
    throw new Error('Expected argument of type coreGRPC.CodeReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_coreGRPC_CodeReply(buffer_arg) {
  return docencia_pb.CodeReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_coreGRPC_CodeRequest(arg) {
  if (!(arg instanceof docencia_pb.CodeRequest)) {
    throw new Error('Expected argument of type coreGRPC.CodeRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_coreGRPC_CodeRequest(buffer_arg) {
  return docencia_pb.CodeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_coreGRPC_FindOrCreateUserReply(arg) {
  if (!(arg instanceof docencia_pb.FindOrCreateUserReply)) {
    throw new Error('Expected argument of type coreGRPC.FindOrCreateUserReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_coreGRPC_FindOrCreateUserReply(buffer_arg) {
  return docencia_pb.FindOrCreateUserReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_coreGRPC_FindOrCreateUserRequest(arg) {
  if (!(arg instanceof docencia_pb.FindOrCreateUserRequest)) {
    throw new Error('Expected argument of type coreGRPC.FindOrCreateUserRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_coreGRPC_FindOrCreateUserRequest(buffer_arg) {
  return docencia_pb.FindOrCreateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_coreGRPC_InfoIDRequest(arg) {
  if (!(arg instanceof docencia_pb.InfoIDRequest)) {
    throw new Error('Expected argument of type coreGRPC.InfoIDRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_coreGRPC_InfoIDRequest(buffer_arg) {
  return docencia_pb.InfoIDRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_coreGRPC_InfoReply(arg) {
  if (!(arg instanceof docencia_pb.InfoReply)) {
    throw new Error('Expected argument of type coreGRPC.InfoReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_coreGRPC_InfoReply(buffer_arg) {
  return docencia_pb.InfoReply.deserializeBinary(new Uint8Array(buffer_arg));
}


// The signer service definition.
var CoreService = exports.CoreService = {
  // Sends a file to sign.
  changePassword: {
    path: '/coreGRPC.Core/ChangePassword',
    requestStream: false,
    responseStream: false,
    requestType: docencia_pb.ChangePasswordRequest,
    responseType: docencia_pb.ChangePasswordReply,
    requestSerialize: serialize_coreGRPC_ChangePasswordRequest,
    requestDeserialize: deserialize_coreGRPC_ChangePasswordRequest,
    responseSerialize: serialize_coreGRPC_ChangePasswordReply,
    responseDeserialize: deserialize_coreGRPC_ChangePasswordReply,
  },
  findOrCreateUser: {
    path: '/coreGRPC.Core/FindOrCreateUser',
    requestStream: false,
    responseStream: false,
    requestType: docencia_pb.FindOrCreateUserRequest,
    responseType: docencia_pb.FindOrCreateUserReply,
    requestSerialize: serialize_coreGRPC_FindOrCreateUserRequest,
    requestDeserialize: deserialize_coreGRPC_FindOrCreateUserRequest,
    responseSerialize: serialize_coreGRPC_FindOrCreateUserReply,
    responseDeserialize: deserialize_coreGRPC_FindOrCreateUserReply,
  },
  getInfoByID: {
    path: '/coreGRPC.Core/GetInfoByID',
    requestStream: false,
    responseStream: false,
    requestType: docencia_pb.InfoIDRequest,
    responseType: docencia_pb.InfoReply,
    requestSerialize: serialize_coreGRPC_InfoIDRequest,
    requestDeserialize: deserialize_coreGRPC_InfoIDRequest,
    responseSerialize: serialize_coreGRPC_InfoReply,
    responseDeserialize: deserialize_coreGRPC_InfoReply,
  },
};

exports.CoreClient = grpc.makeGenericClientConstructor(CoreService);
var CompilerService = exports.CompilerService = {
  compiler: {
    path: '/coreGRPC.Compiler/Compiler',
    requestStream: false,
    responseStream: false,
    requestType: docencia_pb.CodeRequest,
    responseType: docencia_pb.CodeReply,
    requestSerialize: serialize_coreGRPC_CodeRequest,
    requestDeserialize: deserialize_coreGRPC_CodeRequest,
    responseSerialize: serialize_coreGRPC_CodeReply,
    responseDeserialize: deserialize_coreGRPC_CodeReply,
  },
};

exports.CompilerClient = grpc.makeGenericClientConstructor(CompilerService);
