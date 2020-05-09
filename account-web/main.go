package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/zhaojiaming110/openBank/account-web/handler"
	"github.com/zhaojiaming110/openBank/basic"
	"github.com/zhaojiaming110/openBank/basic/common"
	"github.com/zhaojiaming110/openBank/basic/config"
	"github.com/zhaojiaming110/openBank/plugins/breaker"
	"github.com/zhaojiaming110/openBank/plugins/tracer/opentracing/std2micro"
	z "github.com/zhaojiaming110/openBank/plugins/zap"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	appName = "user_web"
	cfg     = &userCfg{}
	log		= z.GetLogger()
)

type userCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	//t, io, err := tracer.NewTracer(cfg.Name, "")
	//if err != nil {
	//	log.Fatal("tracer err", zap.Any("err", err))
	//}
	//defer io.Close()
	//opentracing.SetGlobalTracer(t)

	// 创建新服务
	service := web.NewService(
		web.Name("go.micro.web.account"),
		web.RegisterTTL(15*time.Second),
		web.RegisterInterval(10*time.Second),
		web.Version("latest"),
		web.Registry(micReg),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			// 初始化handler
			handler.Init()
		}),
	); err != nil {
		log.Fatal("初始化服务失败", zap.Any("err", err))
	}

	// 注册登录接口
	//设置采样率
	std2micro.SetSamplingFrequency(50)
	handlerLogin := http.HandlerFunc(handler.Login)
	service.Handle("/user/login", std2micro.TracerWrapper(breaker.BreakerWrapper(handlerLogin)))

	// hystrix-dashboard 将熔断器状态可视化
	//hystrixStreamHandler := hystrix.NewStreamHandler()
	//hystrixStreamHandler.Start()
	//go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal("运行失败", zap.Any("err", err))
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

	log.Info("[initCfg]配置完成")

	return
}
