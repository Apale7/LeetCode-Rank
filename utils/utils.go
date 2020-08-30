package utils

import (
	"LeetCode-Rank/crawler"
	"LeetCode-Rank/db"
	"github.com/spf13/viper"
	"time"
)

func GetSolveNumberToday(username string) int{
	tmp := time.Now()
	t := time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, time.Local)
	count := 0
	db.Db.Model(&db.Accepted{}).Where("username = ? and time >= ?", username, t).Count(&count)
	return count
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
			db.AddProblem(submit.Question.QuestionFrontendID, submit.Question.TranslatedTitle)
			db.AddAccepted(submit.SubmitTime, username, submit.Question.QuestionFrontendID)
		}
	}
}
