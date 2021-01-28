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
		v1.POST("/sort/bubble", bubbleSort)
		v1.POST("/sort/select", selectSort)
		v1.POST("/sort/insert", insertSort)
		v1.POST("/sort/merge", mergeSort)
	}
}

type Param struct {
	Numbers []int `form:"numbers"`
}

func bubbleSort(ctx *gin.Context) {
	var p Param
	ctx.ShouldBind(&p)
	jd := json.NewEncoder(os.Stdout)
	jd.Encode(p)
	pj, _ := json.Marshal(p)
	log.Println(string(pj))
	ctx.JSON(200, handler.BubbleSort(p.Numbers))
}

func selectSort(ctx *gin.Context) {
	var p Param
	ctx.ShouldBind(&p)
	ctx.JSON(200, handler.SelectSort(p.Numbers))
}

func insertSort(ctx *gin.Context) {
	var p Param
	ctx.ShouldBind(&p)
	ctx.JSON(200, handler.InsertSort(p.Numbers))
}

func mergeSort(ctx *gin.Context) {
	var p Param
	ctx.ShouldBind(&p)
	ctx.JSON(200, handler.MergeSort(p.Numbers))
}
