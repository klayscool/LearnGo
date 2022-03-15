package main

import (
	"fmt"
	"strings"
)

var s1, s2 string

func main() {
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
