package handlers

import (
	m "day_3/middleware"
	"day_3/models"
	"day_3/transport"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) CreateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)

	if err := c.Validate(user); err != nil {
		return err
	}

	response := new(transport.Response)
	result, err := h.controller.CreateUser(user)

	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to create user"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Data = result.Data
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *handler) UpdateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)
	extractToken := m.ExtractTokenUserId(c)
	result, err := h.controller.UpdateUser(user, idInt)

	if err != nil || float64(idInt) != extractToken {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to update user"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Data = result.Data
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)

	extractToken := m.ExtractTokenUserId(c)

	result, err := h.controller.DeleteUser(idInt)

	if err != nil || float64(idInt) != extractToken {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to delete user"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Data = result.Data
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)
	result, err := h.controller.GetUserById(idInt)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "User not found"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Data = result.Data
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllUsers(c echo.Context) error {
	response := new(transport.Response)
	result, err := h.controller.GetAllUsers(c.QueryParam("keywords"))

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Users not found"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Data = result.Data
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) UserLogin(c echo.Context) error {
	response := new(transport.Response)
	user := new(models.User)
	c.Bind(user)
	result, err := h.controller.UserLogin(user.Username, user.Password)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Login failed"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Data = result.Data
	}
	return c.JSON(http.StatusOK, response)
}
