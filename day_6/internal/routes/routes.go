package routes

import (
	"day_6/database/config"
	"day_6/internal/handlers"
	m "day_6/internal/middleware"
	repositories "day_6/internal/repositories"
	"day_6/internal/services"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)

	repository := repositories.NewRepositories(config.GetQuery(), config.ConnectDB())
	service := services.NewServices(repository)
	handler := handlers.NewHandlers(service)

	e.Validator = m.NewCustomValidator()

	gJwt := e.Group("/jwt")
	m.SetJwtMiddlewares(gJwt)

	//! Main Routes
	e.GET("/v1/healthcheck", handler.HealthCheck)

	e.POST("/v1/users", handler.CreateUser)
	e.POST("/v1/login", handler.UserLogin)

	e.GET("/v1/books/:id", handler.GetBookById)
	e.GET("/v1/books", handler.GetAllBooks)

	//! Authorization routes
	gJwt.PUT("/v1/users/:id", handler.UpdateUser)
	gJwt.GET("/v1/users/:id", handler.GetUserById)
	gJwt.GET("/v1/users", handler.GetAllUsers)
	gJwt.DELETE("/v1/users/:id", handler.DeleteUser)

	gJwt.PUT("/v1/books/:id", handler.UpdateBook)
	gJwt.DELETE("/v1/books/:id", handler.DeleteBook)
	gJwt.POST("/v1/books", handler.CreateBook)

	gJwt.POST("/v1/reviews", handler.CreateReview)
	gJwt.GET("/v1/reviews", handler.GetReviews)
	gJwt.PUT("/v1/reviews/:id", handler.UpdateReview)
	gJwt.DELETE("/v1/reviews/:id", handler.DeleteReview)

	return e
}
