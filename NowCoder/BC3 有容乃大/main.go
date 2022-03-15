package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var flag bool
	var n1 int8 = 10
	var name string = "小白"

	fmt.Printf("flag的字节大小是：",unsafe.Sizeof(flag))
	fmt.Println(" ")
	fmt.Printf("n1的字节大小是：",unsafe.Sizeof(n1))
	fmt.Println(' ')
	fmt.Printf("name的字节大小是：",unsafe.Sizeof(name))
}