// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/24 下午12:52

package main

import (
	"fmt"
	"math/rand"
)

type job struct {
	Id int
	RandNum int
}

type Result struct {
	*job
}

func main() {
	jobChan := make(chan *job, 128)
	resultChan := make(chan *Result, 128)
	createPoll(64, jobChan, resultChan)
	go func(resultChan <-chan *Result) {
		for v := range resultChan {
			fmt.Printf("job id:%d randum:%d\n", v.Id, v.RandNum)
		}
	}(resultChan)
	var id int
	for {
		id++
		num := rand.Int()
		v := &job{
			Id:      id,
			RandNum: num,
		}
		jobChan <- v
	}
}

func createPoll(num int, jobChan <-chan *job, ResultChan chan<- *Result) {
	for i:=0; i<num; i++ {
		go func(jobchan <-chan *job, resultChan chan<- *Result) {
			for v := range jobChan {
				r := &Result{&job{
					Id:v.Id,
					RandNum:v.RandNum,
				}}
				resultChan <- r
			}
		}(jobChan, ResultChan)
	}
}


