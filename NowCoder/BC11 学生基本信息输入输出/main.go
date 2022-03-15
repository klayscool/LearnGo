package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	num int64
	score1,score2,score3 float64
)

func InputInfo() (int64,float64,float64,float64) {
	//fmt.Println("依次输入一个学生的学号，以及3科（C语言，数学，英语）成绩，在屏幕上输出该学生的学号，3科成绩（注：输出成绩时需进行四舍五入且保留2位小数）。\n\n数据范围：学号满足 1 \\le n \\le 20000000 \\1≤n≤20000000  ，各科成绩使用百分制，且不可能出现负数\n输入描述：\n学号以及3科成绩，学号和成绩之间用英文分号隔开，成绩之间用英文逗号隔开。\n如：\n17140216;80.845,90.55,100.00")
	fmt.Scanf("%d;%f,%f,%f",&num,&score1,&score2,&score3)
	return num,score1,score2,score3
}

func StuInfo(num int64,score1 float64,score2 float64,score3 float64) {
	if num < 0 {
		fmt.Println(errors.New("Student ID must be greater than 0"))
		os.Exit(20)
	}
	if score1 < 0 || score1 > 100 {
		fmt.Println(errors.New("score1 is error"))
		os.Exit(11)
	}
	if score2 < 0 || score2 > 100 {
		fmt.Println(errors.New("score2 is error"))
		os.Exit(12)
	}
	if score3 < 0 || score3 > 100 {
		fmt.Println(errors.New("score3 is error"))
		os.Exit(13)
	}
	score1, _ = strconv.ParseFloat(fmt.Sprintf("%.2f",score1),64)
	score2, _ = strconv.ParseFloat(fmt.Sprintf("%.2f",score2),64)
	score3, _ = strconv.ParseFloat(fmt.Sprintf("%.2f",score3),64)
	fmt.Printf("The each subject score of No. %d is %f, %f, %f.",num,score1,score2,score3)
}

func main()  {
	InputInfo()
	StuInfo(num,score1,score2,score3)
}