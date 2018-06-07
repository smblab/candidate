package transactions

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/smblab/candidate/pkg/balance/pb"
)

// Service an implementation of BalanceService
type Service struct{}

// GetBalance ...
func (s *Service) GetBalance(ctx context.Context, r *pb.GetBalanceRequest) (*pb.BalanceResult, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "Not implemented")
}
