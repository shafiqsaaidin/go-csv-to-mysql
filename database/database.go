package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shafiqsaaidin/go-csv-to-mysql/config"
)

var DB *sql.DB

func Connect() error {
	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Config("DB_USER"),
		config.Config("DB_PASS"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME")))
	if err != nil {
		return err
	}

	fmt.Println("Connected to database")
	return nil
}
