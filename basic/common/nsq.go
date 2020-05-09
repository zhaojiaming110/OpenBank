// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/8 下午2:39

package common

// Nsq配置
type Nsq struct{
	Enabled bool `json:"enabled"`
	Host string	`json:"host"`
	Port int 	`json:"port"`
}