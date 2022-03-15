package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
var n int

func Input() int {
	//fmt.Println("一行，输入一个整数n（1000 <= n <= 9999）：")
	fmt.Scanln(&n)
	if n < 1000 || n > 9999 {
		err := errors.New("n is error")
		fmt.Println(err)
		os.Exit(8)
	}
	return n
}

func Output(n int)  {
	n2 := strconv.Itoa(n)
	a := n2[0:1]
	b := n2[1:2]
	c := n2[2:3]
	d := n2[3:4]
	fmt.Printf("%s%s%s%s\n",d,c,b,a)

	//n3 := d + c + b + a
	//n4, _ := strconv.Atoi(n3)
	//fmt.Println(n4)

}

func main()  {
	Input()
	Output(n)
}