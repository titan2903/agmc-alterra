package controllers

import (
	"day_4/transport"
)

func (c *controllers) HealthCheck() *transport.Response {
	return &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Server successfully running",
	}
}
