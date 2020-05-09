package db

import (
	"database/sql"
	"fmt"
	"github.com/zhaojiaming110/openBank/basic-local/config"
	"sync"

	log "github.com/micro/go-micro/v2/logger"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	log.Info("begin 初始化db")

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Error(err)
		return
	}

	// 如果配置声明使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}

	inited = true

	log.Info("db 初始化完成")
}

// GetDB 获取db
func GetDB() *sql.DB {
	return mysqlDB
}