package transactions

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/smblab/candidate/pkg/transactions/pb"
)

// Service an implementation of TransactionService
type Service struct{}

// CreateTransaction ...
func (s *Service) CreateTransaction(ctx context.Context, r *pb.CreateTransactionRequest) (*pb.Transaction, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "Not implemented")
}

// GetTransaction ...
func (s *Service) GetTransaction(ctx context.Context, r *pb.GetTransactionRequest) (*pb.Transaction, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "Not implemented")
}

// DeleteTransaction ...
func (s *Service) DeleteTransaction(ctx context.Context, r *pb.DeleteTransactionRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "Not implemented")
}

// ListTransactions ...
func (s *Service) ListTransactions(ctx context.Context, r *pb.ListTransactionsRequest) (pb.TransactionService_ListTransactionsClient, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "Not implemented")
}
