package db

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitializeDB(username, password, host string, port int, dbname string) {
	dataSourceName := username + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbname + "?parseTime=true"
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Check database connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connected!")
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func DB() *sql.DB {
	return db
}
