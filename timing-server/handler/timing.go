package handler

import (
	"context"
	z "github.com/zhaojiaming110/openBank/plugins/zap"
	"github.com/zhaojiaming110/openBank/timing-server/model/timing"
	proto "github.com/zhaojiaming110/openBank/timing-server/proto/timing"
	"go.uber.org/zap"
)

var (
	timingService timing.Service
	log           = z.GetLogger()
)

func Init() {
	var err error
	timingService, err = timing.GetService()
	if err != nil {
		log.Fatal("初始化handler错误", zap.Any("err", err))
		return
	}
}

type Server struct{}

func (s *Server) CheckAccounts(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	err := timingService.CheckAccounts(req.Data)
	if err != nil {
		rsp.Result = "failed"
	} else {
		rsp.Result = "success"
	}
	return nil
}

