package main

import (
  "fmt"
)

func add(x, y int) int {
  return x + y;
}

func main() {
  fmt.Printf("%.4d\n", add(40, 29));
}
