package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	DbConnectionString           string
	ApiPort                      string
	Environment                  string
	ApiMaxRequests               int
	ApiRequestsExpiration        time.Duration
	FiberStorageReset            bool
	SMTPHost                     string
	SMTPPort                     int
	SMTPUser                     string
	SMTPPassword                 string
	FiberStorageConnectionString string
}

var config Config

var configOnce sync.Once

func GetConfig() Config {
	configOnce.Do(func() {
		// ENV
		if k_service := os.Getenv("K_SERVICE"); k_service != "" {
			config.Environment = "production"
		} else {
			if environment, ok := os.LookupEnv("ENVIRONMENT"); !ok {
				config.Environment = "development"
			} else {
				config.Environment = environment
			}
		}

		// DB
		host := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		if host == "" || dbPort == "" || user == "" || password == "" || dbName == "" {
			panic("DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, and DB_NAME must be set")
		}
		config.DbConnectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbName)

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

		// Fiber Storage
		config.FiberStorageReset = config.Environment == "development"
		host = os.Getenv("FIBER_STORAGE_HOST")
		user = os.Getenv("FIBER_STORAGE_USER")
		password = os.Getenv("FIBER_STORAGE_PASSWORD")
		dbName = os.Getenv("FIBER_STORAGE_NAME")
		if host == "" || user == "" || password == "" || dbName == "" {
			panic("FIBER_STORAGE_HOST, FIBER_STORAGE_USER, FIBER_STORAGE_PASSWORD, and FIBER_STORAGE_NAME must be set")
		}
		config.FiberStorageConnectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

		// Email Engine
		if config.Environment == "production" {
			config.SMTPHost = os.Getenv("SMTP_HOST")
			config.SMTPPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
			config.SMTPUser = os.Getenv("SMTP_USER")
			config.SMTPPassword = os.Getenv("SMTP_PASSWORD")
		} else {
			config.SMTPHost = "localhost"
			config.SMTPPort = 1025
		}
	})
	return config
}
