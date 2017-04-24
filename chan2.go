package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "AAA"
	messages <- "BBB"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
