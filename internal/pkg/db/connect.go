package db

import (
	"database/sql"
	"sync"

	"github.com/Solidform-labs/newsletter/configs"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
)

var db *sql.DB

var dbOnce sync.Once

func Connect() {
	dbOnce.Do(func() {
		config := configs.GetConfig()
		var err error
		db, err = sql.Open("postgres", config.DbConnectionString)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		log.Info("Connected to database")
	})
}

func GetDB() *sql.DB {
	if db == nil {
		Connect()
	}
	return db
}
