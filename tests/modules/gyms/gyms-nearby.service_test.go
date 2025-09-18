package gyms_e2e_test

import (
	gyms_e2e_test_kit "api-gym-on-go/tests/modules/gyms/testkit"
	"api-gym-on-go/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGymsNearbyE2E(t *testing.T) {
	t.Run("should be able to search gyms nearby", func(t *testing.T) {
		gyms_e2e_test_kit.SetupTest("pre-create-gym")
		opt := utils.HTTPTestOptions{
			Headers: &map[string]string{
				"Authorization": "Bearer " + gyms_e2e_test_kit.Token,
			},
		}
		resp := utils.RunHTTPTestRequest(t, "gyms", "GET", "/gyms/nearby?latitude=1.23456&longitude=1.23456", opt)

		var found bool
		for _, item := range resp.Arr {
			if gymName, ok := item["gym_name"].(string); ok && gymName == "test gym" {
				found = true
				break
			}
		}

		assert.True(t, found, "gym is not present in the response")
	})
}
