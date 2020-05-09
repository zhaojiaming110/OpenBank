// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/8 下午2:02

package db

import (
	"database/sql"
	"github.com/zhaojiaming110/openBank/basic"
	z "github.com/zhaojiaming110/openBank/plugins/zap"
	"go.uber.org/zap"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
	log = z.GetLogger()
)

func init() {
	basic.Register(initDB)
}

// initDB 初始化数据库
func initDB() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		log.Info("mysql已初始化过", zap.Any("err", err))
		return
	}

	initMysql()

	inited = true
}

// GetDB 获取db
func GetDB() *sql.DB {
	return mysqlDB
}