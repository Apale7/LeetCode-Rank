package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Meta     `bson:"inline"`
	Username string `bson:"username"`
	Nickname string `bson:"nickname"`
}

func WithUsername(username string) func(*User) {
	return func(user *User) {
		user.Username = username
	}
}

func WithNickname(nickname string) func(*User) {
	return func(user *User) {
		user.Nickname = nickname
	}
}

func NewUser() *User {
	return &User{
		Meta: Meta{
			ID:          primitive.NilObjectID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Discription: "",
		},
	}
}
