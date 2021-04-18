package main

import (
	"fmt"
	"time"
)

func main() {

    messages := make(chan string)

    go func() { 
    	messages <- "ping"
    	time.Sleep(5*time.Second)
    	fmt.Println("over")
    }()

    msg := <-messages
    fmt.Println(msg)
}