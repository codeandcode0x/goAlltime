package main

import (
	"go-alltime/src/junior/channels"
	"log"
)


func main() {
	log.Println("welcome go alltime !")
	runP()
}

func runP() {
	channels.ChanNoCache()
	channels.Pipeline()
	// channels.ChanWithCache()
	channels.SievePrime()
}