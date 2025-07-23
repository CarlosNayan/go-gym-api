package users_e2e_test

import (
	"api-gym-on-go/tests/utils"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegisterE2E(t *testing.T) {
	utils.ResetDb()
	app := utils.SetupTestApp("users")
	server := httptest.NewServer(utils.FiberToHttpHandler(app.Handler()))
	defer server.Close()

	t.Run("should be able to register", func(t *testing.T) {
		payload := map[string]interface{}{
			"user_name": "Jhon Doe",
			"email":     "user@email.com",
			"password":  "123456",
			"role":      "MEMBER",
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("falha ao codificar payload: %v", err)
		}

		req := httptest.NewRequest("POST", "/users/create", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req, -1)

		assert.Equalf(t, 201, resp.StatusCode, "get HTTP status 201")
	})
}
