package basic_local

import (
	"github.com/zhaojiaming110/openBank/account/basic-local/config"
	"github.com/zhaojiaming110/openBank/account/basic-local/db"
	"github.com/zhaojiaming110/openBank/account/basic-local/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}