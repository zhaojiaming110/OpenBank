// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/23 上午9:50

package main

import (
	"github.com/go-redis/redis/v7"
)


func main() {

	client := redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "",
		DB:                 0,
	})

	rediskey := "user" + "18595050638"
	data := client.Get("18595050638-new").Val()
	var userMessage = []string{
		"count",
		"1",
		"code",
		"110120",
		"data",
		data,
	}

	client.HMSet(rediskey, userMessage)

}
