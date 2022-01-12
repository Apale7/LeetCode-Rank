package utils

import (
	"LeetCode-Rank/db"
	"LeetCode-Rank/model"
	"LeetCode-Rank/service/crawler"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func getScoreToday(username string) (int64, int64, int64) {
	tmp := time.Now()
	t := time.Now()
	if tmp.Hour() >= 8 {
		t = time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, time.Local)
	} else {
		t = time.Date(tmp.Year(), tmp.Month(), tmp.Day()-1, 0, 0, 0, 0, time.Local)
	}

	var easy, medium, hard int64
	table := db.Db.Table("accepteds").Where("username = ?", username).Joins(`left join problems on accepteds.problem_id = problems.id`)
	table.Where("time >= ? and difficulty = ?", t, 0).Count(&easy)
	table.Where("time >= ? and difficulty = ?", t, 1).Count(&medium)
	table.Where("time >= ? and difficulty = ?", t, 2).Count(&hard)
	return easy, medium, hard
}

func getNum7Day(username string) int64 {
	tmp := time.Now()
	lasts7Day := time.Date(tmp.Year(), tmp.Month(), tmp.Day()-6, 0, 0, 0, 0, time.Local)
	table := db.Db.Table("accepteds").Where("username = ?", username).Joins(`left join problems on accepteds.problem_id = problems.id`)
	var num int64
	table.Where("time >= ?", lasts7Day).Count(&num)
	fmt.Println("num: ", num)
	return num
}

func Update() {
	fmt.Println("???")
	viper.AddConfigPath("config")
	viper.SetConfigName("userlist")
	_ = viper.ReadInConfig()
	var Users []string = viper.GetStringSlice("users")
	fmt.Println(Users)
	for _, username := range Users {
		fmt.Printf("username: %s\n", username)
		submits := crawler.GetData(username)
		for _, submit := range submits {
			level := crawler.GetDifficulty(submit.Question.TitleSlug)
			db.AddProblem(submit.Question.QuestionFrontendID, submit.Question.TranslatedTitle, level)
			db.AddAccepted(submit.SubmitTime, username, submit.Question.QuestionFrontendID)
		}
		acInfo := crawler.GetUserAcInfo(username)
		fmt.Printf("%+v\n", acInfo)
		writeRedis(username, acInfo)
		time.Sleep(time.Second * 5)
	}
}
