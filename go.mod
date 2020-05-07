module github.com/zhaojiaming110/openBank

go 1.13

require (
	github.com/go-redis/redis/v7 v7.2.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/go-plugins/broker/nsq/v2 v2.5.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.5.0
	google.golang.org/grpc v1.26.0
)

exclude github.com/micri/go-micro v1.16.0
