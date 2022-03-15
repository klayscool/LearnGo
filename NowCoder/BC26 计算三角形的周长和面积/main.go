package main

import (
	"fmt"
	"math"
)

var l1,l2,l3 float64

func main()  {
	fmt.Scanln(&l1,&l2,&l3)
	circumference := l1 + l2 + l3

	p := (l1 + l2 + l3) / 2.0
	area := math.Sqrt(p * (p - l1) * (p - l2) * (p - l3))

	fmt.Printf("circumference=%.2f area=%.2f",circumference,area)
}