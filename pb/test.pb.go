// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/test.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ParamA struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParamA) Reset()         { *m = ParamA{} }
func (m *ParamA) String() string { return proto.CompactTextString(m) }
func (*ParamA) ProtoMessage()    {}
func (*ParamA) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff63324815cff01, []int{0}
}

func (m *ParamA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamA.Unmarshal(m, b)
}
func (m *ParamA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamA.Marshal(b, m, deterministic)
}
func (m *ParamA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamA.Merge(m, src)
}
func (m *ParamA) XXX_Size() int {
	return xxx_messageInfo_ParamA.Size(m)
}
func (m *ParamA) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamA.DiscardUnknown(m)
}

var xxx_messageInfo_ParamA proto.InternalMessageInfo

func (m *ParamA) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ParamA) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type ResponseA struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	ParentSpanId         string   `protobuf:"bytes,2,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	Name                 string   `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,12,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseA) Reset()         { *m = ResponseA{} }
func (m *ResponseA) String() string { return proto.CompactTextString(m) }
func (*ResponseA) ProtoMessage()    {}
func (*ResponseA) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff63324815cff01, []int{1}
}

func (m *ResponseA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseA.Unmarshal(m, b)
}
func (m *ResponseA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseA.Marshal(b, m, deterministic)
}
func (m *ResponseA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseA.Merge(m, src)
}
func (m *ResponseA) XXX_Size() int {
	return xxx_messageInfo_ResponseA.Size(m)
}
func (m *ResponseA) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseA.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseA proto.InternalMessageInfo

func (m *ResponseA) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *ResponseA) GetParentSpanId() string {
	if m != nil {
		return m.ParentSpanId
	}
	return ""
}

func (m *ResponseA) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResponseA) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type ParamB struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParamB) Reset()         { *m = ParamB{} }
func (m *ParamB) String() string { return proto.CompactTextString(m) }
func (*ParamB) ProtoMessage()    {}
func (*ParamB) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff63324815cff01, []int{2}
}

func (m *ParamB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParamB.Unmarshal(m, b)
}
func (m *ParamB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParamB.Marshal(b, m, deterministic)
}
func (m *ParamB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamB.Merge(m, src)
}
func (m *ParamB) XXX_Size() int {
	return xxx_messageInfo_ParamB.Size(m)
}
func (m *ParamB) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamB.DiscardUnknown(m)
}

var xxx_messageInfo_ParamB proto.InternalMessageInfo

func (m *ParamB) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ParamB) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ResponseB struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	ParentSpanId         string   `protobuf:"bytes,2,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	Name                 string   `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,12,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseB) Reset()         { *m = ResponseB{} }
func (m *ResponseB) String() string { return proto.CompactTextString(m) }
func (*ResponseB) ProtoMessage()    {}
func (*ResponseB) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff63324815cff01, []int{3}
}

func (m *ResponseB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseB.Unmarshal(m, b)
}
func (m *ResponseB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseB.Marshal(b, m, deterministic)
}
func (m *ResponseB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseB.Merge(m, src)
}
func (m *ResponseB) XXX_Size() int {
	return xxx_messageInfo_ResponseB.Size(m)
}
func (m *ResponseB) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseB.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseB proto.InternalMessageInfo

func (m *ResponseB) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *ResponseB) GetParentSpanId() string {
	if m != nil {
		return m.ParentSpanId
	}
	return ""
}

func (m *ResponseB) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResponseB) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*ParamA)(nil), "pb.ParamA")
	proto.RegisterType((*ResponseA)(nil), "pb.ResponseA")
	proto.RegisterType((*ParamB)(nil), "pb.ParamB")
	proto.RegisterType((*ResponseB)(nil), "pb.ResponseB")
}

func init() { proto.RegisterFile("pb/test.proto", fileDescriptor_dff63324815cff01) }

var fileDescriptor_dff63324815cff01 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xd7, 0xfe, 0xf7, 0x5f, 0xdd, 0xbb, 0x4e, 0x24, 0x07, 0xa9, 0xf5, 0x32, 0x8a, 0xc8,
	0x4e, 0x29, 0x28, 0x88, 0xd7, 0x16, 0x3c, 0x0c, 0x14, 0xa4, 0xf3, 0x3e, 0xd2, 0xe5, 0xb5, 0x16,
	0x6c, 0x13, 0x92, 0x38, 0xd0, 0x2f, 0xe2, 0xd7, 0x95, 0xa6, 0xdd, 0xd6, 0xa1, 0x1e, 0xbd, 0xbd,
	0x79, 0xf8, 0x85, 0xdf, 0x03, 0x0f, 0x4c, 0x65, 0x1e, 0x1b, 0xd4, 0x86, 0x4a, 0x25, 0x8c, 0x20,
	0xae, 0xcc, 0xc3, 0xf3, 0x42, 0x88, 0xe2, 0x15, 0x63, 0x9b, 0xe4, 0x6f, 0xcf, 0x31, 0x56, 0xd2,
	0xbc, 0xb7, 0x40, 0x44, 0x61, 0xf4, 0xc8, 0x14, 0xab, 0x12, 0x42, 0x60, 0x58, 0xb3, 0x0a, 0x03,
	0x67, 0xe6, 0xcc, 0xc7, 0x99, 0xbd, 0xc9, 0x09, 0xfc, 0x63, 0x05, 0x06, 0xee, 0xcc, 0x99, 0xff,
	0xcf, 0x9a, 0x33, 0x52, 0x30, 0xce, 0x50, 0x4b, 0x51, 0x6b, 0x4c, 0xc8, 0x19, 0x1c, 0x19, 0xc5,
	0xd6, 0xb8, 0x2a, 0x79, 0xf7, 0xcd, 0xb3, 0xef, 0x05, 0x27, 0x17, 0x70, 0x2c, 0x99, 0xc2, 0xda,
	0xac, 0xb4, 0x64, 0x75, 0x03, 0xb8, 0x16, 0xf0, 0xdb, 0x74, 0x29, 0x59, 0xbd, 0xe0, 0x3b, 0xe7,
	0xe4, 0xbb, 0xd3, 0xdf, 0x3b, 0x6f, 0xba, 0x8e, 0xe9, 0x8f, 0x1d, 0x03, 0xf0, 0x18, 0xe7, 0x0a,
	0xb5, 0xee, 0x14, 0xdb, 0x67, 0xf4, 0xb1, 0xef, 0x9a, 0xfe, 0x4d, 0xd7, 0x9e, 0xdb, 0x3f, 0x70,
	0x5f, 0x7d, 0x3a, 0x30, 0x79, 0x42, 0x6d, 0x96, 0xa8, 0x36, 0xe5, 0x1a, 0xc9, 0x2d, 0x0c, 0xef,
	0xcb, 0x0d, 0x92, 0x53, 0xda, 0xae, 0x41, 0xb7, 0x6b, 0xd0, 0xbb, 0x66, 0x8d, 0xf0, 0x97, 0x3c,
	0x1a, 0x90, 0x4b, 0xf0, 0x1e, 0xd0, 0xbc, 0x08, 0x9e, 0x10, 0xa0, 0x32, 0xa7, 0xed, 0x5c, 0xe1,
	0xb4, 0xb9, 0x77, 0x53, 0xf4, 0xb9, 0xb4, 0xc7, 0xa5, 0x87, 0x5c, 0x1a, 0x0d, 0xf2, 0x91, 0x35,
	0x5c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x12, 0xc9, 0xe5, 0x13, 0x2a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	Live(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	MethodA(ctx context.Context, in *ParamA, opts ...grpc.CallOption) (*ResponseA, error)
	MethodB(ctx context.Context, in *ParamB, opts ...grpc.CallOption) (*ResponseB, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Live(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.TestService/Live", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) MethodA(ctx context.Context, in *ParamA, opts ...grpc.CallOption) (*ResponseA, error) {
	out := new(ResponseA)
	err := c.cc.Invoke(ctx, "/pb.TestService/MethodA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) MethodB(ctx context.Context, in *ParamB, opts ...grpc.CallOption) (*ResponseB, error) {
	out := new(ResponseB)
	err := c.cc.Invoke(ctx, "/pb.TestService/MethodB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	Live(context.Context, *empty.Empty) (*empty.Empty, error)
	MethodA(context.Context, *ParamA) (*ResponseA, error)
	MethodB(context.Context, *ParamB) (*ResponseB, error)
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) Live(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Live not implemented")
}
func (*UnimplementedTestServiceServer) MethodA(ctx context.Context, req *ParamA) (*ResponseA, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodA not implemented")
}
func (*UnimplementedTestServiceServer) MethodB(ctx context.Context, req *ParamB) (*ResponseB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodB not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Live_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Live(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TestService/Live",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Live(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_MethodA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamA)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).MethodA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TestService/MethodA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).MethodA(ctx, req.(*ParamA))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_MethodB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).MethodB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TestService/MethodB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).MethodB(ctx, req.(*ParamB))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Live",
			Handler:    _TestService_Live_Handler,
		},
		{
			MethodName: "MethodA",
			Handler:    _TestService_MethodA_Handler,
		},
		{
			MethodName: "MethodB",
			Handler:    _TestService_MethodB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/test.proto",
}