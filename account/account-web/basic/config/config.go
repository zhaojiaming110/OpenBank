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

	pt := filepath.Join(appPath, "conf")
	log.Info("hello,world")
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

	log.Info(etcdConfig)

	// 标记已经初始化
	inited = true

	log.Info("config 初始化完成")
}


// GetEtcdConfig 获取Etcd配置
func GetEtcdConfig() (ret EtcdConfig) {
	return etcdConfig
}
