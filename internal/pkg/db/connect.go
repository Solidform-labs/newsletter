package db

import (
	"database/sql"
	"fmt"

	"github.com/Solidform-labs/newsletter/configs"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {
	config := configs.GetConfig()
	db, err := sql.Open("postgres", config.DbConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}

func GetDB() *sql.DB {
	if db == nil {
		Connect()
	}
	return db
}
