module github.com/zhaojiaming110/openBank

go 1.13

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/go-redis/redis/v7 v7.2.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/go-plugins/broker/nsq/v2 v2.5.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.5.0
	github.com/micro/go-plugins/micro/cors/v2 v2.5.0 // indirect
	github.com/micro/go-plugins/wrapper/breaker/hystrix v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.5.0
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.5.0 // indirect
	github.com/micro/micro/v2 v2.6.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/common v0.6.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/zap v1.13.0
	google.golang.org/grpc v1.26.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

exclude github.com/micri/go-micro v1.16.0
