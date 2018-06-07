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

type idempotencyKey struct {
	requestID string
	name      string
}

// Service an implementation of TransactionService
type Service struct {
	sync.RWMutex
	repo Repository

	reqs map[idempotencyKey]pb.Transaction
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

	ik := idempotencyKey{
		requestID: r.RequestId,
		name:      r.Transaction.Name,
	}

	if t, found := s.reqs[ik]; found {
		return &t, nil
	}

	var t Transaction
	tstamp, _ := ptypes.Timestamp(r.Transaction.TransactionTime)

	t = Transaction{
		Name:            r.Transaction.Name,
		Parent:          r.Transaction.Parent,
		Message:         r.Transaction.Message,
		Amount:          r.Transaction.Amount.Amount,
		CurrencyCode:    r.Transaction.Amount.CurrencyCode,
		TransactionTime: tstamp,
	}
	fmt.Println(t)

	if err := s.repo.AddTransaction(&t); err != nil {
		return nil, err
	}

	return r.Transaction, nil
}

// GetTransaction ...
func (s *Service) GetTransaction(ctx context.Context, r *pb.GetTransactionRequest) (*pb.Transaction, error) {
	t, err := s.repo.GetTransaction(r.Name)
	if err != nil {
		return nil, errTransactionNotFound
	}
	tstamp, _ := ptypes.TimestampProto(t.TransactionTime)

	return &pb.Transaction{
		Name:    t.Name,
		Parent:  t.Parent,
		Message: t.Message,
		Amount: &pb.Money{
			Amount:       t.Amount,
			CurrencyCode: t.CurrencyCode,
		},
		TransactionTime: tstamp,
	}, nil
}

// DeleteTransaction ...
func (s *Service) DeleteTransaction(ctx context.Context, r *pb.DeleteTransactionRequest) (*empty.Empty, error) {
	if err := s.repo.DeleteTransaction(r.Name); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// ListTransactions ...
func (s *Service) ListTransactions(r *pb.ListTransactionsRequest, stream pb.TransactionService_ListTransactionsServer) error {
	transactions, err := s.repo.ListTransactions(r.Parent)
	if err != nil {
		return err
	}

	for _, t := range transactions {
		tstamp, _ := ptypes.TimestampProto(t.TransactionTime)
		stream.Send(&pb.Transaction{
			Name:   t.Name,
			Parent: t.Parent,
			Amount: &pb.Money{
				Amount:       t.Amount,
				CurrencyCode: t.CurrencyCode,
			},
			TransactionTime: tstamp,
		})
	}
	return nil
}
