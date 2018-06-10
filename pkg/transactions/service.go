package transactions

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/smblab/candidate/pkg/transactions/pb"
)

var (
	errTransactionNotFound = grpc.Errorf(codes.NotFound, "Transaction not found")
	errAccountNotFound     = grpc.Errorf(codes.NotFound, "Account not found")
)

// Service an implementation of TransactionService
type Service struct {
	sync.RWMutex
	repo Repository
}

// NewService initializes a new Transaction Service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateTransaction ...
func (s *Service) CreateTransaction(ctx context.Context, r *pb.CreateTransactionRequest) (*pb.Transaction, error) {
	s.Lock()
	defer s.Unlock()

	var t Transaction
	tstamp, _ := ptypes.Timestamp(r.Transaction.TransactionTime)

	t = Transaction{
		Name:            r.Transaction.Name,
		Parent:          r.Transaction.Parent,
		Message:         r.Transaction.Message,
		RemoteAccount:   r.Transaction.RemoteAccount,
		Amount:          r.Transaction.Amount.Amount,
		CurrencyCode:    r.Transaction.Amount.CurrencyCode,
		TransactionTime: tstamp,
	}
    
	if err := s.repo.AddTransaction(&t); err != nil {
		return nil, err
	}

	return r.Transaction, nil
}

// GetTransaction ...
func (s *Service) GetTransaction(ctx context.Context, r *pb.GetTransactionRequest) (*pb.Transaction, error) {
	s.RLock()
	defer s.RUnlock()

	t, err := s.repo.GetTransaction(r.Name)
	if err != nil {
		return nil, errTransactionNotFound
	}
	tstamp, _ := ptypes.TimestampProto(t.TransactionTime)

	return &pb.Transaction{
		Name:          t.Name,
		Parent:        t.Parent,
		Message:       t.Message,
		RemoteAccount: t.RemoteAccount,
		Amount: &pb.Money{
			Amount:       t.Amount,
			CurrencyCode: t.CurrencyCode,
		},
		TransactionTime: tstamp,
	}, nil
}

// DeleteTransaction ...
func (s *Service) DeleteTransaction(ctx context.Context, r *pb.DeleteTransactionRequest) (*empty.Empty, error) {
	s.Lock()
	defer s.Unlock()

	if err := s.repo.DeleteTransaction(r.Name); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// ListTransactions ...
func (s *Service) ListTransactions(r *pb.ListTransactionsRequest, stream pb.TransactionService_ListTransactionsServer) error {
	s.RLock()
	defer s.RUnlock()

	transactions, err := s.repo.ListTransactions(r.Parent)
	if err != nil {
		return err
	}

	for _, t := range transactions {
		tstamp, _ := ptypes.TimestampProto(t.TransactionTime)
		if err := stream.Send(&pb.Transaction{
			Name:   t.Name,
			Parent: t.Parent,
			Amount: &pb.Money{
				Amount:       t.Amount,
				CurrencyCode: t.CurrencyCode,
			},
			TransactionTime: tstamp,
		}); err != nil {
			return err
		}
	}
	return nil
}
