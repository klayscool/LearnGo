package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	a,b,a2,b2,sum int
	max = int(math.Pow(2,31) - 1)
)

func Input() (int,int) {
	//fmt.Println("一行，输入两个非负整数a和b，用一个空格分隔。（ 0 <= a,b <= 2^31-1 ）：")
	fmt.Scanln(&a,&b)

	if a < 0 || a > max || b < 0 || b > max {
		err := errors.New("a or b is error")
		fmt.Println(err)
		os.Exit(99)
	}

	a1 := strconv.Itoa(a)
	b1 := strconv.Itoa(b)

	La1 := strings.Count(a1,"") - 1
	Lb1 := strings.Count(b1,"") - 1

	if La1 > 2 {
		a1 = a1[len(a1)-2:len(a1)]
	}
	if Lb1 > 2 {
		b1 = b1[len(b1)-2:len(b1)]
	}

	a2, _ = strconv.Atoi(a1)
	//fmt.Printf("a2 is %d\n",a2)
	b2, _ = strconv.Atoi(b1)
	//fmt.Printf("b2 is %d\n",b2)

	return a2,b2
}

func Compute(a2,b2 int) int {
	sum = a2 + b2
	return sum
}

func Output(sum int)  {
	sum1 := strconv.Itoa(sum)
	Lasum := strings.Count(sum1,"") - 1
	if Lasum > 2 {
		sum1 = sum1[len(sum1)-2:len(sum1)]
	}
	sum, _ = strconv.Atoi(sum1)
	fmt.Println(sum)
}

func main()  {
	Input()
	Compute(a2,b2)
	Output(sum)
}