package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

var MySqlRepo *MySQLRepository

func Start(cfg mysql.Config) error {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	fmt.Println("Connected!")

	err = error(nil)
	connection, err := sql.Open("mysql", "root:password@tcp")
	if err != nil {
		return err
	}
	MySqlRepo = NewMySQLRepository(
		connection,
	)
	return nil
}
