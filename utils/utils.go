package utils

import (
	"context"
	"time"

	"github.com/Apale7/LeetCode-Rank/biz/dal"
	"github.com/Apale7/LeetCode-Rank/biz/handler"
	"github.com/Apale7/LeetCode-Rank/db/model"
	"github.com/Apale7/LeetCode-Rank/service/crawler"

	"github.com/sirupsen/logrus"
)

func Update(ctx context.Context) {
	users, err := dal.GetUsers(ctx)
	if err != nil {
		logrus.Errorf("GetUsers error: %v", err)
		return
	}
	for _, user := range users {
		go func(user *model.User) {
			logrus.Printf("username: %s\n", user.Username)
			accepted := crawler.GetUserQuestionProgress(user.Username)
			if accepted == nil {
				logrus.Warnf("username: %s, accepted is nil", user.Username)
				return
			}
			accepted.UserID = user.ID
			err := dal.CreateAccepted(ctx, accepted)
			if err != nil {
				logrus.Errorf("username: %s, CreateAccepted error: %v", user.Username, err)
				return
			}
			logrus.Infof("%+v\n", accepted)
		}(user)
	}

	_, _ = handler.GetListFromCache(ctx, true) // flush cache
}

func GetDateBegin(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
