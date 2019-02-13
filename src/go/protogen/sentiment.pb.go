// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sentiment.proto

package protogen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SentimentRequest struct {
	Input string `protobuf:"bytes,1,opt,name=input" json:"input,omitempty"`
}

func (m *SentimentRequest) Reset()                    { *m = SentimentRequest{} }
func (m *SentimentRequest) String() string            { return proto.CompactTextString(m) }
func (*SentimentRequest) ProtoMessage()               {}
func (*SentimentRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *SentimentRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type SentimentResponse struct {
	Sentiment float32 `protobuf:"fixed32,1,opt,name=sentiment" json:"sentiment,omitempty"`
}

func (m *SentimentResponse) Reset()                    { *m = SentimentResponse{} }
func (m *SentimentResponse) String() string            { return proto.CompactTextString(m) }
func (*SentimentResponse) ProtoMessage()               {}
func (*SentimentResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SentimentResponse) GetSentiment() float32 {
	if m != nil {
		return m.Sentiment
	}
	return 0
}

func init() {
	proto.RegisterType((*SentimentRequest)(nil), "protogen.SentimentRequest")
	proto.RegisterType((*SentimentResponse)(nil), "protogen.SentimentResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SentimentService service

type SentimentServiceClient interface {
	Analyze(ctx context.Context, in *SentimentRequest, opts ...grpc.CallOption) (*SentimentResponse, error)
}

type sentimentServiceClient struct {
	cc *grpc.ClientConn
}

func NewSentimentServiceClient(cc *grpc.ClientConn) SentimentServiceClient {
	return &sentimentServiceClient{cc}
}

func (c *sentimentServiceClient) Analyze(ctx context.Context, in *SentimentRequest, opts ...grpc.CallOption) (*SentimentResponse, error) {
	out := new(SentimentResponse)
	err := grpc.Invoke(ctx, "/protogen.SentimentService/Analyze", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SentimentService service

type SentimentServiceServer interface {
	Analyze(context.Context, *SentimentRequest) (*SentimentResponse, error)
}

func RegisterSentimentServiceServer(s *grpc.Server, srv SentimentServiceServer) {
	s.RegisterService(&_SentimentService_serviceDesc, srv)
}

func _SentimentService_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentimentServiceServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protogen.SentimentService/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentimentServiceServer).Analyze(ctx, req.(*SentimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SentimentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protogen.SentimentService",
	HandlerType: (*SentimentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _SentimentService_Analyze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentiment.proto",
}

func init() { proto.RegisterFile("sentiment.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4e, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xe9,
	0xa9, 0x79, 0x4a, 0x1a, 0x5c, 0x02, 0xc1, 0x30, 0xc9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xbc, 0x82, 0xd2, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x08, 0x47, 0xc9, 0x90, 0x4b, 0x10, 0x49, 0x65, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x0c,
	0x17, 0x27, 0xdc, 0x6c, 0xb0, 0x72, 0xa6, 0x20, 0x84, 0x80, 0x51, 0x18, 0x92, 0xe1, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x4e, 0x5c, 0xec, 0x8e, 0x79, 0x89, 0x39, 0x95, 0x55, 0xa9,
	0x42, 0x52, 0x7a, 0x30, 0x67, 0xe8, 0xa1, 0xbb, 0x41, 0x4a, 0x1a, 0xab, 0x1c, 0xc4, 0xd6, 0x24,
	0x36, 0xb0, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x4b, 0x1c, 0x36, 0xd8, 0x00, 0x00,
	0x00,
}
