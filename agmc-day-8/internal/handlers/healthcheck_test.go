package handlers

import (
	"agmc-day-8/database/config"
	"agmc-day-8/internal/mock"
	"agmc-day-8/internal/repositories"
	"agmc-day-8/internal/services"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	var (
		echoMock = mock.EchoMock{E: echo.New()}
		// repository = repositories.NewRepositories(config.GetQuery(), config.ConnectDB())
		repository = repositories.NewRepositories(config.GetQuery())
		service    = services.NewServices(repository)
		h          = NewHandlers(service)
	)

	t.Run("success", func(t *testing.T) {

		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/v1/healthcheck")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.HealthCheck(c)) {
			asserts.Equal(200, rec.Code)
		}
	})
}
