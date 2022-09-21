package services

// import (
// 	"agmc-day-8/internal/dto"
// 	"agmc-day-8/internal/models"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func (s *services) CreateReview(review *models.Review) (*dto.Response, error) {
// 	reviewMapping := models.Review{
// 		Id:          primitive.NewObjectID(),
// 		Description: review.Description,
// 		Rating:      review.Rating,
// 	}

// 	err := s.repo.CreateReview(reviewMapping)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := &dto.Response{
// 		Code:    201,
// 		Status:  "success",
// 		Message: "Success create review",
// 		Data:    err,
// 	}

// 	return result, nil
// }

// func (s *services) GetReviews() (*dto.Response, error) {

// 	reviews, err := s.repo.GetReviews()
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := &dto.Response{
// 		Code:    200,
// 		Status:  "success",
// 		Message: "Success get reviews",
// 		Data:    reviews,
// 	}

// 	return result, nil
// }

// func (s *services) UpdateReview(review *models.Review, id string) (*dto.Response, error) {
// 	reviewMapping := bson.M{"description": review.Description, "rating": review.Rating}

// 	err := s.repo.UpdateReview(reviewMapping, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := &dto.Response{
// 		Code:    200,
// 		Status:  "success",
// 		Message: "Success update review",
// 		Data:    err,
// 	}

// 	return result, nil
// }

// func (s *services) DeleteReview(id string) (*dto.Response, error) {
// 	err := s.repo.DeleteReview(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := &dto.Response{
// 		Code:    200,
// 		Status:  "success",
// 		Message: "Success delete review",
// 		Data:    err,
// 	}

// 	return result, nil
// }

// func (s *services) GetReview(id string) (*dto.Response, error) {
// 	review, err := s.repo.GetReview(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := &dto.Response{
// 		Code:    200,
// 		Status:  "success",
// 		Message: "Success get review",
// 		Data:    review,
// 	}

// 	return result, nil
// }
