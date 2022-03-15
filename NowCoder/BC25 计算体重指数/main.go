package main

import (
	"fmt"
	"strconv"
)

var weight,high int

func main()  {
	fmt.Scanln(&weight,&high)
	high2 := float64(high) / float64(100)
	res := float64(weight) / ( high2 * high2 )
	res, _ = strconv.ParseFloat(fmt.Sprintf("%.2f",res),64)
	fmt.Println(res)
}