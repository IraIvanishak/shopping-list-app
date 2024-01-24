package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, name)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("error opening database connection:", err)
	}
	DB = db
}
