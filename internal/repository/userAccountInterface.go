package repository

type UserAccountInterface interface {
	DeleteUserAccountByEmail(email string) (int, error)
}

var UserAccountService UserAccountInterface