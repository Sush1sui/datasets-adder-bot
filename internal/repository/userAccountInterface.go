package repository

import "github.com/Sush1sui/datasets_adder/internal/models"

type UserAccountInterface interface {
	GetAllUserAccounts() ([]models.UserAccount, error)
	GetUserByEmail(email string) (*models.UserAccount, error)
	DeleteUserAccountByEmail(email string) (int, error)
}

var UserAccountService UserAccountInterface