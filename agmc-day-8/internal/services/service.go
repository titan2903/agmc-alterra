package services

import (
	"agmc-day-8/internal/dto"
	"agmc-day-8/internal/models"
	repositories "agmc-day-8/internal/repositories"
)

type services struct {
	repo repositories.Repositories
}

type Services interface {
	// ! Book Services
	CreateBook(book *models.Book) (*dto.Response, error)
	UpdateBook(book *models.Book, id int) (*dto.Response, error)
	DeleteBook(id int) (*dto.Response, error)
	GetBookById(id int) (*dto.Response, error)
	GetAllBooks(keywords string) (*dto.Response, error)

	// ! User Services
	CreateUser(user *models.User) (*dto.Response, error)
	UpdateUser(user *models.User, id int) (*dto.Response, error)
	DeleteUser(id int) (*dto.Response, error)
	GetUserById(id int) (*dto.Response, error)
	GetAllUsers(keywords string) (*dto.Response, error)
	UserLogin(username, password string) (*dto.Response, error)

	// ! Health Check Services
	HealthCheck() *dto.Response

	// ! Review Services
	// CreateReview(review *models.Review) (*dto.Response, error)
	// GetReviews() (*dto.Response, error)
	// DeleteReview(id string) (*dto.Response, error)
	// UpdateReview(review *models.Review, id string) (*dto.Response, error)
	// GetReview(id string) (*dto.Response, error)
}

func NewServices(r repositories.Repositories) Services {
	return &services{repo: r}
}
