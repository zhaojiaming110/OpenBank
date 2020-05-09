package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/broker/nsq/v2"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/zhaojiaming110/openBank/asynchronous-server/subscriber"
	"github.com/zhaojiaming110/openBank/basic"
	"github.com/zhaojiaming110/openBank/basic/common"
	"github.com/zhaojiaming110/openBank/basic/config"
	z "github.com/zhaojiaming110/openBank/plugins/zap"
	"go.uber.org/zap"
)

var (
	appName = "timing_srv"
	cfg     = &userCfg{}
	log		= z.GetLogger()
)

type userCfg struct {
	common.AppCfg
}

func main() {

	initCfg()

	micReg := etcd.NewRegistry(registryOptions)


	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.asynchronous"),
		micro.Version("latest"),
		micro.Registry(micReg),
		micro.Broker(brokerOption()),
	)

	// Initialise service
	service.Init()

	// Register Struct as Subscriber
	micro.RegisterSubscriber("OPEN_ACCOUNT", service.Server(), new(subscriber.Asynchronous))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal("server run err", zap.Any("err", err))
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}

func brokerOption() broker.Broker {
	nsqConfig := &common.Nsq{}
	err := config.C().App("nsq", nsqConfig)
	if err != nil {
		panic(err)
	}
	addrs := []string{fmt.Sprintf("%s:%d", nsqConfig.Host, nsqConfig.Port)}
	return nsq.NewBroker(broker.Addrs(addrs...))
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Info("[initCfg]配置", zap.Any("cfg", cfg))

	return
}
