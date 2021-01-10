package handler

import (
	"LeetCode-Rank/db"
	"LeetCode-Rank/model"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func GetList(c *gin.Context) {
	viper.AddConfigPath("config")
	viper.SetConfigName("userlist")
	viper.ReadInConfig()
	var Users []string = viper.GetStringSlice("users")
	viper.SetConfigName("user_map")
	viper.ReadInConfig()
	user_map := viper.GetStringMap("usermap")
	data := []model.Rank{}
	tmp := model.Rank{}
	for _, username := range Users {
		tmp.Name = user_map[username].(string)
		ac7Day, err := db.RedisClient.Get(username+"_ac_7day").Int()
		easy, err := db.RedisClient.Get(username+"_ac_today_easy").Int()
		medium, err := db.RedisClient.Get(username+"_ac_today_medium").Int()
		hard, err := db.RedisClient.Get(username+"_ac_today_hard").Int()

		acNum, err := db.RedisClient.Get(username+"_ac_total").Int()
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