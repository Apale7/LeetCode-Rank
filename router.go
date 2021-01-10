package main

import (
	"LeetCode-Rank/handler"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine){
	r.GET("list", handler.GetList)
}
