package main

import (
	"fmt"
	"strconv"
)

func main() {
	hex := "0xC40C5253"
	val := hex[2:]
	fmt.Println(val)

	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		panic(err)
	}

	n2 := uint32(n)
	fmt.Print(n2)
}