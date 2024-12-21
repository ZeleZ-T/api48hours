package repository

import (
	"api48hours/auth"
	"database/sql"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) FindUserByEmail(email string) (auth.User, error) {
	var user auth.User
	query := "SELECT email, password FROM users WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *MySQLRepository) CreateUser(user auth.User) error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	_, err := r.db.Exec(query, user.Email, user.Password)
	return err
}

func (r *MySQLRepository) ChangePassword(email, password string) error {
	query := "UPDATE users SET password = ? WHERE email = ?"
	_, err := r.db.Exec(query, password, email)
	return err
}

func (r *MySQLRepository) DeleteAccount(email string) error {
	query := "DELETE FROM users WHERE email = ?"
	_, err := r.db.Exec(query, email)
	return err
}
