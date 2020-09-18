package main

import (
	"LeetCode-Rank/router"
	"LeetCode-Rank/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func main() {
	//utils.Update()
	c := cron.New(cron.WithSeconds())
	// utils.Update()
	_, err := c.AddFunc("0 9-59/10 2-23/1 * * *", utils.Update)
	if err != nil {
		log.Error(errors.WithStack(err))
		return
	}
	c.Start()
	Router := gin.Default()
	Router.GET("list", router.GetList)
	Router.GET("leaderboard", router.LeaderBoard)
	err = Router.Run(":4398")
	if err != nil {
		log.Error(errors.WithStack(err))
		return
	}
	//crawler.GetDifficulty(`two-sum`)
}
