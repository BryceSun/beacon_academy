package init_config

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

var Redis *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisConf().Host + ":" + strconv.Itoa(RedisConf().Port),
		Password: RedisConf().Password,
		DB:       RedisConf().Db,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	Redis = client
}
