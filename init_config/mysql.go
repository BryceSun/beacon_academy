package init_config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var Db *sql.DB

func init() {
	configure := MysqlConf()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=30s",
		configure.Username, configure.Password, configure.Host, configure.Port, "bgs"))
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(time.Hour * 24)
	Db = db
}
