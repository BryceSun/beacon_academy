package init_config

import (
	. "github.com/BryceSun/beacon_academy/internal/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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
	case path == "/account/info/v1" && c.Request.Method == "POST":
		return true
	case path == "/account/info/v2" && c.Request.Method == "POST":
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
	t, err := jwt.ParseWithClaims(token, &ClaimsPlus{}, GetKey)
	if err != nil {
		c.JSON(Output2(err))
		c.Abort()
		return
	}
	claims := t.Claims.(*ClaimsPlus)
	c.Set("username", claims.Username)
	c.Set("userId", claims.UserId)
	//检查redis中的token列表
	tokenKey := TokenKey(claims.UserId)
	val, err := Redis.Get(tokenKey).Result()
	if err != nil {
		c.JSON(Output2(err))
		c.Abort()
		return
	}
	if val != token {
		c.JSON(Output2(&NotAuthorized))
		c.Abort()
		return
	}
	iat := claims.IssuedAt
	du := time.Since(time.Unix(iat, 0))
	if du >= TokenLiveTime && c.Request.URL.Path != "/account/token" {
		token, _ = UpdateToken(*claims)
		//将新token加入redis
		Redis.Set(tokenKey, token, TokenLiveTime)
		c.Header("authorization", token)
		c.Header("Access-Control-Expose-Headers", "Authorization")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Connection, User-Agent, Cookie, Authorization")
	}
}
