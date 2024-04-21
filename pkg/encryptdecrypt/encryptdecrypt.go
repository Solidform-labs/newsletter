package encryptdecrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPassword(storedPassword string, passwordToCheck string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(passwordToCheck))
	return err == nil
}
