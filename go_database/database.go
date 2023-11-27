package go_database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_mysql?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Second)

	return db
}