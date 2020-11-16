package main

import (
	"context"
	"strconv"
	"fmt"
)


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	i := 0
	for {
		generateJob(ctx, "itask" + strconv.Itoa(i))
		i++
		if i> 10000 {
			break
		}
	}

	// time.Sleep(10*time.Second)
}

//生成 job
func generateJob(parent context.Context, value string) {
	ch := make(chan int, 0)
	ctx, cancel := context.WithCancel(parent)
	go doTask(ch, ctx, value)
	<-ch
	cancel()
}


//执行任务
func doTask(ch chan<- int, ctx context.Context, job string) {
	select {
	case <-ctx.Done():
		fmt.Println("job is closed", job)
		return
	default:
		fmt.Println(job, "is running")
		ch <- 1
	}
}




