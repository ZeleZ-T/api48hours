package repository

import "api48hours/auth"

type IRepository interface {
	FindUserByEmail(email string) (auth.User, error)
	CreateUser(user auth.User) error
	ChangePassword(email, password string) error
	DeleteAccount(email string) error
}
