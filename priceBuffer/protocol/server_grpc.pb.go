// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

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

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	GetPrice(ctx context.Context, opts ...grpc.CallOption) (CurrencyService_GetPriceClient, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) GetPrice(ctx context.Context, opts ...grpc.CallOption) (CurrencyService_GetPriceClient, error) {
	stream, err := c.cc.NewStream(ctx, &CurrencyService_ServiceDesc.Streams[0], "/proto.CurrencyService/GetPrice", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyServiceGetPriceClient{stream}
	return x, nil
}

type CurrencyService_GetPriceClient interface {
	Send(*GetPriceRequest) error
	Recv() (*GetPriceResponse, error)
	grpc.ClientStream
}

type currencyServiceGetPriceClient struct {
	grpc.ClientStream
}

func (x *currencyServiceGetPriceClient) Send(m *GetPriceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *currencyServiceGetPriceClient) Recv() (*GetPriceResponse, error) {
	m := new(GetPriceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
// All implementations must embed UnimplementedCurrencyServiceServer
// for forward compatibility
type CurrencyServiceServer interface {
	GetPrice(CurrencyService_GetPriceServer) error
	mustEmbedUnimplementedCurrencyServiceServer()
}

// UnimplementedCurrencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (UnimplementedCurrencyServiceServer) GetPrice(CurrencyService_GetPriceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPrice not implemented")
}
func (UnimplementedCurrencyServiceServer) mustEmbedUnimplementedCurrencyServiceServer() {}

// UnsafeCurrencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServiceServer will
// result in compilation errors.
type UnsafeCurrencyServiceServer interface {
	mustEmbedUnimplementedCurrencyServiceServer()
}

func RegisterCurrencyServiceServer(s grpc.ServiceRegistrar, srv CurrencyServiceServer) {
	s.RegisterService(&CurrencyService_ServiceDesc, srv)
}

func _CurrencyService_GetPrice_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CurrencyServiceServer).GetPrice(&currencyServiceGetPriceServer{stream})
}

type CurrencyService_GetPriceServer interface {
	Send(*GetPriceResponse) error
	Recv() (*GetPriceRequest, error)
	grpc.ServerStream
}

type currencyServiceGetPriceServer struct {
	grpc.ServerStream
}

func (x *currencyServiceGetPriceServer) Send(m *GetPriceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *currencyServiceGetPriceServer) Recv() (*GetPriceRequest, error) {
	m := new(GetPriceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyService_ServiceDesc is the grpc.ServiceDesc for CurrencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPrice",
			Handler:       _CurrencyService_GetPrice_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protocol/server.proto",
}
