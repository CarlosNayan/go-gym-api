package gyms_e2e_test_kit

import (
	"api-gym-on-go/tests/utils"
	"strings"
)

var Token string

func SetupTest(params ...string) {
	utils.SetupEnviromentTest()
	Token = utils.CreateAndAuthenticateUser("ADMIN")

	paramsArr := strings.Join(params, ",")

	if strings.Contains(paramsArr, "pre-create-gym") {
		SeedGyms()
	}
}
