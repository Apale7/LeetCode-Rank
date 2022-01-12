package db

import "time"

type User struct {
	Meta
	Username       string
	Nickname       string
	CreatedAt      time.Time
	AcceptedReCord []*Accepted
}

type Accepted struct {
	Meta
	Medium int
	Easy   int
}

func NewUser() *User{
	return &User{
		Meta: Meta{
			ID:          0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Discription: "",
		},
	}		

}