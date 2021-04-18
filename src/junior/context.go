package main

import (
	"go-alltime/src/junior/channels"
	ct "go-alltime/src/junior/context"
	"log"
	"context"
	"strconv"
)

//run processor
func runP() {
	//通道
	channels.ChanNoCache()
	channels.Pipeline()
	// channels.ChanWithCache()
	// channels.SievePrime()
	contextMultiWithValueJob()
}

//多协程管理 vithValue 测试
func contextMultiWithValueJob() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	i := 0
	for {
		ct.GenerateJobWithValue(ctx, "itask" + strconv.Itoa(i))
		i++
		if i> 10 {
			break
		}
	}
	// time.Sleep(10*time.Second)
}


func main() {
	log.Println("welcome go alltime !")
	runP()
}