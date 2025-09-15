package utils

import (
	"math/rand"
)

func GenerateUUID() string {
	// Fallback manual para gerar um UUID (baseado na RFC 4122)
	uuid := make([]byte, 36)
	template := "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"

	for i, c := range template {
		switch c {
		case 'x':
			uuid[i] = "0123456789abcdef"[rand.Intn(16)]
		case 'y':
			uuid[i] = "89ab"[rand.Intn(4)] // Define os valores válidos para 'y' conforme a RFC
		case '-':
			uuid[i] = '-'
		default:
			uuid[i] = byte(c)
		}
	}

	return string(uuid)
}
