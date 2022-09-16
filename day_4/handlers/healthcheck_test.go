package handlers

import (
	"day_4/mock"
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	echoMock = mock.EchoMock{E: echo.New()}
	h        = new(handler)
)

func TestHealthCheckSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	fmt.Println("============== rec", rec)
	fmt.Println("============== rec", rec.Code)
	c.SetPath("/v1/healthcheck")

	asserts := assert.New(t)

	// testing
	if asserts.NoError(h.HealthCheck(c)) {
		fmt.Println("====================masuk")
		asserts.Equal(200, rec.Code)

		// body := rec.Body.String()
		// fmt.Println(body)
		// asserts.Contains(body, "code")
		// asserts.Contains(body, "status")
		// asserts.Contains(body, "message")
		// asserts.Contains(body, "data")
	}
}
