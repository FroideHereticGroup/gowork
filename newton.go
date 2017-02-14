package main

import "fmt"

func square(z float64) float64 {
	return z * z
}

func Sqrt(x float64) float64 {
	var y [12]float64
	y[0] = 4.0
	for i := 0; i < 11; i++ {
		y[i+1] = (square(y[i]) - x) / (2 * y[i])
	}
	return y[11]
}

func main() {
	fmt.Println(Sqrt(2))
}
