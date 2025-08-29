package repository

type UserAccountInterface interface {
	DeleteUserAccountByEmail(email string) error
}

var UserAccountService UserAccountInterface