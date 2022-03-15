package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	a int
	b int
)

func Input() (int,int) {
	//fmt.Println("一行，包括两个整数a和b，依次为被除数和除数（不为零），中间用空格隔开：")
	fmt.Scanln(&a,&b)
	if a <= 0 || a >= 10000 || b <= 0 || b >= 10000 {
		err := errors.New("a or b is error")
		fmt.Println(err)
		os.Exit(8)
	}
	return a,b
}

func Output(a,b int)  {
	sum1 := a / b
	sum2 := a % b
	fmt.Println(sum1,sum2)
}

func main()  {
	Input()
	Output(a,b)
}