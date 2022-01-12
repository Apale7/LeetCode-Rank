package db

import "time"

type Meta struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Discription string
}

func NewMeta() *Meta {
	return &Meta{
		ID:          0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Discription: "",
	}
}
