package main

import (
	"errors"
	"fmt"
)

var (
	a,b,c int
)

func prin(a,b,c int) {
	if a < 0 || a > 100 {
		fmt.Println(errors.New("0≤a≤100"))
	} else {
		if b < 0 || b > 100 {
			fmt.Println(errors.New("0≤b≤100"))
		} else {
			if c < 0 || c > 100 {
				fmt.Println(errors.New("0≤c≤100"))
			} else {
				fmt.Printf("score1=%d,score2=%d,score3=%d",a,b,c)
			}
		}
	}
}

func main()  {
	//fmt.Println("Please enter your score:")
	fmt.Scanln(&a,&b,&c)
	prin(a,b,c)
}