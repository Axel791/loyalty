package repositories

import (
	"context"
	"github.com/Axel791/loyalty/interanal/domains"
)

type LoyaltyRepositories interface {
	GetUserBalance(ctx context.Context, userID int64) (domains.LoyaltyBalance, error)
	ConclusionUserBalance(ctx context.Context, userBalance domains.LoyaltyBalance) error
	InputUserBalance(ctx context.Context, userBalance domains.LoyaltyBalance) error
}
