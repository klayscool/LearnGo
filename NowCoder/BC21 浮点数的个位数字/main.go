package main

import (
	"fmt"
	"strings"
)

var f float64

func Input() float64 {
	fmt.Scanln(&f)
	//fmt.Println(f)
	return f
}

func Output(f float64) {
	m := fmt.Sprint(f)
	//fmt.Println(m)
	n := strings.Split(m,".")[0]
	//fmt.Println(n)
	lenn := strings.Count(n,"") -1
	n = n[lenn-1:lenn]
	fmt.Println(n)
}

func main()  {
	Input()
	Output(f)
}