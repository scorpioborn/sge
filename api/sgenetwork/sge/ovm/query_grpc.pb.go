// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: sgenetwork/sge/ovm/query.proto

package ovm

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
	Query_Params_FullMethodName                    = "/sgenetwork.sge.ovm.Query/Params"
	Query_PubKeys_FullMethodName                   = "/sgenetwork.sge.ovm.Query/PubKeys"
	Query_PublicKeysChangeProposal_FullMethodName  = "/sgenetwork.sge.ovm.Query/PublicKeysChangeProposal"
	Query_PublicKeysChangeProposals_FullMethodName = "/sgenetwork.sge.ovm.Query/PublicKeysChangeProposals"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query defines the gRPC querier service.
type QueryClient interface {
	// Params queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of PubKeys items.
	PubKeys(ctx context.Context, in *QueryPubKeysRequest, opts ...grpc.CallOption) (*QueryPubKeysResponse, error)
	// Queries an proposal by id and status.
	PublicKeysChangeProposal(ctx context.Context, in *QueryPublicKeysChangeProposalRequest, opts ...grpc.CallOption) (*QueryPublicKeysChangeProposalResponse, error)
	// Queries a list of proposal items by status.
	PublicKeysChangeProposals(ctx context.Context, in *QueryPublicKeysChangeProposalsRequest, opts ...grpc.CallOption) (*QueryPublicKeysChangeProposalsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PubKeys(ctx context.Context, in *QueryPubKeysRequest, opts ...grpc.CallOption) (*QueryPubKeysResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPubKeysResponse)
	err := c.cc.Invoke(ctx, Query_PubKeys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PublicKeysChangeProposal(ctx context.Context, in *QueryPublicKeysChangeProposalRequest, opts ...grpc.CallOption) (*QueryPublicKeysChangeProposalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPublicKeysChangeProposalResponse)
	err := c.cc.Invoke(ctx, Query_PublicKeysChangeProposal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PublicKeysChangeProposals(ctx context.Context, in *QueryPublicKeysChangeProposalsRequest, opts ...grpc.CallOption) (*QueryPublicKeysChangeProposalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPublicKeysChangeProposalsResponse)
	err := c.cc.Invoke(ctx, Query_PublicKeysChangeProposals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query defines the gRPC querier service.
type QueryServer interface {
	// Params queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of PubKeys items.
	PubKeys(context.Context, *QueryPubKeysRequest) (*QueryPubKeysResponse, error)
	// Queries an proposal by id and status.
	PublicKeysChangeProposal(context.Context, *QueryPublicKeysChangeProposalRequest) (*QueryPublicKeysChangeProposalResponse, error)
	// Queries a list of proposal items by status.
	PublicKeysChangeProposals(context.Context, *QueryPublicKeysChangeProposalsRequest) (*QueryPublicKeysChangeProposalsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) PubKeys(context.Context, *QueryPubKeysRequest) (*QueryPubKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PubKeys not implemented")
}
func (UnimplementedQueryServer) PublicKeysChangeProposal(context.Context, *QueryPublicKeysChangeProposalRequest) (*QueryPublicKeysChangeProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicKeysChangeProposal not implemented")
}
func (UnimplementedQueryServer) PublicKeysChangeProposals(context.Context, *QueryPublicKeysChangeProposalsRequest) (*QueryPublicKeysChangeProposalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicKeysChangeProposals not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PubKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPubKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PubKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PubKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PubKeys(ctx, req.(*QueryPubKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PublicKeysChangeProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPublicKeysChangeProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PublicKeysChangeProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PublicKeysChangeProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PublicKeysChangeProposal(ctx, req.(*QueryPublicKeysChangeProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PublicKeysChangeProposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPublicKeysChangeProposalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PublicKeysChangeProposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PublicKeysChangeProposals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PublicKeysChangeProposals(ctx, req.(*QueryPublicKeysChangeProposalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sgenetwork.sge.ovm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "PubKeys",
			Handler:    _Query_PubKeys_Handler,
		},
		{
			MethodName: "PublicKeysChangeProposal",
			Handler:    _Query_PublicKeysChangeProposal_Handler,
		},
		{
			MethodName: "PublicKeysChangeProposals",
			Handler:    _Query_PublicKeysChangeProposals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sgenetwork/sge/ovm/query.proto",
}
