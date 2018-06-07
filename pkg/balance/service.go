package balance

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/smblab/candidate/pkg/balance/pb"
	tpb "github.com/smblab/candidate/pkg/transactions/pb"
)

// Service an implementation of BalanceService
type Service struct {
	client tpb.TransactionServiceClient
}

// NewService initializes a new Balance Service
func NewService(tclient tpb.TransactionServiceClient) *Service {
	return &Service{
		client: tclient,
	}
}

// GetBalance ...
func (s *Service) GetBalance(ctx context.Context, r *pb.GetBalanceRequest) (*pb.BalanceResult, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "Not implemented")
}
