package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) string {
	if num%15 == 0 {
		return "fizzbuzz"
	} else if num%3 == 0 {
		return "fizz"
	} else if num%5 == 0 {
		return "buzz"
	} else {
		return strconv.Itoa(num)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
