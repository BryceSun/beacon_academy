package init_config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Configure configure

func init() {
	//f,err := os.Open(os.Args[1])
	f, err := os.Open("D:\\It_project\\go_projects\\beacon_academy\\configs\\conf_dev.json")
	//f, err := os.Open("D:\\projects\\my_module\\bgs_server\\conf\\conf_dev.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	Configure = configure{}
	err = decoder.Decode(&Configure)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Configure.Server.Name)
}

type configure struct {
	Server serverConf
	Mysql  mysqlConf
	Redis  redisConf
	Email  emailConf
}

type serverConf struct {
	Name string
	Port int
}

type mysqlConf struct {
	Host     string
	Port     int
	Username string
	Password string
}

type redisConf struct {
	Host     string
	Port     int
	Password string
	Db       int
}

type emailConf struct {
	Host     string
	Port     int
	Account  string
	Password string
}

func ServerConf() serverConf {
	return Configure.Server
}

func MysqlConf() mysqlConf {
	return Configure.Mysql
}

func RedisConf() redisConf {
	return Configure.Redis
}

func EmailConf() emailConf {
	return Configure.Email
}
