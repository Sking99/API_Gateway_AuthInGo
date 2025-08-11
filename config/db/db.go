package config

import (
	config "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = config.GetString("DB_USER", "root")
	cfg.Passwd = config.GetString("DB_PASSWORD", "")
	cfg.Net = config.GetString("DB_NET", "tcp")
	cfg.Addr = config.GetString("DB_ADDR", "")
	cfg.DBName = config.GetString("DB_NAME", "auth_dev")

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error connecting DB")
		return nil, err
	}

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Error pinging to the DB")
		return nil, pingErr
	}

	fmt.Println("DB Connected Successfully")

	return db, nil
}
