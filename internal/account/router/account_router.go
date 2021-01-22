package router

import (
	"fmt"
	"github.com/BryceSun/beacon_academy/init_config"
	"github.com/BryceSun/beacon_academy/internal/account/handler"
	"github.com/BryceSun/beacon_academy/internal/account/model"
	. "github.com/BryceSun/beacon_academy/internal/common"
	"github.com/gin-gonic/gin"
)

func init() {
	r := init_config.RootRouter
	v1 := r.Group("/account")
	{
		v1.POST("/info", register)
		v1.POST("/token", login)
		v1.DELETE("/token", logout)
		v1.GET("/hello", hello)
	}
}

func register(ctx *gin.Context) {
	var u model.UserAccount
	ctx.Bind(&u)
	ctx.JSON(Output(handler.AddUserAccount(&u)))
}

func login(ctx *gin.Context) {
	var u model.UserAccount
	ctx.Bind(&u)
	ctx.JSON(Output(handler.GetUserToken(&u)))
}

func logout(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	id, _ := uid.(int64)
	ctx.JSON(Output(handler.DeleteUserToken(id)))
}

func hello(ctx *gin.Context) {
	fmt.Println("hello world")
	ctx.JSON(200, "hello world")
}
