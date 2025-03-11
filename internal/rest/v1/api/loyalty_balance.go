package api

type LoyaltyBalance struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
	Count  int64 `json:"count"`
}
