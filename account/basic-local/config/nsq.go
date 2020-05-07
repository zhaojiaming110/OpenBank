// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/7 下午11:58

package config

import (
	"fmt"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/nsq/v2"
)

type NsqConfig interface {
	GetEnabled() bool
	GetHost() string
	GetPort() int
}

type defaultNsqConfig struct{
	Enabled bool `json:"enabled"`
	Host string	`json:"host"`
	Port int 	`json:"port"`
}

func (d defaultNsqConfig) GetEnabled() bool {
	return d.Enabled
}

func (d defaultNsqConfig) GetHost() string {
	return d.Host
}

func (d defaultNsqConfig) GetPort() int {
	return d.Port
}

func GetNsqBroker() broker.Broker {
	ret := GetEtcdConfig()
	return nsq.NewBroker(broker.Addrs([]string{fmt.Sprintf("%s:%d", ret.GetHost(), ret.GetPort())}...))
}


