package gyms_e2e_test

import (
	gyms_e2e_test_kit "api-gym-on-go/tests/modules/gyms/testkit"
	"api-gym-on-go/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGymCreateE2E(t *testing.T) {
	t.Run("should be able to register a gym", func(t *testing.T) {
		gyms_e2e_test_kit.SetupTest()
		opt := utils.HTTPTestOptions{
			Headers: &map[string]string{
				"Authorization": "Bearer " + gyms_e2e_test_kit.Token,
			},
			Body: &map[string]interface{}{
				"gym_name":  "test gym",
				"latitude":  1.23456,
				"longitude": 1.23456,
			},
		}
		resp := utils.RunHTTPTestRequest(t, "gyms", "POST", "/gyms/create", opt)

		assert.Equalf(t, 201, resp.StatusCode, "get HTTP status 201")
	})
}
