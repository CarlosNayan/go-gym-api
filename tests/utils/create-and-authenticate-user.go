package utils

import (
	"api-gym-on-go/models"
	"api-gym-on-go/src/modules/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func CreateAndAuthenticateUser() string {

	app := fiber.New()
	db := models.SetupDatabase("postgresql://root:admin@127.0.0.1:5432/public?sslmode=disable")
	auth.Register(app, db)

	user := models.User{
		ID:           "1e2d4f88-d712-4b0f-9278-41d595c690ad",
		UserName:     "John Doe",
		Email:        "test@test.com",
		PasswordHash: "$2a$10$Dt3LAbYqOJiBPOW5VG/uXOL9Jk8DvqLBz16znHw5WLZiYZQCCED/.",
	}
	db.Create(&user)

	payload := map[string]interface{}{
		"email":    "test@test.com",
		"password": "123456",
	}

	body, err := json.Marshal(payload)
	if err != nil {
		panic(fmt.Sprintf("Erro ao serializar payload: %v", err))
	}

	req := httptest.NewRequest("POST", "/auth", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		panic(fmt.Sprintf("Erro ao testar a requisição: %v", err))
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Erro ao ler o corpo da resposta: %v", err))
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Falha na autenticação. Status: %d, Resposta: %s", resp.StatusCode, string(respBody)))
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(respBody, &responseData)
	if err != nil {
		panic(fmt.Sprintf("Erro ao desserializar a resposta: %v", err))
	}

	token, ok := responseData["token"].(string)
	if !ok {
		panic("Token não encontrado na resposta")
	}

	return token
}
