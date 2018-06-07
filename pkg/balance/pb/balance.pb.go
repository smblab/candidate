// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/balance.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	api/balance.proto

It has these top-level messages:
	GetBalanceRequest
	Bucket
	BalanceResult
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import smblab_candidate_transactions_v1 "github.com/smblab/candidate/pkg/transactions/pb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetBalanceRequest struct {
	Name string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	From *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	To   *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	// Types that are valid to be assigned to Bucket:
	//	*GetBalanceRequest_Buckets
	//	*GetBalanceRequest_Duration
	Bucket isGetBalanceRequest_Bucket `protobuf_oneof:"bucket"`
}

func (m *GetBalanceRequest) Reset()                    { *m = GetBalanceRequest{} }
func (m *GetBalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceRequest) ProtoMessage()               {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isGetBalanceRequest_Bucket interface {
	isGetBalanceRequest_Bucket()
}

type GetBalanceRequest_Buckets struct {
	Buckets int32 `protobuf:"varint,4,opt,name=buckets,oneof"`
}
type GetBalanceRequest_Duration struct {
	Duration *google_protobuf1.Duration `protobuf:"bytes,5,opt,name=duration,oneof"`
}

func (*GetBalanceRequest_Buckets) isGetBalanceRequest_Bucket()  {}
func (*GetBalanceRequest_Duration) isGetBalanceRequest_Bucket() {}

func (m *GetBalanceRequest) GetBucket() isGetBalanceRequest_Bucket {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *GetBalanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetBalanceRequest) GetFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *GetBalanceRequest) GetTo() *google_protobuf.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *GetBalanceRequest) GetBuckets() int32 {
	if x, ok := m.GetBucket().(*GetBalanceRequest_Buckets); ok {
		return x.Buckets
	}
	return 0
}

func (m *GetBalanceRequest) GetDuration() *google_protobuf1.Duration {
	if x, ok := m.GetBucket().(*GetBalanceRequest_Duration); ok {
		return x.Duration
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GetBalanceRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GetBalanceRequest_OneofMarshaler, _GetBalanceRequest_OneofUnmarshaler, _GetBalanceRequest_OneofSizer, []interface{}{
		(*GetBalanceRequest_Buckets)(nil),
		(*GetBalanceRequest_Duration)(nil),
	}
}

func _GetBalanceRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GetBalanceRequest)
	// bucket
	switch x := m.Bucket.(type) {
	case *GetBalanceRequest_Buckets:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Buckets))
	case *GetBalanceRequest_Duration:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Duration); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GetBalanceRequest.Bucket has unexpected type %T", x)
	}
	return nil
}

func _GetBalanceRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GetBalanceRequest)
	switch tag {
	case 4: // bucket.buckets
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Bucket = &GetBalanceRequest_Buckets{int32(x)}
		return true, err
	case 5: // bucket.duration
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Duration)
		err := b.DecodeMessage(msg)
		m.Bucket = &GetBalanceRequest_Duration{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GetBalanceRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GetBalanceRequest)
	// bucket
	switch x := m.Bucket.(type) {
	case *GetBalanceRequest_Buckets:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Buckets))
	case *GetBalanceRequest_Duration:
		s := proto.Size(x.Duration)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Bucket struct {
	From   *google_protobuf.Timestamp              `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To     *google_protobuf.Timestamp              `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount *smblab_candidate_transactions_v1.Money `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *Bucket) Reset()                    { *m = Bucket{} }
func (m *Bucket) String() string            { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()               {}
func (*Bucket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Bucket) GetFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Bucket) GetTo() *google_protobuf.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Bucket) GetAmount() *smblab_candidate_transactions_v1.Money {
	if m != nil {
		return m.Amount
	}
	return nil
}

type BalanceResult struct {
	Name       string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	From       *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	To         *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	NumBuckets int32                      `protobuf:"varint,4,opt,name=num_buckets,json=numBuckets" json:"num_buckets,omitempty"`
	Buckets    []*Bucket                  `protobuf:"bytes,5,rep,name=buckets" json:"buckets,omitempty"`
}

func (m *BalanceResult) Reset()                    { *m = BalanceResult{} }
func (m *BalanceResult) String() string            { return proto.CompactTextString(m) }
func (*BalanceResult) ProtoMessage()               {}
func (*BalanceResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BalanceResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BalanceResult) GetFrom() *google_protobuf.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *BalanceResult) GetTo() *google_protobuf.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *BalanceResult) GetNumBuckets() int32 {
	if m != nil {
		return m.NumBuckets
	}
	return 0
}

func (m *BalanceResult) GetBuckets() []*Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBalanceRequest)(nil), "smblab.candidate.balance.v1.GetBalanceRequest")
	proto.RegisterType((*Bucket)(nil), "smblab.candidate.balance.v1.Bucket")
	proto.RegisterType((*BalanceResult)(nil), "smblab.candidate.balance.v1.BalanceResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BalanceService service

type BalanceServiceClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*BalanceResult, error)
}

type balanceServiceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceServiceClient(cc *grpc.ClientConn) BalanceServiceClient {
	return &balanceServiceClient{cc}
}

func (c *balanceServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*BalanceResult, error) {
	out := new(BalanceResult)
	err := grpc.Invoke(ctx, "/smblab.candidate.balance.v1.BalanceService/GetBalance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BalanceService service

type BalanceServiceServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*BalanceResult, error)
}

func RegisterBalanceServiceServer(s *grpc.Server, srv BalanceServiceServer) {
	s.RegisterService(&_BalanceService_serviceDesc, srv)
}

func _BalanceService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smblab.candidate.balance.v1.BalanceService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BalanceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smblab.candidate.balance.v1.BalanceService",
	HandlerType: (*BalanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _BalanceService_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/balance.proto",
}

func init() { proto.RegisterFile("api/balance.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0x23, 0x27, 0xf1, 0xb2, 0x37, 0x2c, 0x2c, 0x3a, 0x0c, 0xcf, 0x83, 0xc5, 0x78, 0x87,
	0x99, 0x1c, 0x64, 0x92, 0x1d, 0x76, 0x1a, 0x03, 0x33, 0x58, 0x2e, 0xbb, 0x78, 0x3d, 0xf5, 0x52,
	0x64, 0x47, 0x09, 0x6e, 0x2c, 0xc9, 0xb5, 0xa5, 0x40, 0xa1, 0x5f, 0xa7, 0x9f, 0xab, 0xd0, 0x4f,
	0x52, 0xe2, 0x7f, 0x69, 0x1a, 0x48, 0xdb, 0x4b, 0x6f, 0xe2, 0xd5, 0xf3, 0xbe, 0x8f, 0x1f, 0xfd,
	0x5e, 0xc3, 0x98, 0x66, 0x89, 0x1f, 0xd1, 0x94, 0x8a, 0x98, 0x91, 0x2c, 0x97, 0x4a, 0xe2, 0x2f,
	0x05, 0x8f, 0x52, 0x1a, 0x91, 0x98, 0x8a, 0x65, 0xb2, 0xa4, 0x8a, 0x91, 0xe6, 0x7e, 0x3b, 0xb3,
	0x27, 0x6b, 0x29, 0xd7, 0x29, 0xf3, 0x4b, 0x69, 0xa4, 0x57, 0xbe, 0x4a, 0x38, 0x2b, 0x14, 0xe5,
	0x59, 0xd5, 0x6d, 0x7f, 0x7d, 0x2a, 0x58, 0xea, 0x9c, 0xaa, 0x44, 0x8a, 0xfa, 0xfe, 0xd3, 0xce,
	0x50, 0xe5, 0x54, 0x14, 0x34, 0xde, 0x95, 0x8b, 0xaa, 0xee, 0xde, 0x23, 0x18, 0xff, 0x65, 0x2a,
	0xa8, 0xac, 0x42, 0x76, 0xa5, 0x59, 0xa1, 0x30, 0x86, 0x9e, 0xa0, 0x9c, 0x59, 0xc8, 0x41, 0xde,
	0xfb, 0xb0, 0x3c, 0x63, 0x02, 0xbd, 0x55, 0x2e, 0xb9, 0x65, 0x38, 0xc8, 0x1b, 0xce, 0x6d, 0x52,
	0x19, 0x92, 0xc6, 0x90, 0x9c, 0x35, 0x5f, 0x14, 0x96, 0x3a, 0x3c, 0x05, 0x43, 0x49, 0xab, 0xfb,
	0xac, 0xda, 0x50, 0x12, 0xdb, 0xf0, 0x2e, 0xd2, 0xf1, 0x86, 0xa9, 0xc2, 0xea, 0x39, 0xc8, 0xeb,
	0x2f, 0x3a, 0x61, 0x53, 0xc0, 0x3f, 0x61, 0xd0, 0x64, 0xb1, 0xfa, 0xe5, 0xb4, 0xcf, 0x47, 0xd3,
	0xfe, 0xd4, 0x82, 0x45, 0x27, 0x6c, 0xc5, 0xc1, 0x00, 0xcc, 0x6a, 0x86, 0x7b, 0x8b, 0xc0, 0x0c,
	0xca, 0x63, 0x9b, 0x02, 0xbd, 0x2a, 0x85, 0xf1, 0xa2, 0x14, 0xbf, 0xc1, 0xa4, 0x5c, 0x6a, 0xa1,
	0xea, 0xd4, 0xdf, 0xc9, 0x11, 0xd2, 0x03, 0x02, 0xdb, 0x19, 0xf9, 0x27, 0x05, 0xbb, 0x0e, 0xeb,
	0x36, 0xf7, 0x0e, 0xc1, 0x87, 0x96, 0x44, 0xa1, 0xd3, 0xb7, 0x07, 0x31, 0x81, 0xa1, 0xd0, 0xfc,
	0xe2, 0x00, 0x46, 0x08, 0x42, 0xf3, 0xa0, 0xa6, 0xf1, 0x6b, 0x4f, 0xaa, 0xef, 0x74, 0xbd, 0xe1,
	0xfc, 0x1b, 0x39, 0xb1, 0xb7, 0xa4, 0x6a, 0x6b, 0x61, 0xce, 0x6f, 0x60, 0x54, 0x07, 0xfc, 0xcf,
	0xf2, 0x6d, 0x12, 0x33, 0x7c, 0x09, 0xb0, 0xdf, 0x3f, 0x4c, 0x4e, 0x4e, 0x3b, 0x5a, 0x54, 0x7b,
	0x7a, 0xda, 0xfd, 0xf1, 0x5b, 0xba, 0x9d, 0xe0, 0xe3, 0xf9, 0x28, 0xdb, 0xac, 0x9b, 0xff, 0xce,
	0xcf, 0xa2, 0xc8, 0x2c, 0xdf, 0xe1, 0xc7, 0x43, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x3a, 0xc9,
	0x2c, 0x90, 0x03, 0x00, 0x00,
}
