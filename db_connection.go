package repository_testing

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Database configuration
const (
	DB_USER     = "root"
	DB_PASSWORD = "secret"
	DB_NAME     = "golang"
)

// Database connection struct
type DbConnection struct{}

// GetConnection will return database connection instance
func (self *DbConnection) GetConnection() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	return db
}
