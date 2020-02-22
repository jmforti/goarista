// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gnmireverse.proto

package gnmireverse

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	gnmi "github.com/openconfig/gnmi/proto/gnmi"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da0910fdd411c63, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "gnmireverse.Empty")
}

func init() {
	proto.RegisterFile("gnmireverse.proto", fileDescriptor_7da0910fdd411c63)
}

var fileDescriptor_7da0910fdd411c63 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcf, 0xcb, 0xcd,
	0x2c, 0x4a, 0x2d, 0x4b, 0x2d, 0x2a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x92, 0x32, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2f,
	0x48, 0xcd, 0x4b, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x07, 0x29, 0xd1, 0x07, 0x2b, 0x87, 0x30,
	0x41, 0x04, 0x44, 0xbb, 0x12, 0x3b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x91, 0x3b, 0x17,
	0x77, 0xba, 0x9f, 0xaf, 0x67, 0x10, 0xc4, 0x24, 0x21, 0x0b, 0x2e, 0xf6, 0x80, 0xd2, 0xa4, 0x9c,
	0xcc, 0xe2, 0x0c, 0x21, 0x71, 0x3d, 0xb0, 0xfa, 0xe0, 0xd2, 0xa4, 0xe2, 0xe4, 0xa2, 0xcc, 0xa4,
	0xd4, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x29, 0x21, 0x3d, 0x64, 0xe7, 0x80, 0x8d,
	0xd1, 0x60, 0x4c, 0x62, 0x03, 0x1b, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x09, 0xca,
	0x0b, 0xac, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GNMIReverseClient is the client API for GNMIReverse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GNMIReverseClient interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (GNMIReverse_PublishClient, error)
}

type gNMIReverseClient struct {
	cc grpc.ClientConnInterface
}

func NewGNMIReverseClient(cc grpc.ClientConnInterface) GNMIReverseClient {
	return &gNMIReverseClient{cc}
}

func (c *gNMIReverseClient) Publish(ctx context.Context, opts ...grpc.CallOption) (GNMIReverse_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GNMIReverse_serviceDesc.Streams[0], "/gnmireverse.gNMIReverse/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &gNMIReversePublishClient{stream}
	return x, nil
}

type GNMIReverse_PublishClient interface {
	Send(*gnmi.SubscribeResponse) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type gNMIReversePublishClient struct {
	grpc.ClientStream
}

func (x *gNMIReversePublishClient) Send(m *gnmi.SubscribeResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gNMIReversePublishClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GNMIReverseServer is the server API for GNMIReverse service.
type GNMIReverseServer interface {
	Publish(GNMIReverse_PublishServer) error
}

// UnimplementedGNMIReverseServer can be embedded to have forward compatible implementations.
type UnimplementedGNMIReverseServer struct {
}

func (*UnimplementedGNMIReverseServer) Publish(srv GNMIReverse_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}

func RegisterGNMIReverseServer(s *grpc.Server, srv GNMIReverseServer) {
	s.RegisterService(&_GNMIReverse_serviceDesc, srv)
}

func _GNMIReverse_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GNMIReverseServer).Publish(&gNMIReversePublishServer{stream})
}

type GNMIReverse_PublishServer interface {
	SendAndClose(*Empty) error
	Recv() (*gnmi.SubscribeResponse, error)
	grpc.ServerStream
}

type gNMIReversePublishServer struct {
	grpc.ServerStream
}

func (x *gNMIReversePublishServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gNMIReversePublishServer) Recv() (*gnmi.SubscribeResponse, error) {
	m := new(gnmi.SubscribeResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GNMIReverse_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnmireverse.gNMIReverse",
	HandlerType: (*GNMIReverseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _GNMIReverse_Publish_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "gnmireverse.proto",
}
