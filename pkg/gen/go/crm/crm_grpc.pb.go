// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: crm/crm.proto

package crmv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CRMService_CreateLead_FullMethodName        = "/crm.CRMService/CreateLead"
	CRMService_GetLead_FullMethodName           = "/crm.CRMService/GetLead"
	CRMService_ListLeads_FullMethodName         = "/crm.CRMService/ListLeads"
	CRMService_UpdateLead_FullMethodName        = "/crm.CRMService/UpdateLead"
	CRMService_DeleteLead_FullMethodName        = "/crm.CRMService/DeleteLead"
	CRMService_CreateTask_FullMethodName        = "/crm.CRMService/CreateTask"
	CRMService_GetTask_FullMethodName           = "/crm.CRMService/GetTask"
	CRMService_ListTasks_FullMethodName         = "/crm.CRMService/ListTasks"
	CRMService_UpdateTask_FullMethodName        = "/crm.CRMService/UpdateTask"
	CRMService_DeleteTask_FullMethodName        = "/crm.CRMService/DeleteTask"
	CRMService_CreateNote_FullMethodName        = "/crm.CRMService/CreateNote"
	CRMService_GetNote_FullMethodName           = "/crm.CRMService/GetNote"
	CRMService_ListNotes_FullMethodName         = "/crm.CRMService/ListNotes"
	CRMService_UpdateNote_FullMethodName        = "/crm.CRMService/UpdateNote"
	CRMService_DeleteNote_FullMethodName        = "/crm.CRMService/DeleteNote"
	CRMService_CreateChatMessage_FullMethodName = "/crm.CRMService/CreateChatMessage"
	CRMService_GetChatMessage_FullMethodName    = "/crm.CRMService/GetChatMessage"
	CRMService_ListChatMessages_FullMethodName  = "/crm.CRMService/ListChatMessages"
	CRMService_UpdateChatMessage_FullMethodName = "/crm.CRMService/UpdateChatMessage"
	CRMService_DeleteChatMessage_FullMethodName = "/crm.CRMService/DeleteChatMessage"
)

// CRMServiceClient is the client API for CRMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The main service for CRM-related operations.
type CRMServiceClient interface {
	// Leads
	CreateLead(ctx context.Context, in *CreateLeadRequest, opts ...grpc.CallOption) (*CreateLeadResponse, error)
	GetLead(ctx context.Context, in *GetLeadRequest, opts ...grpc.CallOption) (*GetLeadResponse, error)
	ListLeads(ctx context.Context, in *ListLeadsRequest, opts ...grpc.CallOption) (*ListLeadsResponse, error)
	UpdateLead(ctx context.Context, in *UpdateLeadRequest, opts ...grpc.CallOption) (*UpdateLeadResponse, error)
	DeleteLead(ctx context.Context, in *DeleteLeadRequest, opts ...grpc.CallOption) (*DeleteLeadResponse, error)
	// Tasks
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
	// Notes
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*GetNoteResponse, error)
	ListNotes(ctx context.Context, in *ListNotesRequest, opts ...grpc.CallOption) (*ListNotesResponse, error)
	UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error)
	DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error)
	// Chat Messages
	CreateChatMessage(ctx context.Context, in *CreateChatMessageRequest, opts ...grpc.CallOption) (*CreateChatMessageResponse, error)
	GetChatMessage(ctx context.Context, in *GetChatMessageRequest, opts ...grpc.CallOption) (*GetChatMessageResponse, error)
	ListChatMessages(ctx context.Context, in *ListChatMessagesRequest, opts ...grpc.CallOption) (*ListChatMessagesResponse, error)
	UpdateChatMessage(ctx context.Context, in *UpdateChatMessageRequest, opts ...grpc.CallOption) (*UpdateChatMessageResponse, error)
	DeleteChatMessage(ctx context.Context, in *DeleteChatMessageRequest, opts ...grpc.CallOption) (*DeleteChatMessageResponse, error)
}

type cRMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCRMServiceClient(cc grpc.ClientConnInterface) CRMServiceClient {
	return &cRMServiceClient{cc}
}

func (c *cRMServiceClient) CreateLead(ctx context.Context, in *CreateLeadRequest, opts ...grpc.CallOption) (*CreateLeadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLeadResponse)
	err := c.cc.Invoke(ctx, CRMService_CreateLead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) GetLead(ctx context.Context, in *GetLeadRequest, opts ...grpc.CallOption) (*GetLeadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLeadResponse)
	err := c.cc.Invoke(ctx, CRMService_GetLead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) ListLeads(ctx context.Context, in *ListLeadsRequest, opts ...grpc.CallOption) (*ListLeadsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLeadsResponse)
	err := c.cc.Invoke(ctx, CRMService_ListLeads_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) UpdateLead(ctx context.Context, in *UpdateLeadRequest, opts ...grpc.CallOption) (*UpdateLeadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateLeadResponse)
	err := c.cc.Invoke(ctx, CRMService_UpdateLead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) DeleteLead(ctx context.Context, in *DeleteLeadRequest, opts ...grpc.CallOption) (*DeleteLeadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLeadResponse)
	err := c.cc.Invoke(ctx, CRMService_DeleteLead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, CRMService_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, CRMService_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTasksResponse)
	err := c.cc.Invoke(ctx, CRMService_ListTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, CRMService_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, CRMService_DeleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, CRMService_CreateNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*GetNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNoteResponse)
	err := c.cc.Invoke(ctx, CRMService_GetNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) ListNotes(ctx context.Context, in *ListNotesRequest, opts ...grpc.CallOption) (*ListNotesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNotesResponse)
	err := c.cc.Invoke(ctx, CRMService_ListNotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNoteResponse)
	err := c.cc.Invoke(ctx, CRMService_UpdateNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteNoteResponse)
	err := c.cc.Invoke(ctx, CRMService_DeleteNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) CreateChatMessage(ctx context.Context, in *CreateChatMessageRequest, opts ...grpc.CallOption) (*CreateChatMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChatMessageResponse)
	err := c.cc.Invoke(ctx, CRMService_CreateChatMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) GetChatMessage(ctx context.Context, in *GetChatMessageRequest, opts ...grpc.CallOption) (*GetChatMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatMessageResponse)
	err := c.cc.Invoke(ctx, CRMService_GetChatMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) ListChatMessages(ctx context.Context, in *ListChatMessagesRequest, opts ...grpc.CallOption) (*ListChatMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListChatMessagesResponse)
	err := c.cc.Invoke(ctx, CRMService_ListChatMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) UpdateChatMessage(ctx context.Context, in *UpdateChatMessageRequest, opts ...grpc.CallOption) (*UpdateChatMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateChatMessageResponse)
	err := c.cc.Invoke(ctx, CRMService_UpdateChatMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRMServiceClient) DeleteChatMessage(ctx context.Context, in *DeleteChatMessageRequest, opts ...grpc.CallOption) (*DeleteChatMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteChatMessageResponse)
	err := c.cc.Invoke(ctx, CRMService_DeleteChatMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRMServiceServer is the server API for CRMService service.
// All implementations must embed UnimplementedCRMServiceServer
// for forward compatibility.
//
// The main service for CRM-related operations.
type CRMServiceServer interface {
	// Leads
	CreateLead(context.Context, *CreateLeadRequest) (*CreateLeadResponse, error)
	GetLead(context.Context, *GetLeadRequest) (*GetLeadResponse, error)
	ListLeads(context.Context, *ListLeadsRequest) (*ListLeadsResponse, error)
	UpdateLead(context.Context, *UpdateLeadRequest) (*UpdateLeadResponse, error)
	DeleteLead(context.Context, *DeleteLeadRequest) (*DeleteLeadResponse, error)
	// Tasks
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error)
	// Notes
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	GetNote(context.Context, *GetNoteRequest) (*GetNoteResponse, error)
	ListNotes(context.Context, *ListNotesRequest) (*ListNotesResponse, error)
	UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error)
	DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error)
	// Chat Messages
	CreateChatMessage(context.Context, *CreateChatMessageRequest) (*CreateChatMessageResponse, error)
	GetChatMessage(context.Context, *GetChatMessageRequest) (*GetChatMessageResponse, error)
	ListChatMessages(context.Context, *ListChatMessagesRequest) (*ListChatMessagesResponse, error)
	UpdateChatMessage(context.Context, *UpdateChatMessageRequest) (*UpdateChatMessageResponse, error)
	DeleteChatMessage(context.Context, *DeleteChatMessageRequest) (*DeleteChatMessageResponse, error)
	mustEmbedUnimplementedCRMServiceServer()
}

// UnimplementedCRMServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCRMServiceServer struct{}

func (UnimplementedCRMServiceServer) CreateLead(context.Context, *CreateLeadRequest) (*CreateLeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLead not implemented")
}
func (UnimplementedCRMServiceServer) GetLead(context.Context, *GetLeadRequest) (*GetLeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLead not implemented")
}
func (UnimplementedCRMServiceServer) ListLeads(context.Context, *ListLeadsRequest) (*ListLeadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLeads not implemented")
}
func (UnimplementedCRMServiceServer) UpdateLead(context.Context, *UpdateLeadRequest) (*UpdateLeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLead not implemented")
}
func (UnimplementedCRMServiceServer) DeleteLead(context.Context, *DeleteLeadRequest) (*DeleteLeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLead not implemented")
}
func (UnimplementedCRMServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedCRMServiceServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedCRMServiceServer) ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedCRMServiceServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedCRMServiceServer) DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedCRMServiceServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedCRMServiceServer) GetNote(context.Context, *GetNoteRequest) (*GetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNote not implemented")
}
func (UnimplementedCRMServiceServer) ListNotes(context.Context, *ListNotesRequest) (*ListNotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotes not implemented")
}
func (UnimplementedCRMServiceServer) UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedCRMServiceServer) DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedCRMServiceServer) CreateChatMessage(context.Context, *CreateChatMessageRequest) (*CreateChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatMessage not implemented")
}
func (UnimplementedCRMServiceServer) GetChatMessage(context.Context, *GetChatMessageRequest) (*GetChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMessage not implemented")
}
func (UnimplementedCRMServiceServer) ListChatMessages(context.Context, *ListChatMessagesRequest) (*ListChatMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChatMessages not implemented")
}
func (UnimplementedCRMServiceServer) UpdateChatMessage(context.Context, *UpdateChatMessageRequest) (*UpdateChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatMessage not implemented")
}
func (UnimplementedCRMServiceServer) DeleteChatMessage(context.Context, *DeleteChatMessageRequest) (*DeleteChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChatMessage not implemented")
}
func (UnimplementedCRMServiceServer) mustEmbedUnimplementedCRMServiceServer() {}
func (UnimplementedCRMServiceServer) testEmbeddedByValue()                    {}

// UnsafeCRMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRMServiceServer will
// result in compilation errors.
type UnsafeCRMServiceServer interface {
	mustEmbedUnimplementedCRMServiceServer()
}

func RegisterCRMServiceServer(s grpc.ServiceRegistrar, srv CRMServiceServer) {
	// If the following call pancis, it indicates UnimplementedCRMServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CRMService_ServiceDesc, srv)
}

func _CRMService_CreateLead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).CreateLead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_CreateLead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).CreateLead(ctx, req.(*CreateLeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_GetLead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).GetLead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_GetLead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).GetLead(ctx, req.(*GetLeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_ListLeads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLeadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).ListLeads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_ListLeads_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).ListLeads(ctx, req.(*ListLeadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_UpdateLead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).UpdateLead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_UpdateLead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).UpdateLead(ctx, req.(*UpdateLeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_DeleteLead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).DeleteLead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_DeleteLead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).DeleteLead(ctx, req.(*DeleteLeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_ListTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).ListTasks(ctx, req.(*ListTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_CreateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_GetNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).GetNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_GetNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).GetNote(ctx, req.(*GetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_ListNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).ListNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_ListNotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).ListNotes(ctx, req.(*ListNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_UpdateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).UpdateNote(ctx, req.(*UpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_DeleteNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).DeleteNote(ctx, req.(*DeleteNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_CreateChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).CreateChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_CreateChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).CreateChatMessage(ctx, req.(*CreateChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_GetChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).GetChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_GetChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).GetChatMessage(ctx, req.(*GetChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_ListChatMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChatMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).ListChatMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_ListChatMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).ListChatMessages(ctx, req.(*ListChatMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_UpdateChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).UpdateChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_UpdateChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).UpdateChatMessage(ctx, req.(*UpdateChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRMService_DeleteChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRMServiceServer).DeleteChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRMService_DeleteChatMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRMServiceServer).DeleteChatMessage(ctx, req.(*DeleteChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CRMService_ServiceDesc is the grpc.ServiceDesc for CRMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crm.CRMService",
	HandlerType: (*CRMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLead",
			Handler:    _CRMService_CreateLead_Handler,
		},
		{
			MethodName: "GetLead",
			Handler:    _CRMService_GetLead_Handler,
		},
		{
			MethodName: "ListLeads",
			Handler:    _CRMService_ListLeads_Handler,
		},
		{
			MethodName: "UpdateLead",
			Handler:    _CRMService_UpdateLead_Handler,
		},
		{
			MethodName: "DeleteLead",
			Handler:    _CRMService_DeleteLead_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _CRMService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _CRMService_GetTask_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _CRMService_ListTasks_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _CRMService_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _CRMService_DeleteTask_Handler,
		},
		{
			MethodName: "CreateNote",
			Handler:    _CRMService_CreateNote_Handler,
		},
		{
			MethodName: "GetNote",
			Handler:    _CRMService_GetNote_Handler,
		},
		{
			MethodName: "ListNotes",
			Handler:    _CRMService_ListNotes_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _CRMService_UpdateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _CRMService_DeleteNote_Handler,
		},
		{
			MethodName: "CreateChatMessage",
			Handler:    _CRMService_CreateChatMessage_Handler,
		},
		{
			MethodName: "GetChatMessage",
			Handler:    _CRMService_GetChatMessage_Handler,
		},
		{
			MethodName: "ListChatMessages",
			Handler:    _CRMService_ListChatMessages_Handler,
		},
		{
			MethodName: "UpdateChatMessage",
			Handler:    _CRMService_UpdateChatMessage_Handler,
		},
		{
			MethodName: "DeleteChatMessage",
			Handler:    _CRMService_DeleteChatMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crm/crm.proto",
}