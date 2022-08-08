package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	Author      string             `json:"author" validate:"required"`
	Title       string             `json:"title" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Created     time.Time          `json:"created_at"`
	Updated     time.Time          `json:"updated_at"`
}
