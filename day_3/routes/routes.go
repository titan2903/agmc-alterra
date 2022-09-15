package routes

import (
	"day_3/config"
	"day_3/controllers"
	"day_3/handlers"
	lib "day_3/lib/repositories"
	m "day_3/middleware"

	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)

	repository := lib.NewRepositories(config.GetQuery())
	controller := controllers.NewControllers(repository)
	handler := handlers.NewHandlers(controller)

	e.Validator = m.NewCustomValidator()

	e.GET("/v1/healthcheck", handler.HealthCheck)

	gJwt := e.Group("/jwt")
	m.SetJwtMiddlewares(gJwt)

	//! Main Routes
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

	return e
}
