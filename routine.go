package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started.")

	//チャンネル
	finished := make(chan bool)

	go func() {
		log.Print("sleep1 started")
		time.Sleep(1 * time.Second)
		log.Print("sleep1 finished")
		finished <- true
	}()

	go func() {
		log.Print("sleep2 started")
		time.Sleep(2 * time.Second)
		log.Print("sleep2 finished")
		finished <- true
	}()

	go func() {
		log.Print("sleep3 started")
		time.Sleep(3 * time.Second)
		log.Print("sleep3 finished")
		finished <- true
	}()

	for i := 1; i <= 3; i++ {
		<-finished
	}

	log.Print("all finished!!")
}
