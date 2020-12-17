package init_config

import (
	. "github.com/BryceSun/beacon_academy/internal/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
)

var RootRouter *gin.RouterGroup

var Engine *gin.Engine

func init() {
	r := gin.Default()
	Engine = r
	root := r.Group("/")
	root.Use(authorize)
	RootRouter = root
}

func inWhiteList(c *gin.Context) bool {
	path := c.Request.URL.Path
	log.Println(path)
	switch {
	//register
	case path == "/account/info" && c.Request.Method == "POST":
		return true
		//login
	case path == "/account/token" && c.Request.Method == "POST":
		return true
	}
	return false
}

func authorize(c *gin.Context) {
	if inWhiteList(c) {
		return
	}
	token := c.GetHeader("authorization")
	if len(token) == 0 {
		c.JSON(Output2(ParamsAreNeeded.New("token缺失")))
		c.Abort()
		return
	}
	t, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, GetKey)
	if err != nil {
		c.JSON(302, err)
		c.Abort()
		return
	}
	c.Set("username", t.Claims.(ClaimsPlus).Username)
	c.Set("userId", t.Claims.(ClaimsPlus).UserId)
	c.Next()

	status := c.Writer.Status()
	log.Println(status)
}
