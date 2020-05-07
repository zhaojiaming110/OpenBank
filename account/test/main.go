package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	us "github.com/zhaojiaming110/openBank/account/account-srv/proto/account"
	"net/http"
	log "github.com/micro/go-micro/v2/logger"
	"context"
)

var (
	serviceClient us.AccountService
)


func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":9091", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// 调用后台服务
	fmt.Println("hello linux")
	service := micro.NewService()
	service.Init()
	serviceClient = us.NewAccountService("openbank.srv.account", service.Client())
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: "10002",
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Info("调用后台服务error")
		return
	}
	log.Info("调用后台服务结束")
	log.Info(rsp)
}