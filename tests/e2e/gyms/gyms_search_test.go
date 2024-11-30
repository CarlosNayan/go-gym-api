package gyms_e2e_test

import (
	"api-gym-on-go/tests/e2e/gyms/seed"
	"api-gym-on-go/tests/utils"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGymsSearchE2E(t *testing.T) {
	utils.ResetDb()
	token := utils.CreateAndAuthenticateUser("MEMBER")
	app := utils.SetupTestApp("gyms")
	seed.SeedGyms()
	server := httptest.NewServer(utils.FiberToHttpHandler(app.Handler()))

	defer server.Close()

	t.Run("should be able to search gyms nearby", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/gyms/search?query=test", nil)
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
