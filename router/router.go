package router

import (
	"LeetCode-Rank/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"html/template"
	"io/ioutil"
	"net/http"
	"fmt"
)

type Rank struct {
	Name string `json:"name"`
	Easy int `json:"easy"`
	Medium int `json:"medium"`
	Hard int `json:"hard"`
}

func GetList(c *gin.Context) {
	viper.AddConfigPath("config")
	viper.SetConfigName("userlist")
	viper.ReadInConfig()
	var Users []string = viper.GetStringSlice("users")
	viper.SetConfigName("user_map")
	viper.ReadInConfig()
	user_map := viper.GetStringMap("usermap")
	data := []Rank{}
	tmp := Rank{}
	for _, username := range Users {
		tmp.Name = user_map[username].(string)
		tmp.Easy, tmp.Medium, tmp.Hard = utils.GetScoreToday(username)
		data = append(data, tmp)
		// fmt.Println(tmp)
	}
	c.JSON(http.StatusOK, data)
}

func LeaderBoard(c *gin.Context)  {
	viper.AddConfigPath("config")
	viper.SetConfigName("userlist")
	viper.ReadInConfig()
	var Users []string = viper.GetStringSlice("users")
	viper.SetConfigName("user_map")
	viper.ReadInConfig()
	user_map := viper.GetStringMap("usermap")
	// fmt.Println(user_map)
	viper.ReadInConfig()
	data := make(map[string]int)
	for _, username := range Users {
		fmt.Println(username)
		data[user_map[username].(string)] = 0
	}
	bytes, err := ioutil.ReadFile("./static/html/leader_board.html")
	if err != nil {
		log.Warningln("读取模板文件失败")
		log.Warningln(errors.WithStack(err).Error())
		c.Writer.WriteHeader(500)
		return
	}
	tplStr := string(bytes)
	mp := make(map[string]map[string]int)
	mp["Users"] = data
	tpl := template.Must(template.New("leader_board").Parse(tplStr))
	err = tpl.Execute(c.Writer, mp)
	if err != nil {
		log.Warningln("渲染模板失败")
		log.Warningln(errors.WithStack(err))
		c.Writer.WriteHeader(500)
	}
}