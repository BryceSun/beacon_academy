package router

import (
	"fmt"
	"github.com/BryceSun/beacon_academy/init_config"
	"github.com/BryceSun/beacon_academy/internal/account/handler"
	"github.com/BryceSun/beacon_academy/internal/account/model"
	"github.com/gin-gonic/gin"
)

func init() {
	r := init_config.RootRouter
	v1 := r.Group("/account")
	{
		v1.POST("/info", register)
		v1.POST("/token", login)
		v1.GET("/hello", hello)
	}
}

func register(ctx *gin.Context) {
	var u model.UserAccount
	ctx.Bind(&u)
	handler.AddUserAccount(&u)
	ctx.JSON(200, u)
}

func login(ctx *gin.Context) {
	var u model.UserAccount
	ctx.Bind(&u)
	token, err := handler.GetUserToken(&u)
	if err != nil {
		ctx.JSON(304, err)
	}
	ctx.Header("authentication", token)
	ctx.JSON(200, token)
}

func hello(ctx *gin.Context) {

	fmt.Println("hello world")
	ctx.JSON(200, "hello world")
}
