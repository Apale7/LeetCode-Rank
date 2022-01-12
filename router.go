package main

import (
	"LeetCode-Rank/biz/handler"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) {
	r.GET("list", handler.GetList)
}
