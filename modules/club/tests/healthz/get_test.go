package healthz

import (
	responses "app/modules/club/domain/dto/responses/healthz"
	"app/modules/club/tests/setup"
	"app/modules/club/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealthz(t *testing.T) {
	utils.It(t, "should return health check with database information", func(t *testing.T) {
		// arrange
		baseSetup := setup.GetSetupEnvironment(t)

		// act
		sendResponse, err := setup.SendRequest(baseSetup.Echo, http.MethodGet, "/v1/club/healthz", nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, sendResponse.Code)

		// assert
		var res responses.GetHealthzResponseDto
		err = utils.Deserialize(sendResponse.Body.Bytes(), &res)

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Database)
		assert.Equal(t, "16.0", res.Database.Version)
		assert.True(t, res.Database.OpenConnections > 0)
	})
}
