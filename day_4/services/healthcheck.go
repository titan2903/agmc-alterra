package services

import (
	"day_4/transport"
	"net/http"
)

func (c *services) HealthCheck() *transport.Response {
	return &transport.Response{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "Server successfully running",
	}
}
