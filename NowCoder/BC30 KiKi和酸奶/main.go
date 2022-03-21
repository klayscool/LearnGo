package main

import (
	"fmt"
	"math"
)

var n, h, m int
var drink float64

func main() {
	fmt.Scanln(&n, &h, &m)
	drink = float64(m) / float64(h)
	res := math.Ceil(drink)
	res2 := n - int(res)
	fmt.Println(res2)
}
