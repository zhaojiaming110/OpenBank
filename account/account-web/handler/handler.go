package handler

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"io"
	"net/http"
	"time"

	hystrix_go "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	us "github.com/zhaojiaming110/openBank/account/account-srv/proto/account"
)

var (
	serviceClient us.AccountService
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1
	cl := hystrix.NewClientWrapper()(client.DefaultClient)
	serviceClient = us.NewAccountService("open.bank.demo2", cl)
}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {
	log.Info("欢迎来到web服务")
	// 只接受POST请求
	if r.Method != "POST" {
		log.Error("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("pngqqq")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		content = append(content, tmp[:n]...)
	}

	fmt.Println(handler.Filename)

	// 调用后台服务
	log.Info("开始调用后台服务")
	_, err = serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: r.Form.Get("userName"),
		Idfile:   content,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Info("调用后台服务error")
		return
	}
	log.Info("调用后台服务结束")

	// 返回结果
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	response["success"] = true

	//if rsp.User.Pwd == r.Form.Get("pwd") {
	//	response["success"] = true
	//	// 干掉密码返回
	//	rsp.User.Pwd = ""
	//	response["data"] = rsp.User
	//
	//} else {
	//	response["success"] = false
	//	response["error"] = &Error{Detail: "密码错误"}
	//}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
