package main

import (
	"log"
	"time"
)

func mySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	log.Println("Подождём секундочку...")
	mySleep(time.Second)
	log.Println("Подождали.")
}
