package config

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	err	error
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	etcdConfig              defaultEtcdConfig
	mysqlConfig             defaultMysqlConfig
	nsqConfig               defaultNsqConfig
	redisConfig             defaultRedisConfig
	profiles                defaultProfiles
	m                       sync.RWMutex
	inited                  bool
	//sp                      = string(filepath.Separator)
)


// Init 初始化配置
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[Init] 配置已经初始化过")
	}

	// 加载yml配置
	// 先加载基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))
	log.Info(appPath)

	pt := filepath.Join(appPath, "conf")
	log.Info(pt)

	os.Chdir(appPath)


	// 找到application.yml文件
	if err = config.Load(file.NewSource(file.WithPath(pt + "/" + "application.yml"))); err != nil {
		panic(err)
	}

	// 找到需要引入的新配置文件
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Info("[init]加载配置文件 path: ", pt+"/application.yml ", profiles)

	// 开始导入新文件
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")
		sources := make([]source.Source, len(include))

		for i := 0; i < len(include); i++ {
			log.Info(include[i])
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"

			fmt.Printf(filePath + "\n")

			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		// 加载include的文件
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}

	// 赋值
	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)
	config.Get(defaultRootPath, "nsq").Scan(&nsqConfig)
	config.Get(defaultRootPath, "redis").Scan(&redisConfig)

	log.Info(etcdConfig)
	log.Info(mysqlConfig)
	log.Info(nsqConfig)
	log.Info(redisConfig)

	// 标记已经初始化
	inited = true

	log.Info("config 初始化完成")
}

// GetMysqlConfig 获取mysql配置
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetEtcdConfig 获取Etcd配置
func GetEtcdConfig() (ret EtcdConfig) {
	return etcdConfig
}

// 获取Nsq配置
func GetNsqConfig() (ret NsqConfig) {
	return	nsqConfig
}

// GetRedisConfig 获取Redis配置
func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}
