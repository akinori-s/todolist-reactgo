package database

import (
	"fmt"
	// "os"
	"database/sql"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")
	// wip: currently running on local db
	host := "localhost"
	port := "5432"
	user := "golang_usr"
	password := "golang_pass"
	dbname := "golang_db"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
