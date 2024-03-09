package configs

import (
	"fmt"
	"os"
	"sync"
)

type Config struct {
	DbConnectionString string
	ApiPort            string
	Environment        string
}

var config Config

var configOnce sync.Once

func GetConfig() Config {
	configOnce.Do(func() {
		// DB
		host := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbname)
		// API
		apiPort, ok := os.LookupEnv("PORT")
		if !ok {
			apiPort = "8080"
		}
		// ENV
		environment, ok := os.LookupEnv("ENVIRONMENT")
		if !ok {
			environment = "development"
		}
		// Config
		config = Config{
			DbConnectionString: dbConnectionString,
			ApiPort:            apiPort,
			Environment:        environment,
		}
	})
	return config
}
