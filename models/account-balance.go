package models

type AccountBalance struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	AccountID uint    `json:"accountId"`
	Balance   float64 `json:"balance"`
}
