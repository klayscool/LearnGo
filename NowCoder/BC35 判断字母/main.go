package main

import (
	"fmt"
	"unicode"
)

var s rune

func main() {
	fmt.Scanln(&s)
	res := unicode.IsLetter(s)
	fmt.Println(res)
}
