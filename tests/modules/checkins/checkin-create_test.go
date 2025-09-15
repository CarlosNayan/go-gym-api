package checkins_e2e_test

import (
	checkins_e2e_test_kit "api-gym-on-go/tests/modules/checkins/testkit"
	"api-gym-on-go/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckinsCreateE2E(t *testing.T) {
	t.Run("should be able to count history", func(t *testing.T) {
		checkins_e2e_test_kit.SetupTest()
		opt := utils.HTTPTestOptions{
			Headers: &map[string]string{
				"Authorization": "Bearer " + checkins_e2e_test_kit.Token,
			},
			Body: &map[string]interface{}{
				"id_gym":         "2e2d4f88-d712-4b0f-9278-41d595c690ad",
				"user_latitude":  1.23456,
				"user_longitude": 1.23456,
			},
		}

		resp := utils.RunHTTPTestRequest(t, "checkins", "POST", "/checkin/create", opt)

		assert.Equalf(t, 201, resp.StatusCode, "get HTTP status 201")
	})
}
