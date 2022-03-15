package main

import "fmt"

var time int

func main()  {
	fmt.Scanln(&time)
	hour := time / 3600
	min := time % 3600 / 60
	sec := time % 3600 % 60
	fmt.Println(hour,min,sec)
}