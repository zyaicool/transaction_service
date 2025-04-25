package models

type TransactionLog struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	AccountID     uint    `json:"accountId"`
	AccountNumber string  `json:"accountNumber"`
	Amount        float64 `json:"amount"`
}
