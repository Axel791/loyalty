package domains

import "github.com/Axel791/appkit"

type LoyaltyHistory struct {
	ID      int64
	UserID  int64
	OrderID int64
	Count   int64
}

func (v *LoyaltyHistory) ValidateUserID() error {
	if v.UserID <= 0 {
		return appkit.ValidationError("invalid userID")
	}
	return nil
}

func (v *LoyaltyHistory) ValidateOrderID() error {
	if v.OrderID <= 0 {
		return appkit.ValidationError("invalid orderID")
	}
	return nil
}

func (v *LoyaltyHistory) ValidateCount() error {
	if v.Count <= 0 {
		return appkit.ValidationError("invalid count")
	}
	return nil
}
