// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/5 上午12:16

package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/broker/nsq/v2"
	proto "github.com/zhaojiaming110/openBank/account/account-srv/proto/account"
)
type Sub struct {
}

func (s *Sub) Process(ctx context.Context, evt *proto.Request) error {
	log.Info("hhhh")
	log.Info("Receive info: I %s & Timestamp %s  %s\n", evt.UserID, evt.UserPwd, evt.UserName)
	return nil
}


func main() {
	srv := micro.NewService(
		micro.Name("go.oepn.bank.client"),
		micro.Broker(nsq.NewBroker(
			broker.Addrs([]string{"127.0.0.1:4150"}...),
		)),
		)
	srv.Init()

	sOpts := broker.NewSubscribeOptions(
		nsq.WithMaxInFlight(5),
	)

	_ = micro.RegisterSubscriber("demo.message", srv.Server(), &Sub{}, server.SubscriberContext(sOpts.Context))

	if err := srv.Run(); err != nil {
		log.Fatal("error occurs: %v", err)
	}
}