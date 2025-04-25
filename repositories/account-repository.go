package repositories

import (
	"transaction-service/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) CreateAccount(account *models.Account) error {
	return r.DB.Create(account).Error
}

func (r *AccountRepository) GetByNIKOrPhone(nik, phone string) (*models.Account, error) {
	var account models.Account
	result := r.DB.Where("nik = ? OR phone_no = ?", nik, phone).First(&account)
	return &account, result.Error
}

func (r *AccountRepository) GetByAccountNumber(accNum string) (*models.Account, error) {
	var account models.Account
	result := r.DB.Where("account_number = ?", accNum).First(&account)
	return &account, result.Error
}
