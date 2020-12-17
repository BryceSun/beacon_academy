package main

import (
	"github.com/BryceSun/beacon_academy/init_config"
	_ "github.com/BryceSun/beacon_academy/internal/account/router"
	"strconv"
)

func main() {
	init_config.Engine.Run(":" + strconv.Itoa(init_config.ServerConf().Port))
}
