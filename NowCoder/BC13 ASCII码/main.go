package main

import "fmt"

var team = [...] int {73, 32, 99, 97, 110, 32, 100, 111, 32, 105, 116, 33}

func main() {
	a := ""
	for _, v := range team {
		a = a + string(v)
	}
	fmt.Println(a)
}
