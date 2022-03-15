package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)
var score1,score2,score3 float64
func main()  {
	fmt.Scanln(&score1,&score2,&score3)
	sum := score1 + score2 + score3
	scoreAvg := sum / 3
	sum, _ = decimal.NewFromFloat(sum).Round(2).Float64()
	scoreAvg, _ = decimal.NewFromFloat(scoreAvg).Round(2).Float64()
	fmt.Println(sum,scoreAvg)
}