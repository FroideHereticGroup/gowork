package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// ifもステートメントが書ける
	if v := math.Pow(x, n); v < lim {
		return v // vのスコープはif内のみ
	}
	return lim
}

func main() {
	fmt.Println(
		math.Pow(2, 3),
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
}
