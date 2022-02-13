// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/jukebox-app.proto

package grpc

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

// JukeboxClient is the client API for Jukebox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JukeboxClient interface {
	QuerySongsByName(ctx context.Context, in *SongRequest, opts ...grpc.CallOption) (*SongResponse, error)
	ScheduleSong(ctx context.Context, in *SongRequest, opts ...grpc.CallOption) (*SongResponse, error)
}

type jukeboxClient struct {
	cc grpc.ClientConnInterface
}

func NewJukeboxClient(cc grpc.ClientConnInterface) JukeboxClient {
	return &jukeboxClient{cc}
}

func (c *jukeboxClient) QuerySongsByName(ctx context.Context, in *SongRequest, opts ...grpc.CallOption) (*SongResponse, error) {
	out := new(SongResponse)
	err := c.cc.Invoke(ctx, "/Jukebox/QuerySongsByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jukeboxClient) ScheduleSong(ctx context.Context, in *SongRequest, opts ...grpc.CallOption) (*SongResponse, error) {
	out := new(SongResponse)
	err := c.cc.Invoke(ctx, "/Jukebox/ScheduleSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JukeboxServer is the server API for Jukebox service.
// All implementations must embed UnimplementedJukeboxServer
// for forward compatibility
type JukeboxServer interface {
	QuerySongsByName(context.Context, *SongRequest) (*SongResponse, error)
	ScheduleSong(context.Context, *SongRequest) (*SongResponse, error)
	mustEmbedUnimplementedJukeboxServer()
}

// UnimplementedJukeboxServer must be embedded to have forward compatible implementations.
type UnimplementedJukeboxServer struct {
}

func (UnimplementedJukeboxServer) QuerySongsByName(context.Context, *SongRequest) (*SongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySongsByName not implemented")
}
func (UnimplementedJukeboxServer) ScheduleSong(context.Context, *SongRequest) (*SongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleSong not implemented")
}
func (UnimplementedJukeboxServer) mustEmbedUnimplementedJukeboxServer() {}

// UnsafeJukeboxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JukeboxServer will
// result in compilation errors.
type UnsafeJukeboxServer interface {
	mustEmbedUnimplementedJukeboxServer()
}

func RegisterJukeboxServer(s grpc.ServiceRegistrar, srv JukeboxServer) {
	s.RegisterService(&Jukebox_ServiceDesc, srv)
}

func _Jukebox_QuerySongsByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JukeboxServer).QuerySongsByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Jukebox/QuerySongsByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JukeboxServer).QuerySongsByName(ctx, req.(*SongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jukebox_ScheduleSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JukeboxServer).ScheduleSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Jukebox/ScheduleSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JukeboxServer).ScheduleSong(ctx, req.(*SongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Jukebox_ServiceDesc is the grpc.ServiceDesc for Jukebox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jukebox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Jukebox",
	HandlerType: (*JukeboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QuerySongsByName",
			Handler:    _Jukebox_QuerySongsByName_Handler,
		},
		{
			MethodName: "ScheduleSong",
			Handler:    _Jukebox_ScheduleSong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/jukebox-app.proto",
}
