package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	createTables()

	log.Println("DB connection success")
}

func createTables() {
	userTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL
    );
    `

	_, err := DB.Exec(userTable)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}
}

func Execute(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return DB.QueryRow(query, args...)
}
