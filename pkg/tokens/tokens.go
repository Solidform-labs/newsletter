package tokens

import (
	"time"

	"github.com/Solidform-labs/newsletter/configs"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Second * 10).Unix(),
	})

	jwtKey := []byte(configs.GetConfig().JwtSecret)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
