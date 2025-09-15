package users_e2e_test

import (
	users_e2e_test_kit "api-gym-on-go/tests/modules/users/testkit"
	"api-gym-on-go/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserProfileE2E(t *testing.T) {
	t.Run("should be able to get profile", func(t *testing.T) {
		users_e2e_test_kit.SetupTest()
		opt := utils.HTTPTestOptions{
			Headers: &map[string]string{
				"Authorization": "Bearer " + users_e2e_test_kit.Token,
			},
		}
		resp := utils.RunHTTPTestRequest(t, "users", "GET", "/users/me", opt)

		userName, ok := resp.Obj["user_name"].(string)
		assert.True(t, ok, "user_name is not present in the response")
		assert.Contains(t, userName, "John Doe", "user_name does not match")
	})
}
