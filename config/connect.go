package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect_db() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/backend_gdsc")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
