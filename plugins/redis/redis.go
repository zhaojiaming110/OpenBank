// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/8 下午2:08

package redis

import (
	"strings"
	"sync"

	r "github.com/go-redis/redis/v7"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/zhaojiaming110/openBank/basic"
	"github.com/zhaojiaming110/openBank/basic/config"
)

var (
	client *r.Client
	m      sync.RWMutex
	inited bool
)

// redis redis 配置
type redis struct {
	Enabled  bool           `json:"enabled"`
	Conn     string         `json:"conn"`
	Password string         `json:"password"`
	DBNum    int            `json:"dbNum"`
	Timeout  int            `json:"timeout"`
	Sentinel *RedisSentinel `json:"sentinel"`
}

type RedisSentinel struct {
	Enabled bool   `json:"enabled"`
	Master  string `json:"master"`
	XNodes  string `json:"nodes"`
	nodes   []string
}

// Nodes redis 哨兵节点列表
func (s *RedisSentinel) GetNodes() []string {
	if len(s.XNodes) != 0 {
		for _, v := range strings.Split(s.XNodes, ",") {
			v = strings.TrimSpace(v)
			s.nodes = append(s.nodes, v)
		}
	}
	return s.nodes
}

// init 初始化Redis
func init() {
	basic.Register(initRedis)
}

func initRedis() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[initRedis] 已经初始化过Redis...")
		return
	}

	log.Info("[initRedis] 初始化Redis...")

	c := config.C()
	cfg := &redis{}
	err := c.App("redis", cfg)
	log.Info(cfg)
	log.Info(cfg.Sentinel)
	if err != nil {
		log.Warnf("[initRedis] %s", err)
	}

	if !cfg.Enabled {
		log.Info("[initRedis] 未启用redis")
		return
	}

	// 加载哨兵模式
	if cfg.Sentinel != nil && cfg.Sentinel.Enabled {
		log.Info("[initRedis] 初始化Redis，哨兵模式...")
		initSentinel(cfg)
	} else { // 普通模式
		log.Info("[initRedis] 初始化Redis，普通模式...")
		initSingle(cfg)
	}

	log.Info("[initRedis] 初始化Redis，检测连接...")

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("[initRedis] 初始化Redis，检测连接Ping... %s", pong)
}

// Redis 获取redis
func Redis() *r.Client {
	return client
}

func initSentinel(redisConfig *redis) {
	client = r.NewFailoverClient(&r.FailoverOptions{
		MasterName:    redisConfig.Sentinel.Master,
		SentinelAddrs: redisConfig.Sentinel.GetNodes(),
		DB:            redisConfig.DBNum,
		Password:      redisConfig.Password,
	})

}

func initSingle(redisConfig *redis) {
	client = r.NewClient(&r.Options{
		Addr:     redisConfig.Conn,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DBNum,    // use default DB
	})
}