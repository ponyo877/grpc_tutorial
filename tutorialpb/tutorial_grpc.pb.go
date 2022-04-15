// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: tutorialpb/tutorial.proto

package tutorialpb

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

// PrinterClient is the client API for Printer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrinterClient interface {
	PrintStr(ctx context.Context, in *PrintStrRequest, opts ...grpc.CallOption) (*PrintStrResponce, error)
}

type printerClient struct {
	cc grpc.ClientConnInterface
}

func NewPrinterClient(cc grpc.ClientConnInterface) PrinterClient {
	return &printerClient{cc}
}

func (c *printerClient) PrintStr(ctx context.Context, in *PrintStrRequest, opts ...grpc.CallOption) (*PrintStrResponce, error) {
	out := new(PrintStrResponce)
	err := c.cc.Invoke(ctx, "/grpc_tutorial.Printer/PrintStr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrinterServer is the server API for Printer service.
// All implementations must embed UnimplementedPrinterServer
// for forward compatibility
type PrinterServer interface {
	PrintStr(context.Context, *PrintStrRequest) (*PrintStrResponce, error)
	mustEmbedUnimplementedPrinterServer()
}

// UnimplementedPrinterServer must be embedded to have forward compatible implementations.
type UnimplementedPrinterServer struct {
}

func (UnimplementedPrinterServer) PrintStr(context.Context, *PrintStrRequest) (*PrintStrResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintStr not implemented")
}
func (UnimplementedPrinterServer) mustEmbedUnimplementedPrinterServer() {}

// UnsafePrinterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrinterServer will
// result in compilation errors.
type UnsafePrinterServer interface {
	mustEmbedUnimplementedPrinterServer()
}

func RegisterPrinterServer(s grpc.ServiceRegistrar, srv PrinterServer) {
	s.RegisterService(&Printer_ServiceDesc, srv)
}

func _Printer_PrintStr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintStrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServer).PrintStr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_tutorial.Printer/PrintStr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServer).PrintStr(ctx, req.(*PrintStrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Printer_ServiceDesc is the grpc.ServiceDesc for Printer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Printer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_tutorial.Printer",
	HandlerType: (*PrinterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrintStr",
			Handler:    _Printer_PrintStr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tutorialpb/tutorial.proto",
}