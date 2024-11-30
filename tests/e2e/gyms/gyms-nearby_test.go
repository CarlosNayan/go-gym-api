package gyms_e2e_test

import (
	"api-gym-on-go/models"
	"api-gym-on-go/tests/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGymsNearbyE2E(t *testing.T) {
	utils.ResetDb()
	app := utils.SetupTestApp("gyms")
	db := models.SetupDatabase("postgresql://root:admin@127.0.0.1:5432/public?sslmode=disable")
	gym := models.Gym{
		GymName:   "test gym",
		Latitude:  1.23456,
		Longitude: 1.23456,
	}
	db.Create(&gym)
	server := httptest.NewServer(utils.FiberToHttpHandler(app.Handler()))

	defer server.Close()

	t.Run("should be able to search gyms nearby", func(t *testing.T) {
		token := utils.CreateAndAuthenticateUser()

		req := httptest.NewRequest("GET", "/gyms/nearby?latitude=1.23456&longitude=1.23456", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, _ := app.Test(req)

		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("error reading response body: %v", err)
		}

		var responseData []map[string]interface{}
		err = json.Unmarshal(respBody, &responseData)
		if err != nil {
			t.Fatalf("Erro ao parsear JSON: %v", err)
		}
		fmt.Println(responseData)

		var found bool
		for _, item := range responseData {
			if gymName, ok := item["gym_name"].(string); ok && gymName == "test gym" {
				found = true
				break
			}
		}

		assert.True(t, found, "gym is not present in the response")
	})
}
