package services

import (
	"agmc-day-7/internal/dto"
	"net/http"
)

func (s *services) HealthCheck() *dto.Response {
	return &dto.Response{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "Server successfully running",
	}
}
