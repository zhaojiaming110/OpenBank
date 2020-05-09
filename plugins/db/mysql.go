// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/8 下午2:03
package db

import (
	"database/sql"
	"github.com/zhaojiaming110/openBank/basic/config"
	"go.uber.org/zap"
	"time"
	)



// Mysql mySQL 配置
type Mysql struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
	ConnMaxLifetime time.Duration    `json:"connMaxLifetime"`
}

func initMysql() {
	log.Info("[initMysql] 初始化Mysql")

	c := config.C()
	cfg := &Mysql{}

	err := c.App("mysql", cfg)
	log.Info("mysql配置", zap.Any("mysqlConfig", cfg))
	if err != nil {
		log.Fatal("[initMysql] ", zap.Any("err", err))
	}

	if !cfg.Enable {
		log.Fatal("[initMysql] 未启用Mysql")
		return
	}

	// 创建连接
	mysqlDB, err = sql.Open("mysql", cfg.URL)
	if err != nil {
		log.Fatal("连接失败", zap.Any("err", err))
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(cfg.MaxOpenConnection)

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(cfg.MaxIdleConnection)

	//连接数据库闲置断线的问题
	mysqlDB.SetConnMaxLifetime(time.Second * cfg.ConnMaxLifetime)
	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal("激活失败", zap.Any("err", err))
	}

	log.Info("[initMysql] Mysql 连接成功")
}