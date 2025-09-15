package utils

import (
	"api-gym-on-go/src/config/env"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"unicode/utf8"
)

func generateRandomIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}

func Encrypt(text string) (string, error) {
	iv, err := generateRandomIV()
	if err != nil {
		fmt.Println("Erro ao gerar IV:", err)
		return "", err
	}

	key, err := hex.DecodeString(env.CRYPTO_SECRET)
	if err != nil {
		fmt.Println("Erro ao converter chave para bytes:", err)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Erro ao inicializar o bloco de cifra:", err)
		return "", err
	}

	ciphertext := make([]byte, len(text))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, []byte(text))

	return fmt.Sprintf("%x:%x", iv, ciphertext), nil
}

func Decrypt(encryptedText string) (string, error) {
	var iv, ciphertext []byte
	_, err := fmt.Sscanf(encryptedText, "%x:%x", &iv, &ciphertext)
	if err != nil {
		return "", err
	}

	key, err := hex.DecodeString(env.CRYPTO_SECRET)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	if !utf8.ValidString(string(plaintext)) {
		return "", fmt.Errorf("invalid UTF-8 string")
	}

	return string(plaintext), nil
}

func Compare(encryptedText, textToCompare string) (bool, error) {
	decrypted, err := Decrypt(encryptedText)
	if err != nil {
		return false, err
	}
	return decrypted == textToCompare, nil
}
