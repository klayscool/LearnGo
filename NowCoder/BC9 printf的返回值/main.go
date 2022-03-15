package main

import "fmt"


func prin() (n int,error error){
	n,err := fmt.Printf("Hello world!")
	return n,err
}

func main()  {
	code,error := prin()
	fmt.Println(code,error)
}