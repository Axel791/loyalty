package v1

import (
	"context"
	"fmt"

	"github.com/Axel791/loyalty/internal/grpc/v1/pb"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/scenarios"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoyaltyServer struct {
	pb.UnimplementedLoyaltyServiceServer

	createLoyaltyBalanceHandler *scenarios.CreateLoyaltyBalanceHandler
}

func NewLoyaltyServer(
	handler *scenarios.CreateLoyaltyBalanceHandler,
) *LoyaltyServer {
	return &LoyaltyServer{
		createLoyaltyBalanceHandler: handler,
	}
}

// CreateLoyaltyBalance - Реализация метода CreateLoyaltyBalance из .proto
func (s *LoyaltyServer) CreateLoyaltyBalance(
	ctx context.Context,
	req *pb.CreateLoyaltyBalanceRequest,
) (*pb.CreateLoyaltyBalanceResponse, error) {
	err := s.createLoyaltyBalanceHandler.CreateLoyaltyBalance(ctx, req.GetUserId())
	if err != nil {
		return &pb.CreateLoyaltyBalanceResponse{
				Success:      false,
				ErrorMessage: fmt.Sprintf("failed to create loyalty balance: %v", err),
			},
			status.Errorf(codes.Internal, "failed to create loyalty balance: %v", err)
	}

	return &pb.CreateLoyaltyBalanceResponse{
		Success: true,
	}, nil
}
