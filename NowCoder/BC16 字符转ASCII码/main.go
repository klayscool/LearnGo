package main

import "fmt"

var a string

func Input() string {
	//fmt.Println("一行，一个字符：")
	fmt.Scanln(&a)
	return a
}

func Output(a string)  {
	var c rune='a'
	fmt.Println(int(c))
}

func main()  {
	Input()
	Output(a)
}