package utils

import (
	"LeetCode-Rank/biz/dal"
	"LeetCode-Rank/service/crawler"
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// func getScoreToday(username string) (int64, int64, int64) {
// 	tmp := time.Now()
// 	t := time.Now()
// 	if tmp.Hour() >= 8 {
// 		t = time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, time.Local)
// 	} else {
// 		t = time.Date(tmp.Year(), tmp.Month(), tmp.Day()-1, 0, 0, 0, 0, time.Local)
// 	}

// 	var easy, medium, hard int64
// 	table := db.Db.Table("accepteds").Where("username = ?", username).Joins(`left join problems on accepteds.problem_id = problems.id`)
// 	table.Where("time >= ? and difficulty = ?", t, 0).Count(&easy)
// 	table.Where("time >= ? and difficulty = ?", t, 1).Count(&medium)
// 	table.Where("time >= ? and difficulty = ?", t, 2).Count(&hard)
// 	return easy, medium, hard
// }

// func getNum7Day(username string) int64 {
// 	tmp := time.Now()
// 	lasts7Day := time.Date(tmp.Year(), tmp.Month(), tmp.Day()-6, 0, 0, 0, 0, time.Local)
// 	table := db.Db.Table("accepteds").Where("username = ?", username).Joins(`left join problems on accepteds.problem_id = problems.id`)
// 	var num int64
// 	table.Where("time >= ?", lasts7Day).Count(&num)
// 	fmt.Println("num: ", num)
// 	return num
// }

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
