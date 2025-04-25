package services

import (
	"errors"
	"fmt"
	"math/rand"
	"transaction-service/models"
	"transaction-service/repositories"

	"gorm.io/gorm"
)

type AccountService struct {
	Repo *repositories.AccountRepository
	DB   *gorm.DB
}

func NewAccountService(repo *repositories.AccountRepository, db *gorm.DB) *AccountService {
	return &AccountService{Repo: repo, DB: db}
}

func generateAccountNumber() string {
	return fmt.Sprintf("99%08d", rand.Intn(100000000))
}

func (s *AccountService) Register(name, nik, phone string) (string, error) {
	existing, err := s.Repo.GetByNIKOrPhone(nik, phone)
	if err == nil && existing.ID != 0 {
		return "", errors.New("NIK atau No HP sudah terdaftar")
	}

	accountNumber := generateAccountNumber()

	account := models.Account{
		Name:          name,
		NIK:           nik,
		PhoneNo:       phone,
		AccountNumber: accountNumber,
	}

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&account).Error; err != nil {
			return err
		}

		balance := models.AccountBalance{
			AccountID: account.ID,
			Balance:   0,
		}

		return tx.Create(&balance).Error
	})

	return accountNumber, err
}

func (s *AccountService) Deposit(accountNumber string, amount float64) (float64, error) {
	var account models.Account
	if err := s.DB.Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		return 0, errors.New("Nomor rekening tidak ditemukan")
	}

	return s.updateBalance(account, amount)
}

func (s *AccountService) Withdraw(accountNumber string, amount float64) (float64, error) {
	var account models.Account
	if err := s.DB.Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		return 0, errors.New("Nomor rekening tidak ditemukan")
	}

	var balance models.AccountBalance
	if err := s.DB.Where("account_id = ?", account.ID).First(&balance).Error; err != nil {
		return 0, err
	}

	if balance.Balance < amount {
		return balance.Balance, errors.New("Saldo tidak cukup")
	}

	return s.updateBalance(account, -amount)
}

func (s *AccountService) updateBalance(account models.Account, amount float64) (float64, error) {
	returnAmount := 0.0

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		var balance models.AccountBalance
		if err := tx.Where("account_id = ?", account.ID).First(&balance).Error; err != nil {
			return err
		}

		balance.Balance += amount
		returnAmount = balance.Balance

		if err := tx.Save(&balance).Error; err != nil {
			return err
		}

		log := models.TransactionLog{
			AccountID:     account.ID,
			AccountNumber: account.AccountNumber,
			Amount:        amount,
		}

		return tx.Create(&log).Error
	})

	return returnAmount, err
}

func (s *AccountService) GetBalance(accountNumber string) (float64, error) {
	var account models.Account
	if err := s.DB.Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		return 0, errors.New("Nomor rekening tidak ditemukan")
	}

	var balance models.AccountBalance
	if err := s.DB.Where("account_id = ?", account.ID).First(&balance).Error; err != nil {
		return 0, err
	}

	return balance.Balance, nil
}
