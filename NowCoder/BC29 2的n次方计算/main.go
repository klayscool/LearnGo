package main

import (
	"fmt"
	"math"
)

var n float64

func main() {
	fmt.Scanln(&n)
	res := math.Pow(2, n)
	fmt.Println(int(res))
}
