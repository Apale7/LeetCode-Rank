package utils

import (
	"LeetCode-Rank/crawler"
	"LeetCode-Rank/db"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func getScoreToday(username string) (int, int, int) {
	tmp := time.Now()
	t := time.Now()
	if tmp.Hour() >= 8 {
		t = time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, time.Local)
	} else {
		t = time.Date(tmp.Year(), tmp.Month(), tmp.Day()-1, 0, 0, 0, 0, time.Local)
	}

	easy, medium, hard := 0, 0, 0
	table := db.Db.Table("accepteds").Where("username = ?", username).Joins(`left join problems on accepteds.problem_id = problems.id`)
	table.Where("time >= ? and difficulty = ?", t, 0).Count(&easy)
	table.Where("time >= ? and difficulty = ?", t, 1).Count(&medium)
	table.Where("time >= ? and difficulty = ?", t, 2).Count(&hard)
	return easy, medium, hard
}

func getNum7Day(username string) int {
	tmp := time.Now()
	lasts7Day := time.Date(tmp.Year(), tmp.Month(), tmp.Day()-6, 0, 0, 0, 0, time.Local)
	table := db.Db.Table("accepteds").Where("username = ?", username).Joins(`left join problems on accepteds.problem_id = problems.id`)
	num := 0
	table.Where("time >= ?", lasts7Day).Count(&num)
	fmt.Println("num: ", num)
	return num
}

func getAll(username string) int {
	num := 0
	db.Db.Table("accepteds").Where("username = ?", username).Count(&num)
	fmt.Println("num: ", num)
	return num
}

func Update() {
	viper.AddConfigPath("config")
	viper.SetConfigName("userlist")
	_ = viper.ReadInConfig()
	var Users []string = viper.GetStringSlice("users")
	//fmt.Println(Users)
	for _, username := range Users {
		submits := crawler.GetData(username)
		for _, submit := range submits {
			level := crawler.GetDifficulty(submit.Question.TitleSlug)
			db.AddProblem(submit.Question.QuestionFrontendID, submit.Question.TranslatedTitle, level)
			db.AddAccepted(submit.SubmitTime, username, submit.Question.QuestionFrontendID)
		}
		writeRedis(username)
	}
}

func writeRedis(username string) {
	ac7Day := getNum7Day(username)
	easy, medium, hard := getScoreToday(username)
	allAc := getAll(username)
	db.RedisClient.Set(username+"_ac_total", allAc, 0)
	db.RedisClient.Set(username+"_ac_today_easy", easy, 0)
	db.RedisClient.Set(username+"_ac_today_medium", medium, 0)
	db.RedisClient.Set(username+"_ac_today_hard", hard, 0)
	db.RedisClient.Set(username+"_ac_7day", ac7Day, 0)
}
