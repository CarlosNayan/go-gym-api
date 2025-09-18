package checkins_e2e_test

import (
	checkins_e2e_test_kit "api-gym-on-go/tests/modules/checkins/testkit"
	"api-gym-on-go/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckinsValidateE2E(t *testing.T) {
	checkins_e2e_test_kit.SetupTest("create-and-authenticate-admin", "pre-create-checkin")

	t.Run("should be able to count history", func(t *testing.T) {
		opt := utils.HTTPTestOptions{
			Headers: &map[string]string{
				"Authorization": "Bearer " + checkins_e2e_test_kit.Token,
			},
		}
		resp := utils.RunHTTPTestRequest(t, "checkins", "PUT", "/checkin/validate/0ebd4f88-d712-4b0f-9278-41d595c690ad", opt)

		assert.Equalf(t, 200, resp.StatusCode, "put HTTP status 200")
	})
}
