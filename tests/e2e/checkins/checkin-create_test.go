package checkins_e2e_test

import (
	"api-gym-on-go/tests/e2e/checkins/seed"
	"api-gym-on-go/tests/utils"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckinsCreateE2E(t *testing.T) {
	preCreateCheckin := false
	utils.ResetDb()
	token := utils.CreateAndAuthenticateUser()
	app := utils.SetupTestApp("checkins")
	seed.SeedCheckins(preCreateCheckin)
	server := httptest.NewServer(utils.FiberToHttpHandler(app.Handler()))

	defer server.Close()

	t.Run("should be able to count history", func(t *testing.T) {

		payload := map[string]interface{}{
			"id_gym":         "2e2d4f88-d712-4b0f-9278-41d595c690ad",
			"user_latitude":  1.23456,
			"user_longitude": 1.23456,
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("falha ao codificar payload: %v", err)
		}

		req := httptest.NewRequest("POST", "/checkin/create", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, _ := app.Test(req)

		assert.Equalf(t, 201, resp.StatusCode, "get HTTP status 201")
	})
}
