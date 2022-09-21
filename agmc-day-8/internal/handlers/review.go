package handlers

// import (
// 	"agmc-day-8/internal/dto"
// 	"agmc-day-8/internal/models"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func (h *handler) CreateReview(c echo.Context) error {
// 	review := new(models.Review)
// 	response := new(dto.Response)
// 	c.Bind(review)

// 	if err := c.Validate(review); err != nil {
// 		return err
// 	}

// 	result, err := h.service.CreateReview(review)
// 	if err != nil {
// 		response.Code = 400
// 		response.Status = "failed"
// 		response.Message = "Failed to create review"
// 		return c.JSON(http.StatusBadRequest, response)
// 	}

// 	response.Code = result.Code
// 	response.Status = result.Status
// 	response.Message = result.Message
// 	response.Data = result.Data
// 	return c.JSON(http.StatusCreated, response)

// }

// func (h *handler) GetReviews(c echo.Context) error {
// 	response := new(dto.Response)

// 	result, err := h.service.GetReviews()
// 	if err != nil {
// 		response.Code = 404
// 		response.Status = "failed"
// 		response.Message = "Failed to get reviews"
// 		return c.JSON(http.StatusBadRequest, response)
// 	}

// 	response.Code = result.Code
// 	response.Status = result.Status
// 	response.Message = result.Message
// 	response.Data = result.Data
// 	return c.JSON(http.StatusCreated, response)
// }

// func (h *handler) UpdateReview(c echo.Context) error {
// 	review := new(models.Review)
// 	response := new(dto.Response)
// 	c.Bind(review)
// 	id := c.Param("id")
// 	if err := c.Validate(review); err != nil {
// 		return err
// 	}

// 	result, err := h.service.UpdateReview(review, id)
// 	if err != nil {
// 		response.Code = 400
// 		response.Status = "failed"
// 		response.Message = "Failed to update review"
// 		return c.JSON(http.StatusBadRequest, response)
// 	}

// 	response.Code = result.Code
// 	response.Status = result.Status
// 	response.Message = result.Message
// 	response.Data = result.Data
// 	return c.JSON(http.StatusCreated, response)
// }

// func (h *handler) DeleteReview(c echo.Context) error {
// 	response := new(dto.Response)
// 	id := c.Param("id")
// 	result, err := h.service.DeleteReview(id)
// 	if err != nil {
// 		response.Code = 400
// 		response.Status = "failed"
// 		response.Message = "Failed to delete review"
// 		return c.JSON(http.StatusBadRequest, response)
// 	}

// 	response.Code = result.Code
// 	response.Status = result.Status
// 	response.Message = result.Message
// 	response.Data = result.Data
// 	return c.JSON(http.StatusCreated, response)
// }

// func (h *handler) GetReview(c echo.Context) error {
// 	response := new(dto.Response)
// 	id := c.Param("id")
// 	result, err := h.service.GetReview(id)
// 	if err != nil {
// 		response.Code = 404
// 		response.Status = "failed"
// 		response.Message = "Failed to get review"
// 		return c.JSON(http.StatusBadRequest, response)
// 	}

// 	response.Code = result.Code
// 	response.Status = result.Status
// 	response.Message = result.Message
// 	response.Data = result.Data
// 	return c.JSON(http.StatusCreated, response)
// }
