package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meta struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Discription string             `bson:"discription"`
}

func NewMeta() *Meta {
	return &Meta{
		ID:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Discription: "",
	}
}
