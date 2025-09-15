package users_e2e_test

import (
	users_e2e_test_kit "api-gym-on-go/tests/modules/users/testkit"
	"api-gym-on-go/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegisterE2E(t *testing.T) {
	t.Run("should be able to register", func(t *testing.T) {
		users_e2e_test_kit.SetupTest()
		opt := utils.HTTPTestOptions{
			Headers: &map[string]string{
				"Authorization": "Bearer " + users_e2e_test_kit.Token,
			},
			Body: &map[string]interface{}{
				"user_name": "Jhon Doe",
				"email":     "user@email.com",
				"password":  "123456",
				"role":      "MEMBER",
			},
		}
		resp := utils.RunHTTPTestRequest(t, "users", "POST", "/users/create", opt)

		assert.Equalf(t, 201, resp.StatusCode, "get HTTP status 201")
	})
}
