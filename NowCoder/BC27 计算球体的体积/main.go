package main

import (
	"fmt"
	"strconv"
)

var r float64
var pai = 3.1415926

func main()  {
	fmt.Scanln(&r)
	v := pai * r * r * r * 4 / 3
	v, _ = strconv.ParseFloat(fmt.Sprintf("%.3f",v),64)
	fmt.Println(v)
}