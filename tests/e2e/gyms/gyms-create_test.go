package gyms_e2e_test

import (
	"api-gym-on-go/tests/utils"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGymCreateE2E(t *testing.T) {
	utils.ResetDb()
	token := utils.CreateAndAuthenticateUser("MEMBER")
	app := utils.SetupTestApp("gyms")
	server := httptest.NewServer(utils.FiberToHttpHandler(app.Handler()))
	defer server.Close()

	t.Run("should be able to register a gym", func(t *testing.T) {
		payload := map[string]interface{}{
			"gym_name":  "test gym",
			"latitude":  1.23456,
			"longitude": 1.23456,
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("falha ao codificar payload: %v", err)
		}

		req := httptest.NewRequest("POST", "/gyms/create", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, _ := app.Test(req, -1)

		assert.Equalf(t, 201, resp.StatusCode, "get HTTP status 201")
	})
}
