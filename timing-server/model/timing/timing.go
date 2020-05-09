// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/8 下午9:57

package timing

import (
	"bufio"
	"fmt"
	"github.com/zhaojiaming110/openBank/plugins/db"
	"github.com/zhaojiaming110/openBank/plugins/redis"
	z "github.com/zhaojiaming110/openBank/plugins/zap"
	"go.uber.org/zap"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	s *service
	m sync.RWMutex
	log = z.GetLogger()

	key1 = "key1"
	key2 = "key2"
	diffKey1Key2 = "diffKey1Key2"
	diffKey2Key1 = "diffKey2Key1"
	unionKey1Key2 = "unionKey1Key2"
	key1Elements []string
	key2Elements []string

)

type Service interface {
	CheckAccounts(data string) error
}

type service struct{}

func (s *service) CheckAccounts(data string) error {
	client := redis.Redis()
	pipe := client.Pipeline()
	file1, err := os.Open("./1.txt")
	if err != nil {
		log.Fatal("open file1 error")
		return err
	}
	defer file1.Close()

	reader := bufio.NewReader(file1)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				key1Elements = append(key1Elements, line)
			}
			log.Info("file has read!")
			break
		}
		if err != nil {
			log.Fatal("read file failed")
			return err
		}
		key1Elements = append(key1Elements, line)
	}

	file2, err := os.Open("./2.txt")
	if err != nil {
		log.Fatal("open file error")
		return nil
	}
	defer file1.Close()

	reader2 := bufio.NewReader(file2)
	for {
		line, err := reader2.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				key2Elements = append(key2Elements, line)
			}
			log.Info("file has read")
			break
		}
		if err != nil {
			log.Fatal("read file failed")
			return nil
		}
		key2Elements = append(key2Elements, line)
	}
	pipe.SAdd(key1, key1Elements)
	pipe.SAdd(key2, key2Elements)
	pipe.SUnionStore(unionKey1Key2, key1, key2)
	pipe.SDiffStore(diffKey1Key2, key1, key2)
	pipe.SDiffStore(diffKey2Key1, key2, key1)
	_, err = pipe.Exec()
	if err != nil {
		log.Fatal("pipe err", zap.Any("err", err))
	}

	diffKey1Key2s := client.SMembers(diffKey1Key2).Val()
	diffKey2Key1s := client.SMembers(diffKey2Key1).Val()
	fmt.Println(diffKey1Key2s)
	fmt.Println(diffKey2Key1s)

	mapDiffKey1Key2 := make(map[string][]string, 3)
	for _, v := range diffKey1Key2s {
		res := strings.Split(v, "|")
		mapDiffKey1Key2[res[0]] = res
	}

	mapDiffKey2Key1 := make(map[string][]string, 3)
	for _, v := range diffKey2Key1s {
		res := strings.Split(v, "|")
		mapDiffKey2Key1[res[0]] = res
	}

	sql := db.GetDB()
	insertString := "INSERT INTO check_accounts(list,money1,money2,state1,state2) VALUE (?,?,?,?,?)"
	stmt, err := sql.Prepare(insertString)
	if err != nil {
		log.Fatal("prepare failed, err[%s]", zap.Any("err", err))
	}
	defer stmt.Close()


	for index, val := range mapDiffKey1Key2 {
		// 判断是否出现在 mapDiffKey2Key1中
		if v, ok := mapDiffKey2Key1[index]; ok {
			// 流水同时存在，状态或金额不一致
			fmt.Println("流水同时存在，状态或金额不一致, map1: ", val, " map2 ", v)
			stmt.Exec(val[0], val[1], v[1], val[2], v[2])
		} else {
			// map1存在流水，map1多单的情况
			fmt.Println("map1多单: ", val)
			sql.Exec("INSERT INTO check_accounts(list,money1,state1) VALUE (?,?,?)",val[0],val[1],val[2])
		}
	}

	for index, val := range mapDiffKey2Key1 {
		// 判断是否出现在 mapDiffKey2Key1中
		if _, ok := mapDiffKey1Key2[index]; !ok {
			// map2存在流水，map2多单的情况
			fmt.Println("map2多单: ", val)
			sql.Exec("INSERT INTO check_accounts(list,money2,state2) VALUE (?,?,?)",val[0],val[1],val[2])
		}
		// 同时存在已经在上个循环处理过了。
	}

	return nil
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService]未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}
	s = &service{}
}
