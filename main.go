package main

import (
	"LeetCode-Rank/router"
	"LeetCode-Rank/utils"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	//utils.Update()
	//utils.GetSolveNumberToday("a_haw-2")
	c := cron.New()
	utils.Update()
	c.AddFunc("0 30 * * * *", utils.Update)
	Router := gin.Default()
	Router.GET("list", router.GetList)
	Router.GET("leaderboard", router.LeaderBoard)
	Router.Run(":4396")
}
