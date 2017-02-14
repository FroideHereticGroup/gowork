package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on ")
	//switch文はbreakがなくても大丈夫。caseの後にプログラムを抜ける
	switch os := runtime.GOOS; os { // switch文でもステートメント変数をかける
	case "darwin":
		fmt.Print("OS X.\n")
	case "linux":
		fmt.Print("Linux.\n")
	default:
		fmt.Print(os + ".\n")
	}
}
