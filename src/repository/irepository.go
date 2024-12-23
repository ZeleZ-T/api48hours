package repository

import (
	"api48hours/models"
)

type IRepository interface {
	EmailExists(email string) bool

	CreateUser(user models.User) error
	FindUserByEmail(email string) (models.User, error)
	ChangePassword(email, password string) error
	DeleteAccount(email string) error

	MapExistsToUser(email, name string) bool

	SaveMap(email string, params models.MapSaveParams) error
	FindMap(email, name string) (models.WorldMap, error)
	ChangeMapName(email, mapName, newName string) error
	DeleteMap(email, name string) error
}
