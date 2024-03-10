package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	DbConnectionString    string
	ApiPort               string
	Environment           string
	ApiMaxRequests        int
	ApiRequestsExpiration time.Duration
	FiberStorageReset     bool
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
		if host == "" || dbPort == "" || user == "" || password == "" || dbname == "" {
			panic("DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, and DB_NAME must be set")
		}
		config.DbConnectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbname)
		// API
		if apiPort, ok := os.LookupEnv("PORT"); !ok {
			config.ApiPort = "8080"
		} else {
			config.ApiPort = apiPort
		}
		if apiMaxRequests, ok := os.LookupEnv("MAX_REQUESTS"); ok {
			var err error
			if config.ApiMaxRequests, err = strconv.Atoi(apiMaxRequests); err != nil {
				config.ApiMaxRequests = 20
			}
		} else {
			config.ApiMaxRequests = 20
		}
		if requestsExpiration, ok := os.LookupEnv("REQUESTS_EXPIRATION"); !ok {
			config.ApiRequestsExpiration = 30 * time.Second
		} else {
			if apiRequestsExpiration, err := time.ParseDuration(requestsExpiration); err != nil {
				config.ApiRequestsExpiration = 30 * time.Second
			} else {
				config.ApiRequestsExpiration = apiRequestsExpiration
			}
		}
		// ENV
		if environment, ok := os.LookupEnv("ENVIRONMENT"); !ok {
			config.Environment = "development"
		} else {
			config.Environment = environment
		}
		// Fiber Storage
		config.FiberStorageReset = config.Environment == "development"
	})
	return config
}
