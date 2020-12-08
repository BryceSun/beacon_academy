package init_config

import (
	"github.com/BryceSun/beacon_academy/internal/common/utils"
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
	root.Use(Authorize)
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

func Authorize(c *gin.Context) {
	if inWhiteList(c) {
		return
	}
	token := c.GetHeader("authorization")
	if len(token) == 0 {
		c.JSON(302, gin.H{"error": "token is needed"})
		c.Abort()
		return
	}
	t, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, utils.GetKey)
	if err != nil {
		c.JSON(302, err)
		c.Abort()
		return
	}
	c.Set("username", t.Claims.(jwt.MapClaims)["unm"])
	c.Next()
	status := c.Writer.Status()
	log.Println(status)
}

//
//func authorize() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t := time.Now()
//
//		// Set example variable
//		//c.Set("example", "12345")
//
//		// before request
//
//		c.Next()
//
//		// after request
//		latency := time.Since(t)
//		log.Print(latency)
//
//		// access the status we are sending
//		status := c.Writer.Status()
//		log.Println(status)
//	}
//}
