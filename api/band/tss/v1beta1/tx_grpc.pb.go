// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: band/tss/v1beta1/tx.proto

package tssv1beta1

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

const (
	Msg_SubmitDKGRound1_FullMethodName = "/band.tss.v1beta1.Msg/SubmitDKGRound1"
	Msg_SubmitDKGRound2_FullMethodName = "/band.tss.v1beta1.Msg/SubmitDKGRound2"
	Msg_Complain_FullMethodName        = "/band.tss.v1beta1.Msg/Complain"
	Msg_Confirm_FullMethodName         = "/band.tss.v1beta1.Msg/Confirm"
	Msg_SubmitDEs_FullMethodName       = "/band.tss.v1beta1.Msg/SubmitDEs"
	Msg_ResetDE_FullMethodName         = "/band.tss.v1beta1.Msg/ResetDE"
	Msg_SubmitSignature_FullMethodName = "/band.tss.v1beta1.Msg/SubmitSignature"
	Msg_UpdateParams_FullMethodName    = "/band.tss.v1beta1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// SubmitDKGRound1 submits dkg for computing round 1.
	SubmitDKGRound1(ctx context.Context, in *MsgSubmitDKGRound1, opts ...grpc.CallOption) (*MsgSubmitDKGRound1Response, error)
	// SubmitDKGRound2 submits dkg for computing round 2.
	SubmitDKGRound2(ctx context.Context, in *MsgSubmitDKGRound2, opts ...grpc.CallOption) (*MsgSubmitDKGRound2Response, error)
	// Complain submits proof for complaining malicious.
	Complain(ctx context.Context, in *MsgComplain, opts ...grpc.CallOption) (*MsgComplainResponse, error)
	// Confirm submits own signature for proof that it can derive the secret.
	Confirm(ctx context.Context, in *MsgConfirm, opts ...grpc.CallOption) (*MsgConfirmResponse, error)
	// SubmitDEs submits list of pre-commits DE for signing process.
	SubmitDEs(ctx context.Context, in *MsgSubmitDEs, opts ...grpc.CallOption) (*MsgSubmitDEsResponse, error)
	// ResetDE resets the submitted DEs that being stored on chain.
	ResetDE(ctx context.Context, in *MsgResetDE, opts ...grpc.CallOption) (*MsgResetDEResponse, error)
	// SubmitSignature submits signature on task participant need to do.
	SubmitSignature(ctx context.Context, in *MsgSubmitSignature, opts ...grpc.CallOption) (*MsgSubmitSignatureResponse, error)
	// UpdateParams defines a governance operation for updating the x/tss module
	// parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitDKGRound1(ctx context.Context, in *MsgSubmitDKGRound1, opts ...grpc.CallOption) (*MsgSubmitDKGRound1Response, error) {
	out := new(MsgSubmitDKGRound1Response)
	err := c.cc.Invoke(ctx, Msg_SubmitDKGRound1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitDKGRound2(ctx context.Context, in *MsgSubmitDKGRound2, opts ...grpc.CallOption) (*MsgSubmitDKGRound2Response, error) {
	out := new(MsgSubmitDKGRound2Response)
	err := c.cc.Invoke(ctx, Msg_SubmitDKGRound2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Complain(ctx context.Context, in *MsgComplain, opts ...grpc.CallOption) (*MsgComplainResponse, error) {
	out := new(MsgComplainResponse)
	err := c.cc.Invoke(ctx, Msg_Complain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Confirm(ctx context.Context, in *MsgConfirm, opts ...grpc.CallOption) (*MsgConfirmResponse, error) {
	out := new(MsgConfirmResponse)
	err := c.cc.Invoke(ctx, Msg_Confirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitDEs(ctx context.Context, in *MsgSubmitDEs, opts ...grpc.CallOption) (*MsgSubmitDEsResponse, error) {
	out := new(MsgSubmitDEsResponse)
	err := c.cc.Invoke(ctx, Msg_SubmitDEs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ResetDE(ctx context.Context, in *MsgResetDE, opts ...grpc.CallOption) (*MsgResetDEResponse, error) {
	out := new(MsgResetDEResponse)
	err := c.cc.Invoke(ctx, Msg_ResetDE_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitSignature(ctx context.Context, in *MsgSubmitSignature, opts ...grpc.CallOption) (*MsgSubmitSignatureResponse, error) {
	out := new(MsgSubmitSignatureResponse)
	err := c.cc.Invoke(ctx, Msg_SubmitSignature_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// SubmitDKGRound1 submits dkg for computing round 1.
	SubmitDKGRound1(context.Context, *MsgSubmitDKGRound1) (*MsgSubmitDKGRound1Response, error)
	// SubmitDKGRound2 submits dkg for computing round 2.
	SubmitDKGRound2(context.Context, *MsgSubmitDKGRound2) (*MsgSubmitDKGRound2Response, error)
	// Complain submits proof for complaining malicious.
	Complain(context.Context, *MsgComplain) (*MsgComplainResponse, error)
	// Confirm submits own signature for proof that it can derive the secret.
	Confirm(context.Context, *MsgConfirm) (*MsgConfirmResponse, error)
	// SubmitDEs submits list of pre-commits DE for signing process.
	SubmitDEs(context.Context, *MsgSubmitDEs) (*MsgSubmitDEsResponse, error)
	// ResetDE resets the submitted DEs that being stored on chain.
	ResetDE(context.Context, *MsgResetDE) (*MsgResetDEResponse, error)
	// SubmitSignature submits signature on task participant need to do.
	SubmitSignature(context.Context, *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error)
	// UpdateParams defines a governance operation for updating the x/tss module
	// parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) SubmitDKGRound1(context.Context, *MsgSubmitDKGRound1) (*MsgSubmitDKGRound1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDKGRound1 not implemented")
}
func (UnimplementedMsgServer) SubmitDKGRound2(context.Context, *MsgSubmitDKGRound2) (*MsgSubmitDKGRound2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDKGRound2 not implemented")
}
func (UnimplementedMsgServer) Complain(context.Context, *MsgComplain) (*MsgComplainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complain not implemented")
}
func (UnimplementedMsgServer) Confirm(context.Context, *MsgConfirm) (*MsgConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (UnimplementedMsgServer) SubmitDEs(context.Context, *MsgSubmitDEs) (*MsgSubmitDEsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDEs not implemented")
}
func (UnimplementedMsgServer) ResetDE(context.Context, *MsgResetDE) (*MsgResetDEResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetDE not implemented")
}
func (UnimplementedMsgServer) SubmitSignature(context.Context, *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitSignature not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_SubmitDKGRound1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitDKGRound1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitDKGRound1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitDKGRound1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitDKGRound1(ctx, req.(*MsgSubmitDKGRound1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitDKGRound2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitDKGRound2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitDKGRound2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitDKGRound2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitDKGRound2(ctx, req.(*MsgSubmitDKGRound2))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Complain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgComplain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Complain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Complain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Complain(ctx, req.(*MsgComplain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConfirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Confirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Confirm(ctx, req.(*MsgConfirm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitDEs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitDEs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitDEs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitDEs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitDEs(ctx, req.(*MsgSubmitDEs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ResetDE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgResetDE)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ResetDE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ResetDE_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ResetDE(ctx, req.(*MsgResetDE))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitSignature)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitSignature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitSignature(ctx, req.(*MsgSubmitSignature))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "band.tss.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitDKGRound1",
			Handler:    _Msg_SubmitDKGRound1_Handler,
		},
		{
			MethodName: "SubmitDKGRound2",
			Handler:    _Msg_SubmitDKGRound2_Handler,
		},
		{
			MethodName: "Complain",
			Handler:    _Msg_Complain_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _Msg_Confirm_Handler,
		},
		{
			MethodName: "SubmitDEs",
			Handler:    _Msg_SubmitDEs_Handler,
		},
		{
			MethodName: "ResetDE",
			Handler:    _Msg_ResetDE_Handler,
		},
		{
			MethodName: "SubmitSignature",
			Handler:    _Msg_SubmitSignature_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "band/tss/v1beta1/tx.proto",
}
