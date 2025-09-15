package checkins_e2e_test

import (
	checkins_e2e_test_kit "api-gym-on-go/tests/modules/checkins/testkit"
	"api-gym-on-go/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckinsCountHistoryE2E(t *testing.T) {
	t.Run("should be able to count history", func(t *testing.T) {
		checkins_e2e_test_kit.SetupTest()
		opt := utils.HTTPTestOptions{
			Headers: &map[string]string{
				"Authorization": "Bearer " + checkins_e2e_test_kit.Token,
			},
		}
		resp := utils.RunHTTPTestRequest(t, "checkins", "GET", "/checkin/history/count", opt)

		assert.Equal(t, 0, int(resp.Obj["count"].(float64)), "count does not match")
		assert.Equalf(t, 200, resp.StatusCode, "get HTTP status 200")
	})

	t.Run("should be able to count history", func(t *testing.T) {
		checkins_e2e_test_kit.SetupTest("pre-create-checkin")
		opt := utils.HTTPTestOptions{
			Headers: &map[string]string{
				"Authorization": "Bearer " + checkins_e2e_test_kit.Token,
			},
		}
		resp := utils.RunHTTPTestRequest(t, "checkins", "GET", "/checkin/history/count", opt)

		assert.Equal(t, 1, int(resp.Obj["count"].(float64)), "count does not match")
		assert.Equalf(t, 200, resp.StatusCode, "get HTTP status 200")
	})
}
