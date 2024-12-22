package repository

import (
	"api48hours/models"
)

type IRepository interface {
	EmailExists(email string) bool
	FindUserByEmail(email string) (models.User, error)
	CreateUser(user models.User) error
	ChangePassword(email, password string) error
	DeleteAccount(email string) error
}
