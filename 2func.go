package main

import "fmt"

//関数の呼び出しだとコンパイルが若干遅くなってしまう
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world");
	fmt.Println(a, b);

  //fmt.Println(a)だとエラーが出てしまう
}
