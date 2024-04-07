package apikeygenerator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateAPIKey(length int) (string, error) {
	fmt.Println("Generating an api key")

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var apiKeyBuilder string

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		apiKeyBuilder += string(charset[index.Int64()])
	}

	return apiKeyBuilder, nil
}
