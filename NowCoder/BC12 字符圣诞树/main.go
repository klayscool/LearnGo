package main

import "fmt"

var (
	num int
)


func Tree(num int)  {
	for i := 1 ; i <= 5; i ++ {
		for z := 5 - i ; z >= 0 ; z -- {
			for j := 1 ; j <= i; j ++ {
				fmt.Printf("%d",num)
			}
			fmt.Printf(" ")
		}
		fmt.Println("\n")
	}
}

func main()  {
	fmt.Println("输入一个整数：")
	fmt.Scanln(&num)
	Tree(num)
}