package main

import (
	"github.com/BryceSun/beacon_academy/init_config"
	_ "github.com/BryceSun/beacon_academy/internal/account/router"
)

func main() {
	init_config.Engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
