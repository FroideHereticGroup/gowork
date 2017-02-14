package main

import "fmt"

var num int = 0

func main() {
	var a int = 1
	var b string = "juho"
	//関数内であれば「:=」で暗黙の型変換が行われる
	c := 4
	d := "anchan"
	//var e string = [1,2,3,4];
	fmt.Println(a, b, c, d, e)
}
