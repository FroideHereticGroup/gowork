//並行処理

package main

import (
	"fmt"
	"time"
)

func repeat(name string, count int) {
	for i := 1; i <= count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name, fmt.Sprintf("%d回目", i))
	}
}

func main() {
	go repeat("笹本\t", 2) //2秒かかるgoroutine
	repeat("あん\t", 4)
	go repeat("椛島\t", 5) //5秒かかるgoroutine
	go repeat("山口\t", 3) //3秒かかるgoroutine
	//	go repeat("所崎\t", 4) //4秒かかるgoroutine
	//	repeat("あん\t", 6) //mainというgoroutineが終わるのに6秒かかる
}
