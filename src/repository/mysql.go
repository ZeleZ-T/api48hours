package repository

import (
	"api48hours/models"
	"database/sql"
	"github.com/KEINOS/go-noise"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) EmailExists(email string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *MySQLRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT email, password FROM users WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *MySQLRepository) CreateUser(user models.User) error {
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

func (r *MySQLRepository) MapExistsToUser(email, name string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM maps WHERE owner = ? AND name = ?)"
	err := r.db.QueryRow(query, email, name).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *MySQLRepository) SaveMap(email string, params models.MapSaveParams) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var paramsID int64
	query := "INSERT INTO map_params (seed, width, height, smoothness, water_level, perlin_noise) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.Exec(query, params.CreationParams.Seed, params.CreationParams.Width, params.CreationParams.Height, params.CreationParams.Smoothness, params.CreationParams.WaterSmoothness, params.CreationParams.NoiseType == noise.Perlin)
	if err != nil {
		tx.Rollback()
		return err
	}
	paramsID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	println(paramsID)

	query = "INSERT INTO maps (name, owner, params_id) VALUES (?, ?, ?)"
	_, err = tx.Exec(query, params.Name, email, paramsID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *MySQLRepository) FindMap(email, name string) (models.WorldMap, error) {
	var worldMap models.WorldMap
	var paramsID int64
	query := "SELECT params_id FROM maps WHERE owner = ? AND name = ?"
	err := r.db.QueryRow(query, email, name).Scan(&paramsID)
	if err != nil {
		return worldMap, err
	}

	query = "SELECT seed, width, height, smoothness, water_level, perlin_noise FROM map_params WHERE id = ?"
	var creationParams models.MapCreationParams
	err = r.db.QueryRow(query, paramsID).Scan(&creationParams.Seed, &creationParams.Width, &creationParams.Height, &creationParams.Smoothness, &creationParams.WaterSmoothness, &creationParams.NoiseType)
	if err != nil {
		return worldMap, err
	}

	worldMap.Seed = creationParams.Seed
	worldMap.Height = creationParams.Height
	worldMap.Width = creationParams.Width
	worldMap.MapData = make(map[int]map[int]float64)

	return worldMap, nil
}

func (r *MySQLRepository) ChangeMapName(email, mapName, newName string) error {
	query := "UPDATE maps SET name = ? WHERE owner = ? AND name = ?"
	_, err := r.db.Exec(query, newName, email, mapName)
	return err
}

func (r *MySQLRepository) DeleteMap(email, name string) error {
	query := "DELETE FROM maps WHERE owner = ? AND name = ?"
	_, err := r.db.Exec(query, email, name)
	return err
}
