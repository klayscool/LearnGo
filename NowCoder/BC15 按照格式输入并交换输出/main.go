package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

var (
	a int
	b int
	num int
	min = int(-math.Pow(2,31))
	max = int(math.Pow(2,31)-1)
)

func Input() (int,int) {
	//fmt.Println("输入只有一行，按照格式输入两个整数，范围，中间用“,”分隔：")
	fmt.Scanf("a=%d,b=%d",&a,&b)
	if a < min || a > max || b < min || b > max {
		err := errors.New("a or b is error")
		fmt.Println(err)
		os.Exit(99)
	}
	return a,b
}

func Output(a,b int)  {
	num = a
	a = b
	b = num
	fmt.Printf("a=%d,b=%d",a,b)
}

func main()  {
	Input()
	Output(a,b)
}