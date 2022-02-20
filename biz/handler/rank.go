package handler

import (
	"LeetCode-Rank/biz/dal"
	"LeetCode-Rank/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetList(c *gin.Context) {
	users, err := dal.GetUsers(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	for _, user := range users {
		tmp.Name = user_map[username].(string)
		ac7Day, err := db.RedisClient.Get(username + "_ac_7day").Int()
		easy, err := db.RedisClient.Get(username + "_ac_today_easy").Int()
		medium, err := db.RedisClient.Get(username + "_ac_today_medium").Int()
		hard, err := db.RedisClient.Get(username + "_ac_today_hard").Int()

		acNum, err := db.RedisClient.Get(username + "_ac_total").Int()
		tmp.TotalAC = acNum
		tmp.Easy = easy
		tmp.Medium = medium
		tmp.Hard = hard
		tmp.TotalAC7Day = ac7Day
		if err != nil {
			log.Warnf("获取redis数据err: ", err)
		}
		fmt.Printf("%+v\n", tmp)
		data = append(data, tmp)
		// fmt.Println(tmp)
	}
	c.JSON(http.StatusOK, data)
}
