package main

import (
	"fmt"
	"strings"
)

func main(){
	str := strings.Repeat(" ",5)
	pri := strings.Repeat("*",2)
	str2 := strings.Repeat("*",12)

	fmt.Println(str + pri + str)
	fmt.Println(str + pri + str)
	fmt.Println(str2)
	fmt.Println(str2)
	fmt.Println(str + pri + str)
	fmt.Println(str + pri + str)
}