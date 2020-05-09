package db

import (
	"database/sql"
	"github.com/zhaojiaming110/openBank/basic-local/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/micro/go-micro/v2/logger"
)

func initMysql() {
	var err error

	// 创建连接
	log.Info("initMysql()")
	log.Info(config.GetMysqlConfig().GetURL())
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Info("Connect Mysql success")

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	//连接数据库闲置断线的问题
	mysqlDB.SetConnMaxLifetime(time.Second * config.GetMysqlConfig().GetConnMaxLifetime())
	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
