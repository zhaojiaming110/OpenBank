package handler

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	acc "github.com/zhaojiaming110/openBank/account-srv/model/account"
	s "github.com/zhaojiaming110/openBank/account-srv/proto/account"
	proto "github.com/zhaojiaming110/openBank/asynchronous-server/proto/asynchronous"
	"os"
)

type Service struct{}

var (
	accountService acc.Service
)

// Init 初始化handler
func Init() {
	var err error
	accountService, err = acc.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	log.Info("欢迎进入后台查询接口")
	file, err := os.OpenFile("11.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	file.Write(req.Idfile)
	user, err := accountService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return err
	}

	rsp.User = user
	rsp.Success = true

	log.Info("欢迎下次再来")
	return nil
}

// CreateUser 用户注册
func (e *Service) CreateUser(ctx context.Context, req *s.Request, rsp *s.Response) error {
	log.Info("欢迎CreateuSER")
	log.Info(req.Flag)
	log.Info(req)
	//if req.Flag == "1" {	// 异步处理
		log.Info("欢迎异步处理")
		pub := micro.NewEvent("OPEN_ACCOUNT", client.DefaultClient)
		err := pub.Publish(ctx, &proto.Message{
			Id:                   "62270021",
			Name:                 "test01",
			Pwd:                  "121231",
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Info("发送异步消息成功")
		return nil
	//}
	//if err := accountService.CreateUser(req); err != nil {
	//	rsp.Success = false
	//	rsp.Error = &s.Error{
	//		Code:   500,
	//		Detail: err.Error(),
	//	}
	//	return nil
	//}

	rsp.Success = true

	return nil
}
