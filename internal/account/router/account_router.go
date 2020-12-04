package router

import (
	"github.com/BryceSun/beacon_academy/init_config"
	"github.com/gin-gonic/gin"
)

func init() {
	r := init_config.RootRouter
	v1 := r.Group("/account")
	{
		v1.GET("/add", createUserAccount)
	}
}


func createUserAccount(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"hello": "world",
	})
}