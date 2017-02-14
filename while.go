package main

import "fmt"

func main() {
	sum := 1
	// goではwhile文はなくforで書く
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
