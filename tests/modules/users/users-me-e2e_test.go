package users_e2e_test

import (
	"api-gym-on-go/tests/utils"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserProfileE2E(t *testing.T) {
	utils.ResetDb()
	app := utils.SetupTestApp("users")
	server := httptest.NewServer(utils.FiberToHttpHandler(app.Handler()))
	defer server.Close()

	t.Run("should be able to get profile", func(t *testing.T) {
		token := utils.CreateAndAuthenticateUser("MEMBER")

		req := httptest.NewRequest("GET", "/users/me", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, _ := app.Test(req, -1)

		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("error reading response body: %v", err)
		}

		var responseData map[string]interface{}
		err = json.Unmarshal(respBody, &responseData)
		if err != nil {
			t.Fatalf("Erro ao parsear JSON: %v", err)
		}

		userName, ok := responseData["user_name"].(string)
		assert.True(t, ok, "user_name is not present in the response")
		assert.Contains(t, userName, "John Doe", "user_name does not match")
	})
}
