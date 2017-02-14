//fizzbuzz

package main

import (
	"fmt"
	"strconv"
)

//関数の中はreturnが一つだけ戻ってくるように設計する
//中でfor文を使うときはswitch文などでreturnが一つになるように設計すること
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
	for i := 1; i <= 15; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
