package users_e2e_test_kit

import (
	"api-gym-on-go/tests/utils"
	"strings"
)

var Token string

func SetupTest(params ...string) {
	utils.SetupEnviromentTest()
	paramsArr := strings.Join(params, ",")

	if strings.Contains(paramsArr, "create-and-authenticate-admin") {
		Token = utils.CreateAndAuthenticateUser("ADMIN")
	} else {

		Token = utils.CreateAndAuthenticateUser("MEMBER")
	}
}
