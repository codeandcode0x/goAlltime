/**
* 多协程任务管理
*/

package context

import (
	"context"
	"log"
)

//生成 job with value
func GenerateJobWithValue(parent context.Context, value string) {
	ch := make(chan int, 0)
	valueKey := "value"
	ctx := context.WithValue(parent, valueKey, value)
	go doTaskWithValue(ch, ctx, valueKey)
	<-ch
}

//生成 job with cancel
func GenerateJob(parent context.Context, value string) {
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
		log.Println("job is closed", job)
		return
	default:
		log.Println(job, "is running")
		ch <- 1
	}
}

//带参数 context, 执行任务
func doTaskWithValue(ch chan<- int, ctx context.Context, valueKey string) {
    if v := ctx.Value(valueKey); v != nil {
        log.Println("found value:", v, ctx.Value(valueKey))
    }else{
    	log.Println("key not found:", valueKey)
    }
    ch <- 1
}




