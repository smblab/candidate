// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/transactions.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	api/transactions.proto

It has these top-level messages:
	Money
	Transaction
	CreateTransactionRequest
	GetTransactionRequest
	DeleteTransactionRequest
	ListTransactionsRequest
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

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

type Money struct {
	Amount       float64 `protobuf:"fixed64,1,opt,name=amount" json:"amount,omitempty"`
	CurrencyCode string  `protobuf:"bytes,2,opt,name=currency_code,json=currencyCode" json:"currency_code,omitempty"`
}

func (m *Money) Reset()                    { *m = Money{} }
func (m *Money) String() string            { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()               {}
func (*Money) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Money) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Money) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

type Transaction struct {
	Name            string                      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parent          string                      `protobuf:"bytes,2,opt,name=parent" json:"parent,omitempty"`
	Message         string                      `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	RemoteAccount   string                      `protobuf:"bytes,4,opt,name=remote_account,json=remoteAccount" json:"remote_account,omitempty"`
	Amount          *Money                      `protobuf:"bytes,5,opt,name=amount" json:"amount,omitempty"`
	TransactionTime *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=transaction_time,json=transactionTime" json:"transaction_time,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Transaction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Transaction) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Transaction) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Transaction) GetRemoteAccount() string {
	if m != nil {
		return m.RemoteAccount
	}
	return ""
}

func (m *Transaction) GetAmount() *Money {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Transaction) GetTransactionTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.TransactionTime
	}
	return nil
}

type CreateTransactionRequest struct {
	RequestId   string       `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction" json:"transaction,omitempty"`
}

func (m *CreateTransactionRequest) Reset()                    { *m = CreateTransactionRequest{} }
func (m *CreateTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTransactionRequest) ProtoMessage()               {}
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateTransactionRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *CreateTransactionRequest) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type GetTransactionRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parent string `protobuf:"bytes,2,opt,name=parent" json:"parent,omitempty"`
}

func (m *GetTransactionRequest) Reset()                    { *m = GetTransactionRequest{} }
func (m *GetTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTransactionRequest) ProtoMessage()               {}
func (*GetTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetTransactionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetTransactionRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

type DeleteTransactionRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parent string `protobuf:"bytes,2,opt,name=parent" json:"parent,omitempty"`
}

func (m *DeleteTransactionRequest) Reset()                    { *m = DeleteTransactionRequest{} }
func (m *DeleteTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTransactionRequest) ProtoMessage()               {}
func (*DeleteTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteTransactionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteTransactionRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

type ListTransactionsRequest struct {
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
}

func (m *ListTransactionsRequest) Reset()                    { *m = ListTransactionsRequest{} }
func (m *ListTransactionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTransactionsRequest) ProtoMessage()               {}
func (*ListTransactionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListTransactionsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func init() {
	proto.RegisterType((*Money)(nil), "smblab.candidate.transactions.v1.Money")
	proto.RegisterType((*Transaction)(nil), "smblab.candidate.transactions.v1.Transaction")
	proto.RegisterType((*CreateTransactionRequest)(nil), "smblab.candidate.transactions.v1.CreateTransactionRequest")
	proto.RegisterType((*GetTransactionRequest)(nil), "smblab.candidate.transactions.v1.GetTransactionRequest")
	proto.RegisterType((*DeleteTransactionRequest)(nil), "smblab.candidate.transactions.v1.DeleteTransactionRequest")
	proto.RegisterType((*ListTransactionsRequest)(nil), "smblab.candidate.transactions.v1.ListTransactionsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TransactionService service

type TransactionServiceClient interface {
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (TransactionService_ListTransactionsClient, error)
}

type transactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionServiceClient(cc *grpc.ClientConn) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := grpc.Invoke(ctx, "/smblab.candidate.transactions.v1.TransactionService/CreateTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := grpc.Invoke(ctx, "/smblab.candidate.transactions.v1.TransactionService/GetTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/smblab.candidate.transactions.v1.TransactionService/DeleteTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (TransactionService_ListTransactionsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TransactionService_serviceDesc.Streams[0], c.cc, "/smblab.candidate.transactions.v1.TransactionService/ListTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionServiceListTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransactionService_ListTransactionsClient interface {
	Recv() (*Transaction, error)
	grpc.ClientStream
}

type transactionServiceListTransactionsClient struct {
	grpc.ClientStream
}

func (x *transactionServiceListTransactionsClient) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TransactionService service

type TransactionServiceServer interface {
	CreateTransaction(context.Context, *CreateTransactionRequest) (*Transaction, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error)
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*google_protobuf1.Empty, error)
	ListTransactions(*ListTransactionsRequest, TransactionService_ListTransactionsServer) error
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smblab.candidate.transactions.v1.TransactionService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smblab.candidate.transactions.v1.TransactionService/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smblab.candidate.transactions.v1.TransactionService/DeleteTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTransactionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactionServiceServer).ListTransactions(m, &transactionServiceListTransactionsServer{stream})
}

type TransactionService_ListTransactionsServer interface {
	Send(*Transaction) error
	grpc.ServerStream
}

type transactionServiceListTransactionsServer struct {
	grpc.ServerStream
}

func (x *transactionServiceListTransactionsServer) Send(m *Transaction) error {
	return x.ServerStream.SendMsg(m)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smblab.candidate.transactions.v1.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionService_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _TransactionService_GetTransaction_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _TransactionService_DeleteTransaction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTransactions",
			Handler:       _TransactionService_ListTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/transactions.proto",
}

func init() { proto.RegisterFile("api/transactions.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0x86, 0x26, 0x28, 0x13, 0x5a, 0xda, 0x45, 0x0d, 0x56, 0x00, 0x11, 0x19, 0x10, 0x21,
	0x08, 0x9b, 0xa4, 0x07, 0xd4, 0x5e, 0x10, 0xa4, 0x05, 0x21, 0x81, 0x90, 0x42, 0x4f, 0x5c, 0xa2,
	0x8d, 0x3d, 0x44, 0x16, 0xb1, 0xd7, 0x78, 0x37, 0x91, 0xa2, 0xaa, 0x97, 0x1e, 0xb9, 0x72, 0xe3,
	0x11, 0xe0, 0x09, 0x78, 0x03, 0xee, 0xbc, 0x02, 0x0f, 0x82, 0xbc, 0xb6, 0xdb, 0x75, 0x93, 0xc8,
	0xf8, 0xe6, 0xf9, 0xfb, 0xe6, 0x9b, 0x6f, 0x66, 0x0d, 0x4d, 0x16, 0x7a, 0xb6, 0x8c, 0x58, 0x20,
	0x98, 0x23, 0x3d, 0x1e, 0x08, 0x2b, 0x8c, 0xb8, 0xe4, 0xb4, 0x2d, 0xfc, 0xf1, 0x94, 0x8d, 0x2d,
	0x87, 0x05, 0xae, 0xe7, 0x32, 0x89, 0x56, 0x2e, 0x69, 0xde, 0x6b, 0xdd, 0x9e, 0x70, 0x3e, 0x99,
	0xa2, 0x1d, 0x03, 0xb0, 0x20, 0xe0, 0x92, 0x69, 0xf5, 0xad, 0x5b, 0x69, 0x54, 0x59, 0xe3, 0xd9,
	0x27, 0x1b, 0xfd, 0x50, 0x2e, 0xd2, 0xe0, 0xdd, 0xcb, 0x41, 0xe9, 0xf9, 0x28, 0x24, 0xf3, 0xc3,
	0x24, 0xc1, 0x3c, 0x84, 0xea, 0x3b, 0x1e, 0xe0, 0x82, 0x36, 0xa1, 0xc6, 0x7c, 0x3e, 0x0b, 0xa4,
	0x41, 0xda, 0xa4, 0x43, 0x86, 0xa9, 0x45, 0xef, 0xc1, 0xa6, 0x33, 0x8b, 0x22, 0x0c, 0x9c, 0xc5,
	0xc8, 0xe1, 0x2e, 0x1a, 0x95, 0x36, 0xe9, 0xd4, 0x87, 0xd7, 0x32, 0xe7, 0x80, 0xbb, 0x68, 0x9e,
	0x55, 0xa0, 0x71, 0x7c, 0xc1, 0x9a, 0x52, 0xd8, 0x08, 0x98, 0x8f, 0x0a, 0xaa, 0x3e, 0x54, 0xdf,
	0x71, 0x83, 0x90, 0x45, 0x18, 0xc8, 0x14, 0x21, 0xb5, 0xa8, 0x01, 0x57, 0x7d, 0x14, 0x82, 0x4d,
	0xd0, 0xb8, 0xa2, 0x02, 0x99, 0x49, 0x1f, 0xc0, 0x56, 0x84, 0x3e, 0x97, 0x38, 0x62, 0x8e, 0xa3,
	0xa8, 0x6d, 0xa8, 0x84, 0xcd, 0xc4, 0xfb, 0x22, 0x71, 0xd2, 0xe7, 0xe7, 0xcc, 0xab, 0x6d, 0xd2,
	0x69, 0xf4, 0x1f, 0x5a, 0x45, 0x8a, 0x5a, 0x6a, 0xe4, 0xf3, 0x11, 0x8f, 0x60, 0x5b, 0x4b, 0x18,
	0xc5, 0x12, 0x19, 0x35, 0x05, 0xd5, 0xb2, 0x12, 0xfd, 0xac, 0x4c, 0x3f, 0xeb, 0x38, 0xd3, 0x6f,
	0x78, 0x5d, 0xab, 0x89, 0xbd, 0xe6, 0x57, 0x02, 0xc6, 0x20, 0x42, 0x26, 0x51, 0x93, 0x62, 0x88,
	0x5f, 0x66, 0x28, 0x24, 0xbd, 0x03, 0x10, 0x25, 0x9f, 0x23, 0xcf, 0x4d, 0x75, 0xa9, 0xa7, 0x9e,
	0x37, 0x2e, 0x7d, 0x0f, 0x0d, 0x0d, 0x4e, 0x29, 0xd4, 0xe8, 0x3f, 0x29, 0x1e, 0x44, 0xef, 0xa4,
	0x23, 0x98, 0x03, 0xd8, 0x7d, 0x8d, 0x72, 0x05, 0x91, 0x12, 0xab, 0x31, 0x5f, 0x81, 0x71, 0x88,
	0x53, 0x5c, 0x39, 0x50, 0x19, 0x9c, 0x1e, 0xdc, 0x7c, 0xeb, 0x09, 0x9d, 0x8d, 0xc8, 0x60, 0x2e,
	0x4a, 0x88, 0x5e, 0xd2, 0xff, 0x5d, 0x05, 0xaa, 0xe5, 0x7f, 0xc0, 0x68, 0xee, 0x39, 0x48, 0x7f,
	0x11, 0xd8, 0x59, 0xd2, 0x98, 0x1e, 0x14, 0x0b, 0xb5, 0x6e, 0x31, 0xad, 0x72, 0x22, 0x9b, 0xfb,
	0x67, 0x7f, 0xfe, 0x7e, 0xab, 0xec, 0x99, 0x96, 0x3d, 0xef, 0xd9, 0xe9, 0x59, 0x0a, 0xfb, 0x44,
	0x2b, 0xb1, 0x92, 0x09, 0x4e, 0x73, 0x6f, 0xfd, 0x80, 0x74, 0xe9, 0x4f, 0x02, 0x5b, 0xf9, 0x9d,
	0xd0, 0x67, 0xc5, 0xcd, 0x57, 0x6e, 0xb1, 0x2c, 0xeb, 0x9e, 0x62, 0xfd, 0x98, 0x3e, 0xca, 0xb3,
	0x5e, 0xc5, 0xd4, 0x3e, 0x89, 0x57, 0x79, 0x4a, 0xbf, 0x13, 0xd8, 0x59, 0x5a, 0xfe, 0xff, 0x28,
	0xbd, 0xee, 0x62, 0x5a, 0xcd, 0xa5, 0xc7, 0x74, 0x14, 0xff, 0xa9, 0x32, 0x72, 0xdd, 0x12, 0xe4,
	0x7e, 0x10, 0xd8, 0xbe, 0x7c, 0x51, 0x74, 0xbf, 0x98, 0xdb, 0x9a, 0x2b, 0x2c, 0x2b, 0x67, 0x57,
	0x31, 0xbe, 0x4f, 0xcd, 0x62, 0xc6, 0x4f, 0xc9, 0xcb, 0xdd, 0x8f, 0x37, 0xc2, 0xcf, 0x93, 0xfc,
	0x1c, 0xe1, 0x78, 0x5c, 0x53, 0x32, 0xec, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x10, 0x3e,
	0x42, 0x18, 0x06, 0x00, 0x00,
}
