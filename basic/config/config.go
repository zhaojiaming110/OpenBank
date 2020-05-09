package config

import (
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"sync"

	"github.com/micro/go-micro/v2/config"
)

var (
	m      sync.RWMutex
	inited bool

	// 默认配置器
	c = &configurator{}
)

// Configurator 配置器
type Configurator interface {
	App(name string, config interface{}) (err error)
	Path(path string, config interface{}) (err error)
}

// configurator 配置器
type configurator struct {
	conf config.Config
	appName string
}

func (c *configurator) App(name string, config interface{}) (err error) {

	v := c.conf.Get(name)
	if v != nil {
		err = v.Scan(config)
	} else {
		err = fmt.Errorf("[App] 配置不存在，err：%s", name)
	}

	return
}

func (c *configurator) Path(path string, config interface{}) (err error) {
	v := c.conf.Get(c.appName, path)
	if v != nil {
		err = v.Scan(config)
	} else {
		err = fmt.Errorf("[Path] 配置不存在，err：%s", path)
	}

	return
}

// c 配置器
func C() Configurator {
	return c
}

func (c *configurator) init(ops Options) (err error) {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[init] 配置已经初始化过")
		return
	}

	c.conf, _ = config.NewConfig()
	// 加载配置
	err = c.conf.Load(ops.Sources...)
	if err != nil {
		log.Fatal("加载配置出错")
	}

	go func() {
		log.Info("[init] 侦听配置变动 ...")

		// 开始侦听变动事件
		watcher, err := c.conf.Watch()
		if err != nil {
			log.Fatal("侦听配置出错")
		}

		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatal("侦听配置出错")
			}

			log.Fatal("侦听配置变动%v", v)
		}
	}()

	// 标记已经初始化
	inited = true
	return
}

// Init 初始化配置
func Init(opts ...Option) {

	ops := Options{}
	for _, o := range opts {
		o(&ops)
	}

	c = &configurator{}

	c.init(ops)
}
