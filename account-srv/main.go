package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/broker/nsq/v2"
	z "github.com/zhaojiaming110/openBank/plugins/zap"

	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/zhaojiaming110/openBank/account-srv/handler"
	"github.com/zhaojiaming110/openBank/account-srv/model"
	s "github.com/zhaojiaming110/openBank/account-srv/proto/account"
	"github.com/zhaojiaming110/openBank/basic"
	"github.com/zhaojiaming110/openBank/basic/common"
	"github.com/zhaojiaming110/openBank/basic/config"
	"go.uber.org/zap"
)

var (
	appName = "user_srv"
	cfg     = &userCfg{}
	log		= z.GetLogger()
)

type userCfg struct {
	common.AppCfg
}

func main() {
	// 本地初始化配置、数据库等信息
	//basic_local.Init()


	// grpc初始化配置、数据库等信息
	initCfg()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	nsqBorker := getNsqBroker()

	// 新建服务
	service := micro.NewService(
		micro.Name("go.micro.openbank.account.arv"),
		micro.Registry(micReg),
		micro.Version("latest"),
		micro.Broker(nsqBorker),
		//micro.WrapHandler(openTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
	)

	// Register Handler   注册服务
	s.RegisterAccountHandler(service.Server(), new(handler.Service))

	// Run service    启动服务
	if err := service.Run(); err != nil {
		log.Fatal("service run failed")
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

func getNsqBroker() broker.Broker {
	nsqCfg := &common.Nsq{}
	err := config.C().App("nsq", nsqCfg)
	log.Info("[nsqConfig]配置", zap.Any("nsqConfig", nsqCfg))
	if err != nil {
		panic(err)
	}
	return nsq.NewBroker(broker.Addrs([]string{fmt.Sprintf("%s:%d", nsqCfg.Host, nsqCfg.Port)}...))
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
