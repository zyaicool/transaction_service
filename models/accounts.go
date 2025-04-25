package models

type Account struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	AccountNumber  string         `gorm:"unique" json:"accountNumber"`
	Name           string         `json:"name"`
	NIK            string         `gorm:"unique" json:"nik"`
	PhoneNo        string         `gorm:"unique" json:"phoneNo"`
	AccountBalance AccountBalance `gorm:"foreignKey:AccountID"`
}
