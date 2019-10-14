// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lend.proto

package lend

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type LendRequest struct {
	BookNumber           int32    `protobuf:"varint,1,opt,name=bookNumber,proto3" json:"bookNumber,omitempty"`
	StudentNumber        int32    `protobuf:"varint,2,opt,name=studentNumber,proto3" json:"studentNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LendRequest) Reset()         { *m = LendRequest{} }
func (m *LendRequest) String() string { return proto.CompactTextString(m) }
func (*LendRequest) ProtoMessage()    {}
func (*LendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e4e4748fbb082f, []int{0}
}

func (m *LendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LendRequest.Unmarshal(m, b)
}
func (m *LendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LendRequest.Marshal(b, m, deterministic)
}
func (m *LendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LendRequest.Merge(m, src)
}
func (m *LendRequest) XXX_Size() int {
	return xxx_messageInfo_LendRequest.Size(m)
}
func (m *LendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LendRequest proto.InternalMessageInfo

func (m *LendRequest) GetBookNumber() int32 {
	if m != nil {
		return m.BookNumber
	}
	return 0
}

func (m *LendRequest) GetStudentNumber() int32 {
	if m != nil {
		return m.StudentNumber
	}
	return 0
}

type LendResponse struct {
	BookNumber           int32    `protobuf:"varint,1,opt,name=bookNumber,proto3" json:"bookNumber,omitempty"`
	StudentNumber        int32    `protobuf:"varint,2,opt,name=studentNumber,proto3" json:"studentNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LendResponse) Reset()         { *m = LendResponse{} }
func (m *LendResponse) String() string { return proto.CompactTextString(m) }
func (*LendResponse) ProtoMessage()    {}
func (*LendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e4e4748fbb082f, []int{1}
}

func (m *LendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LendResponse.Unmarshal(m, b)
}
func (m *LendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LendResponse.Marshal(b, m, deterministic)
}
func (m *LendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LendResponse.Merge(m, src)
}
func (m *LendResponse) XXX_Size() int {
	return xxx_messageInfo_LendResponse.Size(m)
}
func (m *LendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LendResponse proto.InternalMessageInfo

func (m *LendResponse) GetBookNumber() int32 {
	if m != nil {
		return m.BookNumber
	}
	return 0
}

func (m *LendResponse) GetStudentNumber() int32 {
	if m != nil {
		return m.StudentNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*LendRequest)(nil), "lend.LendRequest")
	proto.RegisterType((*LendResponse)(nil), "lend.LendResponse")
}

func init() { proto.RegisterFile("lend.proto", fileDescriptor_e9e4e4748fbb082f) }

var fileDescriptor_e9e4e4748fbb082f = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x49, 0xcd, 0x4b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x82, 0xb9, 0xb8, 0x7d, 0x52,
	0xf3, 0x52, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xe4, 0xb8, 0xb8, 0x92, 0xf2, 0xf3,
	0xb3, 0xfd, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x90, 0x44,
	0x84, 0x54, 0xb8, 0x78, 0x8b, 0x4b, 0x4a, 0x53, 0x52, 0xf3, 0x4a, 0xa0, 0x4a, 0x98, 0xc0, 0x4a,
	0x50, 0x05, 0x95, 0x42, 0xb8, 0x78, 0x20, 0x86, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x52, 0xc7,
	0x54, 0x23, 0x1b, 0x88, 0x53, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x74, 0xb9, 0x58,
	0x40, 0x5c, 0x21, 0x41, 0x3d, 0xb0, 0xa7, 0x90, 0x7c, 0x21, 0x25, 0x84, 0x2c, 0x04, 0x71, 0x43,
	0x12, 0x1b, 0xd8, 0xd7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x70, 0xfd, 0x3d, 0x03,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LendServiceClient is the client API for LendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LendServiceClient interface {
	Lend(ctx context.Context, in *LendRequest, opts ...grpc.CallOption) (*LendResponse, error)
}

type lendServiceClient struct {
	cc *grpc.ClientConn
}

func NewLendServiceClient(cc *grpc.ClientConn) LendServiceClient {
	return &lendServiceClient{cc}
}

func (c *lendServiceClient) Lend(ctx context.Context, in *LendRequest, opts ...grpc.CallOption) (*LendResponse, error) {
	out := new(LendResponse)
	err := c.cc.Invoke(ctx, "/lend.LendService/Lend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LendServiceServer is the server API for LendService service.
type LendServiceServer interface {
	Lend(context.Context, *LendRequest) (*LendResponse, error)
}

// UnimplementedLendServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLendServiceServer struct {
}

func (*UnimplementedLendServiceServer) Lend(ctx context.Context, req *LendRequest) (*LendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lend not implemented")
}

func RegisterLendServiceServer(s *grpc.Server, srv LendServiceServer) {
	s.RegisterService(&_LendService_serviceDesc, srv)
}

func _LendService_Lend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LendServiceServer).Lend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lend.LendService/Lend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LendServiceServer).Lend(ctx, req.(*LendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LendService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lend.LendService",
	HandlerType: (*LendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lend",
			Handler:    _LendService_Lend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lend.proto",
}