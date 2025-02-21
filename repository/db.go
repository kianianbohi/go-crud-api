// repository/db.go
package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	log.Println("Connected to MySQL")
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}
