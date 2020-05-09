package account

import (
	log "github.com/micro/go-micro/v2/logger"
	proto "github.com/zhaojiaming110/openBank/account-srv/proto/account"
	"github.com/zhaojiaming110/openBank/basic-local/db"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	//queryString := `SELECT user_id, user_name, pwd FROM account WHERE user_name = ?`

	// 获取数据库
	//o := db.GetDB()

	//ret = &proto.User{}

	log.Info("hello, world")
	log.Info(userName)

	//pub := micro.NewEvent("demo.message", client.DefaultClient)
	//
	//go func() {
	//	time.Sleep(1*time.Second)
	//	log.Info("publish")
	//	err := pub.Publish(context.TODO(), &proto.Request{
	//		UserID:               "140723199301120036",
	//		UserName:             "zjm",
	//		UserPwd:              "110120",
	//	})
	//	fmt.Println(err)
	//}()

	// 查询
	//err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	//if err != nil {
	//	log.Error("[QueryUserByName] 查询数据失败，err：%s", err)
	//	return
	//}
	log.Info("查询结束")
	return
}

func (s *service) CreateUser(req *proto.Request) error {
	insertString := `INSERT INTO account (user_id, user_name, pwd) VALUES(?, ?, ?)`

	mysql := db.GetDB()
	stmt, err := mysql.Prepare(insertString)
	if err != nil {
		log.Error("prepare failed, err[%s]", err)
	}

	res, err := stmt.Exec(req.UserID, req.UserID, req.UserPwd)
	defer stmt.Close()

	if err != nil {
		log.Error("exec failed, err[%s]", err)
	}


	r, err := res.RowsAffected()
	if err != nil {
		log.Error("insert failed err: [%s]", err)
	}
	if r == 1 {
		log.Info("insert success")
	} else {
		log.Info("insert success.")
	}

	return nil
}
