package config

import (
	"fmt"
	"github.com/micro/go-micro/v2/registry"
)

// EtcdConfig Etcd 配置
type EtcdConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

// defaultEtcdConfig 默认Etcd 配置
type defaultEtcdConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// GetPort Etcd 端口
func (c defaultEtcdConfig) GetPort() int {
	return c.Port
}

// GetEnabled Etcd 激活
func (c defaultEtcdConfig) GetEnabled() bool {
	return c.Enabled
}

// GetHost Etcd 主机地址
func (c defaultEtcdConfig) GetHost() string {
	return c.Host
}

// ETCD注册函数
func RegistryOptions(ops *registry.Options) {
	etcdCfg := GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}