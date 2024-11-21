package utils

import (
	"crypto/rand"
	"fmt"
)

// GenerateUUIDv4 gera um UUIDv4 com Go nativo
func GenerateUUIDv4() (string, error) {
	// Gerar 16 bytes aleatórios
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Definir os bits de versão (UUIDv4) nos 6 primeiros bits do 7º byte
	bytes[6] = (bytes[6] & 0x0f) | 0x40 // versão 4
	// Definir os bits de variante nos 2 primeiros bits do 9º byte
	bytes[8] = (bytes[8] & 0x3f) | 0x80 // variante 10xx

	// Formatar os bytes como uma string UUID
	return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:]), nil
}

func main() {
	uuid, err := GenerateUUIDv4()
	if err != nil {
		fmt.Println("Erro ao gerar UUID:", err)
		return
	}

	fmt.Println("UUID gerado:", uuid)
}
