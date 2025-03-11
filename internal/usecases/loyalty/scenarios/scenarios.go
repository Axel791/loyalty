package scenarios

import (
	"context"

	"github.com/Axel791/loyalty/internal/domains"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
)

type InputUserBalanceUseCase interface {
	Execute(ctx context.Context, balance dto.LoyaltyBalance) error
}

type GetUserBalanceUseCase interface {
	Execute(ctx context.Context, userID int64) (dto.LoyaltyBalance, error)
}

type ConclusionUserBalanceUseCase interface {
	ConclusionUserBalance(ctx context.Context, userBalance domains.LoyaltyBalance) error
}

type CreateUserBalanceUseCase interface {
	CreateLoyaltyBalance(ctx context.Context, userID int64) error
}
