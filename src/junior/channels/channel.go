package channels

import (
	"go-alltime/src/util/client/http"
	"log"

)

//chan 同步通道 (无缓存通道)
func ChanNoCache() {
	ch := make(chan int, 0)

	go func() {
		var sum int = 0
		for i :=0; i<10; i++ {
			sum = sum + i
		}
		ch <- sum
	}()
	//在计算sum和的goroutine没有执行完，把值赋给ch通道之前，
	//fmt.Println(<-ch)会一直等待
	log.Println(<-ch)

}

//chan 通道 (有缓存)
func ChanWithCache()  string {
	response := make(chan string, 3)

	go func() { response <- http.Request("https://godoc.org/google.golang.org/grpc") }()
	go func() { response <- http.Request("https://godoc.org/debug/gosym") }()
	go func() { response <- http.Request("https://godoc.org/context") }()

	//输出所有数据
	for i:=0 ; i< cap(response); i++ {
		log.Println(<-response)
		log.Println("----------", i)
	}

	//返回最快的获取到数据
	return <- response
}

//pipeline 管道
func Pipeline() {
	begin := make(chan int, 0)
	end := make(chan int, 0)

	go func() {
		begin <- 10
	}()


	go func() {
		num := <- begin
		end <- num
	}()

	log.Println(<-end)

}

//经典实例
//查找素数 - Sieve Prime 
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	count := 0
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
		count++
	}
}

// The prime sieve: Daisy-chain Filter processes.
func SievePrime() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 100; i++ {
		prime := <-ch
		log.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}






