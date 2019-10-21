// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mymodels.proto

package mymodels

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

type Subscribe struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Agents               []string `protobuf:"bytes,2,rep,name=agents,proto3" json:"agents,omitempty"`
	Queues               []string `protobuf:"bytes,3,rep,name=queues,proto3" json:"queues,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subscribe) Reset()         { *m = Subscribe{} }
func (m *Subscribe) String() string { return proto.CompactTextString(m) }
func (*Subscribe) ProtoMessage()    {}
func (*Subscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_e47b40523705aad6, []int{0}
}

func (m *Subscribe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscribe.Unmarshal(m, b)
}
func (m *Subscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscribe.Marshal(b, m, deterministic)
}
func (m *Subscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscribe.Merge(m, src)
}
func (m *Subscribe) XXX_Size() int {
	return xxx_messageInfo_Subscribe.Size(m)
}
func (m *Subscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscribe.DiscardUnknown(m)
}

var xxx_messageInfo_Subscribe proto.InternalMessageInfo

func (m *Subscribe) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *Subscribe) GetAgents() []string {
	if m != nil {
		return m.Agents
	}
	return nil
}

func (m *Subscribe) GetQueues() []string {
	if m != nil {
		return m.Queues
	}
	return nil
}

type Event struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	TopicId              string   `protobuf:"bytes,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Application          string   `protobuf:"bytes,3,opt,name=application,proto3" json:"application,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e47b40523705aad6, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Event) GetTopicId() string {
	if m != nil {
		return m.TopicId
	}
	return ""
}

func (m *Event) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TestEvent struct {
	Somestuff            string   `protobuf:"bytes,1,opt,name=somestuff,proto3" json:"somestuff,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestEvent) Reset()         { *m = TestEvent{} }
func (m *TestEvent) String() string { return proto.CompactTextString(m) }
func (*TestEvent) ProtoMessage()    {}
func (*TestEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e47b40523705aad6, []int{2}
}

func (m *TestEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestEvent.Unmarshal(m, b)
}
func (m *TestEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestEvent.Marshal(b, m, deterministic)
}
func (m *TestEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestEvent.Merge(m, src)
}
func (m *TestEvent) XXX_Size() int {
	return xxx_messageInfo_TestEvent.Size(m)
}
func (m *TestEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TestEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TestEvent proto.InternalMessageInfo

func (m *TestEvent) GetSomestuff() string {
	if m != nil {
		return m.Somestuff
	}
	return ""
}

func init() {
	proto.RegisterType((*Subscribe)(nil), "mymodels.Subscribe")
	proto.RegisterType((*Event)(nil), "mymodels.Event")
	proto.RegisterType((*TestEvent)(nil), "mymodels.TestEvent")
}

func init() { proto.RegisterFile("mymodels.proto", fileDescriptor_e47b40523705aad6) }

var fileDescriptor_e47b40523705aad6 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4b, 0xf4, 0x30,
	0x10, 0xc6, 0xdf, 0x6c, 0xf7, 0x2b, 0xf3, 0xfa, 0x45, 0x14, 0xc9, 0x8a, 0x87, 0xd2, 0x53, 0x05,
	0x29, 0xa2, 0x78, 0x97, 0x05, 0x59, 0x3c, 0x14, 0x96, 0xea, 0x5d, 0xd2, 0x76, 0x76, 0x09, 0xb4,
	0x4d, 0x6c, 0x52, 0xb1, 0x57, 0xff, 0x72, 0x69, 0x5a, 0xbb, 0xde, 0x9e, 0xdf, 0x6f, 0x26, 0x61,
	0x78, 0xe0, 0xa4, 0x6c, 0x4b, 0x95, 0x63, 0x61, 0x22, 0x5d, 0x2b, 0xab, 0xd8, 0xf2, 0x97, 0x83,
	0x18, 0xe8, 0x6b, 0x93, 0x9a, 0xac, 0x96, 0x29, 0xb2, 0x33, 0xf0, 0x84, 0xd6, 0x9c, 0xf8, 0x24,
	0xa4, 0x49, 0x17, 0xd9, 0x25, 0xcc, 0xc5, 0x1e, 0x2b, 0x6b, 0xf8, 0xc4, 0xf7, 0x42, 0x9a, 0x0c,
	0xd4, 0xf9, 0x8f, 0x06, 0x1b, 0x34, 0xdc, 0xeb, 0x7d, 0x4f, 0xc1, 0x37, 0x81, 0xd9, 0xf3, 0x27,
	0x56, 0x96, 0x5d, 0xc0, 0xcc, 0x2a, 0x2d, 0xb3, 0xe1, 0xb7, 0x1e, 0xd8, 0x0a, 0x96, 0x2e, 0xbc,
	0xcb, 0x9c, 0x4f, 0xdc, 0x60, 0xe1, 0xf8, 0x25, 0x67, 0x3e, 0xfc, 0x17, 0x5a, 0x17, 0x32, 0x13,
	0x56, 0xaa, 0x8a, 0x7b, 0x6e, 0xfa, 0x57, 0x31, 0x06, 0x53, 0xdb, 0x6a, 0xe4, 0x53, 0x37, 0x72,
	0xb9, 0x73, 0xb9, 0xb0, 0x82, 0xcf, 0x7c, 0x12, 0x1e, 0x25, 0x2e, 0x07, 0x37, 0x40, 0xdf, 0xd0,
	0xd8, 0xfe, 0x8e, 0x6b, 0xa0, 0x46, 0x95, 0x68, 0x6c, 0xb3, 0xdb, 0x0d, 0xb7, 0x1c, 0xc4, 0xfd,
	0x13, 0x2c, 0x36, 0x35, 0xa2, 0xc5, 0x9a, 0x3d, 0x02, 0xdd, 0x60, 0xff, 0xc8, 0xb0, 0xf3, 0x68,
	0x6c, 0x6c, 0xac, 0xe7, 0xea, 0xf4, 0x20, 0xdd, 0x5a, 0xf0, 0xef, 0x8e, 0xac, 0x6f, 0x61, 0x25,
	0x55, 0xb4, 0xaf, 0x75, 0x16, 0xe1, 0x97, 0x28, 0x75, 0x81, 0x66, 0x5c, 0x5b, 0x1f, 0xc7, 0x6d,
	0xec, 0xd2, 0xb6, 0xab, 0x7d, 0x4b, 0xd2, 0xb9, 0xeb, 0xff, 0xe1, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x40, 0x3a, 0x84, 0x88, 0x91, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	GetEvents(ctx context.Context, in *Subscribe, opts ...grpc.CallOption) (Greeter_GetEventsClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetEvents(ctx context.Context, in *Subscribe, opts ...grpc.CallOption) (Greeter_GetEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/mymodels.Greeter/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_GetEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type greeterGetEventsClient struct {
	grpc.ClientStream
}

func (x *greeterGetEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	GetEvents(*Subscribe, Greeter_GetEventsServer) error
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) GetEvents(req *Subscribe, srv Greeter_GetEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Subscribe)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).GetEvents(m, &greeterGetEventsServer{stream})
}

type Greeter_GetEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type greeterGetEventsServer struct {
	grpc.ServerStream
}

func (x *greeterGetEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mymodels.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _Greeter_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mymodels.proto",
}
