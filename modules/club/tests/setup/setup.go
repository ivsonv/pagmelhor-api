package setup

import (
	"app/configs"
	"app/modules/club"
	"bytes"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

type TestEnvironment struct {
	Echo *echo.Echo
}

func GetSetupEnvironment(t *testing.T) *TestEnvironment {
	// Setup test configuration
	cfg := &configs.Config{
		DBDriver:   "postgres",
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "tests",
		DBPassword: "tests",
		DBName:     "postgres",
		DBSSLMode:  "disable",
	}
	os.Setenv("DB_NAME", cfg.DBName)

	e := echo.New()

	// Start the club module with test configuration
	club.Start(e.Group("v1/club"), cfg)

	return &TestEnvironment{
		Echo: e,
	}
}

func SendRequest(e *echo.Echo, method, path string, body []byte) (*httptest.ResponseRecorder, error) {
	// arrange
	req := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	// act
	e.ServeHTTP(res, req)

	return res, nil
}
