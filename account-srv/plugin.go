package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/zhaojiaming110/openBank/plugins/db"
	_ "github.com/zhaojiaming110/openBank/plugins/redis"
	_ "github.com/zhaojiaming110/openBank/plugins/zap"
)