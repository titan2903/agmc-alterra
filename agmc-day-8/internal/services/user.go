package services

import (
	"agmc-day-8/internal/dto"
	m "agmc-day-8/internal/middleware"
	"agmc-day-8/internal/models"
	"agmc-day-8/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (s *services) CreateUser(user *models.User) (*dto.Response, error) {

	hash, _ := utils.HashPassword(user.Password)
	userMapping := &models.User{
		Password: hash,
		Username: user.Username,
		Email:    user.Email,
	}

	err := s.repo.CreateUser(userMapping)
	if err != nil {
		return nil, err
	}

	result := &dto.Response{
		Code:    201,
		Status:  "success",
		Message: "Success to create user",
		Data:    err,
	}

	return result, err
}

func (s *services) UpdateUser(user *models.User, id int) (*dto.Response, error) {
	userMapping := &models.User{
		Username: user.Username,
		Email:    user.Email,
	}

	err := s.repo.UpdateUser(userMapping, id)
	if err != nil {
		return nil, err
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success to update user",
		Data:    err,
	}

	return result, nil
}

func (s *services) DeleteUser(id int) (*dto.Response, error) {

	err := s.repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success to delete user",
		Data:    err,
	}

	return result, nil
}

func (s *services) GetUserById(id int) (*dto.Response, error) {

	user, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userMapping := &dto.UserMapping{
		Username:  user.Username,
		Email:     user.Email,
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get user",
		Data:    userMapping,
	}
	return result, nil
}

func (s *services) GetAllUsers(keywords string) (*dto.Response, error) {
	users, err := s.repo.GetAllUsers(keywords)
	if err != nil {
		return nil, err
	}

	var arrUsers []dto.UserMapping
	for _, user := range users {
		userMapping := dto.UserMapping{
			Username:  user.Username,
			Email:     user.Email,
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		arrUsers = append(arrUsers, userMapping)
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get all users",
		Data:    arrUsers,
	}
	return result, nil
}

func (s *services) UserLogin(username, password string) (*dto.Response, error) {
	user, err := s.repo.UserLogin(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// If the two passwords don't match, return a 401 status.
		return nil, err
	}

	userReplicated := dto.LoginResponse{}
	userReplicated.Email = user.Email
	userReplicated.Username = user.Username
	userReplicated.ID = user.ID
	userReplicated.Token, err = m.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	result := &dto.Response{
		Code:    200,
		Status:  "success",
		Message: "Success Login",
		Data:    userReplicated,
	}
	return result, nil
}
