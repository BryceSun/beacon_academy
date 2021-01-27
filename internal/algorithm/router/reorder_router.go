package router

import (
	"encoding/json"
	"github.com/BryceSun/beacon_academy/init_config"
	"github.com/BryceSun/beacon_academy/internal/algorithm/handler"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	r := init_config.RootRouter
	v1 := r.Group("/algorithm")
	//花括号可不要，但为了可读性还是留着吧
	{
		v1.POST("/reorder/bubble", bubbleReorder)
		v1.POST("/reorder/select", selectReorder)
		v1.POST("/reorder/insert", insertReorder)

	}
}

type Param struct {
	Numbers []int `form:"numbers"`
}

func bubbleReorder(ctx *gin.Context) {
	var p Param
	ctx.ShouldBind(&p)
	jd := json.NewEncoder(os.Stdout)
	jd.Encode(p)
	pj, _ := json.Marshal(p)
	log.Println(string(pj))
	ctx.JSON(200, handler.BubbleReorder(p.Numbers))
}

func selectReorder(ctx *gin.Context) {
	var p Param
	ctx.ShouldBind(&p)
	ctx.JSON(200, handler.SelectReorder(p.Numbers))
}

func insertReorder(ctx *gin.Context) {
	var p Param
	ctx.ShouldBind(&p)
	ctx.JSON(200, handler.InsertReorder(p.Numbers))
}
