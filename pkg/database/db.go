package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // driver
)

// init db conn
func NewPostgresDB() (*sql.DB, error) {
	// load .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// read env variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// conn string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)

	// open db conn
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// verify conn
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database successfully")
	return db, nil
}
