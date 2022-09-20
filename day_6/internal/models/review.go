package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Description string             `json:"description" form:"description"`
	Rating      int64              `json:"rating" form:"rating" validate:"required"`
}
