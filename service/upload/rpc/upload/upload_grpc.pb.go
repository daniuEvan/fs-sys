// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: upload.proto

package upload

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UploadClient is the client API for Upload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadClient interface {
	// FileUpload 文件上传
	FileUpload(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadResponse, error)
	// 更新文件表 todo 与更新用户表的事务问题
	UpdateFileTable(ctx context.Context, in *UserTableUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	// 更新用户文件表
	UpdateUserTable(ctx context.Context, in *UserTableUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	// 秒传
	FastUpload(ctx context.Context, in *FastUploadRequest, opts ...grpc.CallOption) (*FastUploadResponse, error)
}

type uploadClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadClient(cc grpc.ClientConnInterface) UploadClient {
	return &uploadClient{cc}
}

func (c *uploadClient) FileUpload(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadResponse, error) {
	out := new(FileUploadResponse)
	err := c.cc.Invoke(ctx, "/upload.Upload/FileUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadClient) UpdateFileTable(ctx context.Context, in *UserTableUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/upload.Upload/UpdateFileTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadClient) UpdateUserTable(ctx context.Context, in *UserTableUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/upload.Upload/UpdateUserTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadClient) FastUpload(ctx context.Context, in *FastUploadRequest, opts ...grpc.CallOption) (*FastUploadResponse, error) {
	out := new(FastUploadResponse)
	err := c.cc.Invoke(ctx, "/upload.Upload/FastUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadServer is the server API for Upload service.
// All implementations must embed UnimplementedUploadServer
// for forward compatibility
type UploadServer interface {
	// FileUpload 文件上传
	FileUpload(context.Context, *FileUploadRequest) (*FileUploadResponse, error)
	// 更新文件表 todo 与更新用户表的事务问题
	UpdateFileTable(context.Context, *UserTableUpdateRequest) (*Empty, error)
	// 更新用户文件表
	UpdateUserTable(context.Context, *UserTableUpdateRequest) (*Empty, error)
	// 秒传
	FastUpload(context.Context, *FastUploadRequest) (*FastUploadResponse, error)
	mustEmbedUnimplementedUploadServer()
}

// UnimplementedUploadServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServer struct {
}

func (UnimplementedUploadServer) FileUpload(context.Context, *FileUploadRequest) (*FileUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileUpload not implemented")
}
func (UnimplementedUploadServer) UpdateFileTable(context.Context, *UserTableUpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFileTable not implemented")
}
func (UnimplementedUploadServer) UpdateUserTable(context.Context, *UserTableUpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserTable not implemented")
}
func (UnimplementedUploadServer) FastUpload(context.Context, *FastUploadRequest) (*FastUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FastUpload not implemented")
}
func (UnimplementedUploadServer) mustEmbedUnimplementedUploadServer() {}

// UnsafeUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServer will
// result in compilation errors.
type UnsafeUploadServer interface {
	mustEmbedUnimplementedUploadServer()
}

func RegisterUploadServer(s grpc.ServiceRegistrar, srv UploadServer) {
	s.RegisterService(&Upload_ServiceDesc, srv)
}

func _Upload_FileUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).FileUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/upload.Upload/FileUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).FileUpload(ctx, req.(*FileUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upload_UpdateFileTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTableUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).UpdateFileTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/upload.Upload/UpdateFileTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).UpdateFileTable(ctx, req.(*UserTableUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upload_UpdateUserTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTableUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).UpdateUserTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/upload.Upload/UpdateUserTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).UpdateUserTable(ctx, req.(*UserTableUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upload_FastUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).FastUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/upload.Upload/FastUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).FastUpload(ctx, req.(*FastUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Upload_ServiceDesc is the grpc.ServiceDesc for Upload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "upload.Upload",
	HandlerType: (*UploadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileUpload",
			Handler:    _Upload_FileUpload_Handler,
		},
		{
			MethodName: "UpdateFileTable",
			Handler:    _Upload_UpdateFileTable_Handler,
		},
		{
			MethodName: "UpdateUserTable",
			Handler:    _Upload_UpdateUserTable_Handler,
		},
		{
			MethodName: "FastUpload",
			Handler:    _Upload_FastUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload.proto",
}
