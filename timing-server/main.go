package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/zhaojiaming110/openBank/basic"
	"github.com/zhaojiaming110/openBank/basic/common"
	"github.com/zhaojiaming110/openBank/basic/config"
	z "github.com/zhaojiaming110/openBank/plugins/zap"
	"github.com/zhaojiaming110/openBank/timing-server/handler"
	"github.com/zhaojiaming110/openBank/timing-server/model"
	proto "github.com/zhaojiaming110/openBank/timing-server/proto/timing"
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
		micro.Name("micro.open.bank.service.timing"),
		micro.Version("latest"),
		micro.Registry(micReg),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) error {
			model.Init()
			handler.Init()
			return nil
		}),
	)

	// Register Handler
	proto.RegisterTimingHandler(service.Server(), new(handler.Server))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal("service run failed", zap.Any("err", err))
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
