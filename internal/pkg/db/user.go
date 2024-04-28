package db

import (
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/models"
	"github.com/gofiber/fiber/v2/log"
)

func GetUserByEmail(email string, user *models.User) error {
	db := GetDB()
	err := db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&user.Password)
	if err != nil {
		log.Warn("error: ", err)
		return err
	}

	return nil
}
