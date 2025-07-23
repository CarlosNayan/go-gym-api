package checkins_e2e_test

import (
	"api-gym-on-go/tests/modules/checkins/seed"
	"api-gym-on-go/tests/utils"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckinsValidateE2E(t *testing.T) {
	preCreateCheckin := true
	utils.ResetDb()
	token := utils.CreateAndAuthenticateUser("ADMIN")
	app := utils.SetupTestApp("checkins")
	seed.SeedCheckins(preCreateCheckin)
	server := httptest.NewServer(utils.FiberToHttpHandler(app.Handler()))

	defer server.Close()

	t.Run("should be able to count history", func(t *testing.T) {

		req := httptest.NewRequest("PUT", "/checkin/validate/0ebd4f88-d712-4b0f-9278-41d595c690ad", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, _ := app.Test(req)

		assert.Equalf(t, 200, resp.StatusCode, "put HTTP status 200")
	})
}
