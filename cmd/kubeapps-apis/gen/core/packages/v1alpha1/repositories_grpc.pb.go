// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

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

// RepositoriesServiceClient is the client API for RepositoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoriesServiceClient interface {
	AddPackageRepository(ctx context.Context, in *AddPackageRepositoryRequest, opts ...grpc.CallOption) (*AddPackageRepositoryResponse, error)
}

type repositoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoriesServiceClient(cc grpc.ClientConnInterface) RepositoriesServiceClient {
	return &repositoriesServiceClient{cc}
}

func (c *repositoriesServiceClient) AddPackageRepository(ctx context.Context, in *AddPackageRepositoryRequest, opts ...grpc.CallOption) (*AddPackageRepositoryResponse, error) {
	out := new(AddPackageRepositoryResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.core.packages.v1alpha1.RepositoriesService/AddPackageRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoriesServiceServer is the server API for RepositoriesService service.
// All implementations should embed UnimplementedRepositoriesServiceServer
// for forward compatibility
type RepositoriesServiceServer interface {
	AddPackageRepository(context.Context, *AddPackageRepositoryRequest) (*AddPackageRepositoryResponse, error)
}

// UnimplementedRepositoriesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRepositoriesServiceServer struct {
}

func (UnimplementedRepositoriesServiceServer) AddPackageRepository(context.Context, *AddPackageRepositoryRequest) (*AddPackageRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPackageRepository not implemented")
}

// UnsafeRepositoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoriesServiceServer will
// result in compilation errors.
type UnsafeRepositoriesServiceServer interface {
	mustEmbedUnimplementedRepositoriesServiceServer()
}

func RegisterRepositoriesServiceServer(s grpc.ServiceRegistrar, srv RepositoriesServiceServer) {
	s.RegisterService(&RepositoriesService_ServiceDesc, srv)
}

func _RepositoriesService_AddPackageRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPackageRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoriesServiceServer).AddPackageRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.core.packages.v1alpha1.RepositoriesService/AddPackageRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoriesServiceServer).AddPackageRepository(ctx, req.(*AddPackageRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RepositoriesService_ServiceDesc is the grpc.ServiceDesc for RepositoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepositoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeappsapis.core.packages.v1alpha1.RepositoriesService",
	HandlerType: (*RepositoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPackageRepository",
			Handler:    _RepositoriesService_AddPackageRepository_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeappsapis/core/packages/v1alpha1/repositories.proto",
}