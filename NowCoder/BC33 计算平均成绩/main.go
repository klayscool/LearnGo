package main

import "fmt"

var a, b, c, d, e int

func main() {
	fmt.Scanln(&a, &b, &c, &d, &e)
	sum := (float64(a) + float64(b) + float64(c) + float64(d) + float64(e)) / 5
	fmt.Printf("%.1f", sum)
}
