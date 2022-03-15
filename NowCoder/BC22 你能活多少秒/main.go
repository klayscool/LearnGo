package main

import (
	"fmt"
	"math"
)

var age int

func main()  {
	sec := int(3.156 * math.Pow10(7))
	fmt.Scanln(&age)
	res := age * sec
	fmt.Println(res)
}