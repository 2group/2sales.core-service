// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: user/user.proto

#include "user/user.pb.h"
#include "user/user.grpc.pb.h"

#include <functional>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/impl/channel_interface.h>
#include <grpcpp/impl/client_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/rpc_service_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/sync_stream.h>
namespace user {

static const char* UserService_method_names[] = {
  "/user.UserService/Login",
  "/user.UserService/Register",
  "/user.UserService/UpdateUser",
  "/user.UserService/PatchUser",
  "/user.UserService/GetUser",
  "/user.UserService/DeleteUser",
  "/user.UserService/ListUsers",
  "/user.UserService/CreateUser",
  "/user.UserService/CreateRole",
  "/user.UserService/ListRoles",
  "/user.UserService/UpdateRole",
  "/user.UserService/DeleteRole",
};

std::unique_ptr< UserService::Stub> UserService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< UserService::Stub> stub(new UserService::Stub(channel, options));
  return stub;
}

UserService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Login_(UserService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Register_(UserService_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_UpdateUser_(UserService_method_names[2], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_PatchUser_(UserService_method_names[3], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetUser_(UserService_method_names[4], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_DeleteUser_(UserService_method_names[5], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_ListUsers_(UserService_method_names[6], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_CreateUser_(UserService_method_names[7], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_CreateRole_(UserService_method_names[8], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_ListRoles_(UserService_method_names[9], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_UpdateRole_(UserService_method_names[10], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_DeleteRole_(UserService_method_names[11], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status UserService::Stub::Login(::grpc::ClientContext* context, const ::user::LoginRequest& request, ::user::LoginResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::LoginRequest, ::user::LoginResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Login_, context, request, response);
}

void UserService::Stub::async::Login(::grpc::ClientContext* context, const ::user::LoginRequest* request, ::user::LoginResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::LoginRequest, ::user::LoginResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Login_, context, request, response, std::move(f));
}

void UserService::Stub::async::Login(::grpc::ClientContext* context, const ::user::LoginRequest* request, ::user::LoginResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Login_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::LoginResponse>* UserService::Stub::PrepareAsyncLoginRaw(::grpc::ClientContext* context, const ::user::LoginRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::LoginResponse, ::user::LoginRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Login_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::LoginResponse>* UserService::Stub::AsyncLoginRaw(::grpc::ClientContext* context, const ::user::LoginRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncLoginRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::Register(::grpc::ClientContext* context, const ::user::RegisterRequest& request, ::user::RegisterResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::RegisterRequest, ::user::RegisterResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Register_, context, request, response);
}

void UserService::Stub::async::Register(::grpc::ClientContext* context, const ::user::RegisterRequest* request, ::user::RegisterResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::RegisterRequest, ::user::RegisterResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Register_, context, request, response, std::move(f));
}

void UserService::Stub::async::Register(::grpc::ClientContext* context, const ::user::RegisterRequest* request, ::user::RegisterResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Register_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::RegisterResponse>* UserService::Stub::PrepareAsyncRegisterRaw(::grpc::ClientContext* context, const ::user::RegisterRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::RegisterResponse, ::user::RegisterRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Register_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::RegisterResponse>* UserService::Stub::AsyncRegisterRaw(::grpc::ClientContext* context, const ::user::RegisterRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncRegisterRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::UpdateUser(::grpc::ClientContext* context, const ::user::UpdateUserRequest& request, ::user::UpdateUserResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::UpdateUserRequest, ::user::UpdateUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_UpdateUser_, context, request, response);
}

void UserService::Stub::async::UpdateUser(::grpc::ClientContext* context, const ::user::UpdateUserRequest* request, ::user::UpdateUserResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::UpdateUserRequest, ::user::UpdateUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_UpdateUser_, context, request, response, std::move(f));
}

void UserService::Stub::async::UpdateUser(::grpc::ClientContext* context, const ::user::UpdateUserRequest* request, ::user::UpdateUserResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_UpdateUser_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::UpdateUserResponse>* UserService::Stub::PrepareAsyncUpdateUserRaw(::grpc::ClientContext* context, const ::user::UpdateUserRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::UpdateUserResponse, ::user::UpdateUserRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_UpdateUser_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::UpdateUserResponse>* UserService::Stub::AsyncUpdateUserRaw(::grpc::ClientContext* context, const ::user::UpdateUserRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncUpdateUserRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::PatchUser(::grpc::ClientContext* context, const ::user::PatchUserRequest& request, ::user::PatchUserResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::PatchUserRequest, ::user::PatchUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_PatchUser_, context, request, response);
}

void UserService::Stub::async::PatchUser(::grpc::ClientContext* context, const ::user::PatchUserRequest* request, ::user::PatchUserResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::PatchUserRequest, ::user::PatchUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_PatchUser_, context, request, response, std::move(f));
}

void UserService::Stub::async::PatchUser(::grpc::ClientContext* context, const ::user::PatchUserRequest* request, ::user::PatchUserResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_PatchUser_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::PatchUserResponse>* UserService::Stub::PrepareAsyncPatchUserRaw(::grpc::ClientContext* context, const ::user::PatchUserRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::PatchUserResponse, ::user::PatchUserRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_PatchUser_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::PatchUserResponse>* UserService::Stub::AsyncPatchUserRaw(::grpc::ClientContext* context, const ::user::PatchUserRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncPatchUserRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::GetUser(::grpc::ClientContext* context, const ::user::GetUserRequest& request, ::user::GetUserResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::GetUserRequest, ::user::GetUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetUser_, context, request, response);
}

void UserService::Stub::async::GetUser(::grpc::ClientContext* context, const ::user::GetUserRequest* request, ::user::GetUserResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::GetUserRequest, ::user::GetUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetUser_, context, request, response, std::move(f));
}

void UserService::Stub::async::GetUser(::grpc::ClientContext* context, const ::user::GetUserRequest* request, ::user::GetUserResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetUser_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::GetUserResponse>* UserService::Stub::PrepareAsyncGetUserRaw(::grpc::ClientContext* context, const ::user::GetUserRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::GetUserResponse, ::user::GetUserRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetUser_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::GetUserResponse>* UserService::Stub::AsyncGetUserRaw(::grpc::ClientContext* context, const ::user::GetUserRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetUserRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::DeleteUser(::grpc::ClientContext* context, const ::user::DeleteUserRequest& request, ::user::DeleteUserResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::DeleteUserRequest, ::user::DeleteUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_DeleteUser_, context, request, response);
}

void UserService::Stub::async::DeleteUser(::grpc::ClientContext* context, const ::user::DeleteUserRequest* request, ::user::DeleteUserResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::DeleteUserRequest, ::user::DeleteUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_DeleteUser_, context, request, response, std::move(f));
}

void UserService::Stub::async::DeleteUser(::grpc::ClientContext* context, const ::user::DeleteUserRequest* request, ::user::DeleteUserResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_DeleteUser_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::DeleteUserResponse>* UserService::Stub::PrepareAsyncDeleteUserRaw(::grpc::ClientContext* context, const ::user::DeleteUserRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::DeleteUserResponse, ::user::DeleteUserRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_DeleteUser_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::DeleteUserResponse>* UserService::Stub::AsyncDeleteUserRaw(::grpc::ClientContext* context, const ::user::DeleteUserRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncDeleteUserRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::ListUsers(::grpc::ClientContext* context, const ::user::ListUsersRequest& request, ::user::ListUsersResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::ListUsersRequest, ::user::ListUsersResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_ListUsers_, context, request, response);
}

void UserService::Stub::async::ListUsers(::grpc::ClientContext* context, const ::user::ListUsersRequest* request, ::user::ListUsersResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::ListUsersRequest, ::user::ListUsersResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_ListUsers_, context, request, response, std::move(f));
}

void UserService::Stub::async::ListUsers(::grpc::ClientContext* context, const ::user::ListUsersRequest* request, ::user::ListUsersResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_ListUsers_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::ListUsersResponse>* UserService::Stub::PrepareAsyncListUsersRaw(::grpc::ClientContext* context, const ::user::ListUsersRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::ListUsersResponse, ::user::ListUsersRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_ListUsers_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::ListUsersResponse>* UserService::Stub::AsyncListUsersRaw(::grpc::ClientContext* context, const ::user::ListUsersRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncListUsersRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::CreateUser(::grpc::ClientContext* context, const ::user::CreateUserRequest& request, ::user::CreateUserResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::CreateUserRequest, ::user::CreateUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_CreateUser_, context, request, response);
}

void UserService::Stub::async::CreateUser(::grpc::ClientContext* context, const ::user::CreateUserRequest* request, ::user::CreateUserResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::CreateUserRequest, ::user::CreateUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_CreateUser_, context, request, response, std::move(f));
}

void UserService::Stub::async::CreateUser(::grpc::ClientContext* context, const ::user::CreateUserRequest* request, ::user::CreateUserResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_CreateUser_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::CreateUserResponse>* UserService::Stub::PrepareAsyncCreateUserRaw(::grpc::ClientContext* context, const ::user::CreateUserRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::CreateUserResponse, ::user::CreateUserRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_CreateUser_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::CreateUserResponse>* UserService::Stub::AsyncCreateUserRaw(::grpc::ClientContext* context, const ::user::CreateUserRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncCreateUserRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::CreateRole(::grpc::ClientContext* context, const ::user::CreateRoleRequest& request, ::user::CreateRoleResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::CreateRoleRequest, ::user::CreateRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_CreateRole_, context, request, response);
}

void UserService::Stub::async::CreateRole(::grpc::ClientContext* context, const ::user::CreateRoleRequest* request, ::user::CreateRoleResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::CreateRoleRequest, ::user::CreateRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_CreateRole_, context, request, response, std::move(f));
}

void UserService::Stub::async::CreateRole(::grpc::ClientContext* context, const ::user::CreateRoleRequest* request, ::user::CreateRoleResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_CreateRole_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::CreateRoleResponse>* UserService::Stub::PrepareAsyncCreateRoleRaw(::grpc::ClientContext* context, const ::user::CreateRoleRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::CreateRoleResponse, ::user::CreateRoleRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_CreateRole_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::CreateRoleResponse>* UserService::Stub::AsyncCreateRoleRaw(::grpc::ClientContext* context, const ::user::CreateRoleRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncCreateRoleRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::ListRoles(::grpc::ClientContext* context, const ::user::ListRolesRequest& request, ::user::ListRolesResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::ListRolesRequest, ::user::ListRolesResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_ListRoles_, context, request, response);
}

void UserService::Stub::async::ListRoles(::grpc::ClientContext* context, const ::user::ListRolesRequest* request, ::user::ListRolesResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::ListRolesRequest, ::user::ListRolesResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_ListRoles_, context, request, response, std::move(f));
}

void UserService::Stub::async::ListRoles(::grpc::ClientContext* context, const ::user::ListRolesRequest* request, ::user::ListRolesResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_ListRoles_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::ListRolesResponse>* UserService::Stub::PrepareAsyncListRolesRaw(::grpc::ClientContext* context, const ::user::ListRolesRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::ListRolesResponse, ::user::ListRolesRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_ListRoles_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::ListRolesResponse>* UserService::Stub::AsyncListRolesRaw(::grpc::ClientContext* context, const ::user::ListRolesRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncListRolesRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::UpdateRole(::grpc::ClientContext* context, const ::user::UpdateRoleRequest& request, ::user::UpdateRoleResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::UpdateRoleRequest, ::user::UpdateRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_UpdateRole_, context, request, response);
}

void UserService::Stub::async::UpdateRole(::grpc::ClientContext* context, const ::user::UpdateRoleRequest* request, ::user::UpdateRoleResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::UpdateRoleRequest, ::user::UpdateRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_UpdateRole_, context, request, response, std::move(f));
}

void UserService::Stub::async::UpdateRole(::grpc::ClientContext* context, const ::user::UpdateRoleRequest* request, ::user::UpdateRoleResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_UpdateRole_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::UpdateRoleResponse>* UserService::Stub::PrepareAsyncUpdateRoleRaw(::grpc::ClientContext* context, const ::user::UpdateRoleRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::UpdateRoleResponse, ::user::UpdateRoleRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_UpdateRole_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::UpdateRoleResponse>* UserService::Stub::AsyncUpdateRoleRaw(::grpc::ClientContext* context, const ::user::UpdateRoleRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncUpdateRoleRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status UserService::Stub::DeleteRole(::grpc::ClientContext* context, const ::user::DeleteRoleRequest& request, ::user::DeleteRoleResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::user::DeleteRoleRequest, ::user::DeleteRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_DeleteRole_, context, request, response);
}

void UserService::Stub::async::DeleteRole(::grpc::ClientContext* context, const ::user::DeleteRoleRequest* request, ::user::DeleteRoleResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::user::DeleteRoleRequest, ::user::DeleteRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_DeleteRole_, context, request, response, std::move(f));
}

void UserService::Stub::async::DeleteRole(::grpc::ClientContext* context, const ::user::DeleteRoleRequest* request, ::user::DeleteRoleResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_DeleteRole_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::user::DeleteRoleResponse>* UserService::Stub::PrepareAsyncDeleteRoleRaw(::grpc::ClientContext* context, const ::user::DeleteRoleRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::user::DeleteRoleResponse, ::user::DeleteRoleRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_DeleteRole_, context, request);
}

::grpc::ClientAsyncResponseReader< ::user::DeleteRoleResponse>* UserService::Stub::AsyncDeleteRoleRaw(::grpc::ClientContext* context, const ::user::DeleteRoleRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncDeleteRoleRaw(context, request, cq);
  result->StartCall();
  return result;
}

UserService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::LoginRequest, ::user::LoginResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::LoginRequest* req,
             ::user::LoginResponse* resp) {
               return service->Login(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::RegisterRequest, ::user::RegisterResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::RegisterRequest* req,
             ::user::RegisterResponse* resp) {
               return service->Register(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::UpdateUserRequest, ::user::UpdateUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::UpdateUserRequest* req,
             ::user::UpdateUserResponse* resp) {
               return service->UpdateUser(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[3],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::PatchUserRequest, ::user::PatchUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::PatchUserRequest* req,
             ::user::PatchUserResponse* resp) {
               return service->PatchUser(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[4],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::GetUserRequest, ::user::GetUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::GetUserRequest* req,
             ::user::GetUserResponse* resp) {
               return service->GetUser(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[5],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::DeleteUserRequest, ::user::DeleteUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::DeleteUserRequest* req,
             ::user::DeleteUserResponse* resp) {
               return service->DeleteUser(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[6],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::ListUsersRequest, ::user::ListUsersResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::ListUsersRequest* req,
             ::user::ListUsersResponse* resp) {
               return service->ListUsers(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[7],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::CreateUserRequest, ::user::CreateUserResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::CreateUserRequest* req,
             ::user::CreateUserResponse* resp) {
               return service->CreateUser(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[8],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::CreateRoleRequest, ::user::CreateRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::CreateRoleRequest* req,
             ::user::CreateRoleResponse* resp) {
               return service->CreateRole(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[9],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::ListRolesRequest, ::user::ListRolesResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::ListRolesRequest* req,
             ::user::ListRolesResponse* resp) {
               return service->ListRoles(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[10],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::UpdateRoleRequest, ::user::UpdateRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::UpdateRoleRequest* req,
             ::user::UpdateRoleResponse* resp) {
               return service->UpdateRole(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      UserService_method_names[11],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< UserService::Service, ::user::DeleteRoleRequest, ::user::DeleteRoleResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](UserService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::user::DeleteRoleRequest* req,
             ::user::DeleteRoleResponse* resp) {
               return service->DeleteRole(ctx, req, resp);
             }, this)));
}

UserService::Service::~Service() {
}

::grpc::Status UserService::Service::Login(::grpc::ServerContext* context, const ::user::LoginRequest* request, ::user::LoginResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::Register(::grpc::ServerContext* context, const ::user::RegisterRequest* request, ::user::RegisterResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::UpdateUser(::grpc::ServerContext* context, const ::user::UpdateUserRequest* request, ::user::UpdateUserResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::PatchUser(::grpc::ServerContext* context, const ::user::PatchUserRequest* request, ::user::PatchUserResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::GetUser(::grpc::ServerContext* context, const ::user::GetUserRequest* request, ::user::GetUserResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::DeleteUser(::grpc::ServerContext* context, const ::user::DeleteUserRequest* request, ::user::DeleteUserResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::ListUsers(::grpc::ServerContext* context, const ::user::ListUsersRequest* request, ::user::ListUsersResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::CreateUser(::grpc::ServerContext* context, const ::user::CreateUserRequest* request, ::user::CreateUserResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::CreateRole(::grpc::ServerContext* context, const ::user::CreateRoleRequest* request, ::user::CreateRoleResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::ListRoles(::grpc::ServerContext* context, const ::user::ListRolesRequest* request, ::user::ListRolesResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::UpdateRole(::grpc::ServerContext* context, const ::user::UpdateRoleRequest* request, ::user::UpdateRoleResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status UserService::Service::DeleteRole(::grpc::ServerContext* context, const ::user::DeleteRoleRequest* request, ::user::DeleteRoleResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace user

