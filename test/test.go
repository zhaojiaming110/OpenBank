// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/5/6 下午3:24

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
* Barrier
 */
type barrierResp struct {
	Err error
	Resp string
	Status int
}

// 构造请求
func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}

	client := http.Client{
		Timeout: time.Duration(2*time.Microsecond),
	}

	resp, err := client.Get(url)
	if resp != nil {
		res.Status = resp.StatusCode
	}
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}

// 合并结果
func barrier(endpoints ...string) {
	requestNumber := len(endpoints)

	in := make(chan barrierResp, requestNumber)
	response := make([]barrierResp, requestNumber)

	defer close(in)

	for _, endpoints := range endpoints {
		go makeRequest(in, endpoints)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err, resp.Status)
			hasError = true
		}
		response[i] = resp
	}
	if !hasError {
		for _, resp := range response {
			fmt.Println(resp.Status)
		}
	}
}

func main() {
	barrier([]string{"https://www.baidu.com", "https://www.baidu.com", "https://www.baidu.com"}...)
}