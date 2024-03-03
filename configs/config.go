package configs

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	DbConnectionString string
	ApiPort            string
}

var config Config

var configOnce sync.Once

func GetConfig() Config {
	configOnce.Do(func() {
		// DB
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbname := os.Getenv("DB_NAME")
		dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, dbPort, dbname)
		// API
		apiPort, ok := os.LookupEnv("API_PORT")
		if !ok {
			apiPort = "3000"
		}
		// Config
		config = Config{
			DbConnectionString: dbConnectionString,
			ApiPort:            apiPort,
		}
	})
	return config
}
