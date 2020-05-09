package subscriber

import (
	"context"
	proto "github.com/zhaojiaming110/openBank/asynchronous-server/proto/asynchronous"
	"github.com/zhaojiaming110/openBank/plugins/db"
	z "github.com/zhaojiaming110/openBank/plugins/zap"
	"go.uber.org/zap"
)

type Asynchronous struct{}

var (
	log = z.GetLogger()
)

func (e *Asynchronous) Handle(ctx context.Context, msg *proto.Message) error {
	return Handler(ctx, msg)
}

func Handler(ctx context.Context, msg *proto.Message) error {
	log.Info("异步开户信息", zap.Any("msg", msg))
	mysql := db.GetDB()
	_, err := mysql.Exec("INSERT INTO account(user_id,user_name,pwd) VALUE (?,?,?)", msg.Id,msg.Name,msg.Pwd)
	if err != nil {
		log.Fatal("登记失败", zap.Any("err", err))
		return err
	}
	return nil
}
