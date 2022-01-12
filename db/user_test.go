package db

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := NewUser()
	user.Username = "apale7"
	user.Nickname = "apale"
	user.AcceptedReCord = []*Accepted{
		&Accepted{
			Meta:   *NewMeta(),
			Medium: 1,
		},
	}
	res, err := client.Database("leetcode").Collection("user").InsertOne(context.TODO(), user)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("res: %v\n", res)
}
