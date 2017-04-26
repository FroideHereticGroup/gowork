package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("goroutine started")

	//チャンネルの作成
	messages := make(chan bool)

	go func() {
		fmt.Println("goroutine1 started")
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine1 finished")
		messages <- true
	}()

	go func() {
		fmt.Println("goroutine2 started")
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine2 finished")
		messages <- true
	}()

	go func() {
		fmt.Println("goroutine3 started")
		time.Sleep(3 * time.Second)
		fmt.Println("goroutine3 finished")
		messages <- true
	}()

	for i := 0; i < 3; i++ {
		<-messages
	}

	fmt.Println("All finished!!!")
}
