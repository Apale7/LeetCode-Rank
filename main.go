package main

import (
	"LeetCode-Rank/router"
	"LeetCode-Rank/utils"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	// utils.Update()
	c.AddFunc("0 9-59/10 2-23/1 * * *", utils.Update)
	c.Start()
	Router := gin.Default()
	Router.GET("list", router.GetList)
	Router.GET("leaderboard", router.LeaderBoard)
	Router.Run(":4398")
}
