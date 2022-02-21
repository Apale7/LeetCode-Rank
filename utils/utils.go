package utils

import (
	"LeetCode-Rank/biz/dal"
	"LeetCode-Rank/service/crawler"
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func Update(ctx context.Context) {
	users, err := dal.GetUsers(ctx)
	if err != nil {
		logrus.Errorf("GetUsers error: %v", err)
		return
	}
	for _, user := range users {
		fmt.Printf("username: %s\n", user.Username)
		accepted := crawler.GetUserQuestionProgress(user.Username)
		accepted.UserID = user.ID
		err := dal.CreateAccepted(ctx, accepted)
		if err != nil {
			logrus.Errorf("username: %s, CreateAccepted error: %v", user.Username, err)
			continue
		}
		fmt.Printf("%+v\n", accepted)
		time.Sleep(time.Second * 5)
	}
}

func GetDateBegin(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
