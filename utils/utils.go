package utils

import (
	"LeetCode-Rank/crawler"
	"LeetCode-Rank/db"
	"github.com/spf13/viper"
	"time"
)

func GetScoreToday(username string) (int, int, int) {
	tmp := time.Now()
	t := time.Now()
	if tmp.Hour() >= 8 {
		t = time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, time.Local)
	} else {
		t = time.Date(tmp.Year(), tmp.Month(), tmp.Day()-1, 0, 0, 0, 0, time.Local)
	}

	easy, medium, hard := 0,0,0
	table := db.Db.Table("accepteds").Where("username = ?", username).Joins(`left join problems on accepteds.problem_id = problems.id`)
	table.Where("time >= ? and difficulty = ?", t, 0).Count(&easy)
	table.Where("time >= ? and difficulty = ?", t, 1).Count(&medium)
	table.Where("time >= ? and difficulty = ?", t, 2).Count(&hard)
	return easy, medium, hard
}

func Update() {
	viper.AddConfigPath("config")
	viper.SetConfigName("userlist")
	viper.ReadInConfig()
	var Users []string = viper.GetStringSlice("users")
	//fmt.Println(Users)
	for _, username := range Users {
		submits := crawler.GetData(username)
		for _, submit := range submits {
			level := crawler.GetDifficulty(submit.Question.TitleSlug)
			db.AddProblem(submit.Question.QuestionFrontendID, submit.Question.TranslatedTitle, level)
			db.AddAccepted(submit.SubmitTime, username, submit.Question.QuestionFrontendID)
		}
	}
}

