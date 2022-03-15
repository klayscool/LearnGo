package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	age string
	year int
	month int
	date int
)

func Input() (string,int,int,int) {
	//fmt.Println("请输入一个人的出生日期（包括年月日）：\n\n数据范围：年份满足1990≤y≤2015 ，月份满足1≤m≤12  ，日满足1≤d≤30 ")
	fmt.Scanln(&age)
	year, _ = strconv.Atoi(age[0:4])
	month, _ = strconv.Atoi(age[4:6])
	date, _ = strconv.Atoi(age[6:8])
	return age,year,month,date
}

func Output(age string, year int, month int, date int)  {

	if 1990 <= year && year <= 2015 {
		if 1 <= month && month <= 12 {
			if 1 <= date && date <= 31 {
				fmt.Printf("year=%s\n",age[0:4])
				fmt.Printf("month=%s\n",age[4:6])
				fmt.Printf("date=%s",age[6:8])
			} else {
				fmt.Println("date in error")
				os.Exit(99)
			}
		} else {
			fmt.Println("month is error")
			os.Exit(99)
		}
	} else {
		fmt.Println("year is error")
		os.Exit(99)
	}
}

func main()  {
	Input()
	Output(age,year,month,date)
}