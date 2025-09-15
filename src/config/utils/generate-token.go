package utils

import (
	"math/rand"
	"time"
)

func GenerateToken(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	token := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex := rng.Intn(len(charset))
		token[i] = charset[randomIndex]
	}

	return string(token)
}
