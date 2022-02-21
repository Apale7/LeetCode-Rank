package main

import (
	config "LeetCode-Rank/config_loader"
	"LeetCode-Rank/db"
	"LeetCode-Rank/utils"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	config.Init()
	db.Init(ctx)
	// utils.Update()
	c := cron.New(cron.WithSeconds())
	// utils.Update()
	_, err := c.AddFunc("0 9-59/10 2-23/1 * * *", func() {
		utils.Update(ctx)
	})
	// go func() {
	// 	utils.Update(ctx)
	// }()
	if err != nil {
		log.Error(errors.WithStack(err))
		return
	}
	c.Start()
	r := gin.Default()
	CollectRouter(r)

	err = r.Run(":4398")
	// err = r.Run(":6799")
	if err != nil {
		log.Error(errors.WithStack(err))
		return
	}
	// crawler.GetDifficulty(`two-sum`)
}
