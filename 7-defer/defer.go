package main

import "fmt"

func func1(){
	fmt.Println("A")
}

func func2(){
	fmt.Println("BB")
}

func func3(){
	fmt.Println("CCC")
}

func main(){
	//写入defer 关键字,栈——后进先出
	//defer fmt.Println("main end1")
	//defer fmt.Println("main end2")
	//
	//fmt.Println("main::hello go 1")
	//fmt.Println("main::hello go 2")


	//知识点1：defer的执行顺序
	defer func1()
	defer func2()
	defer func3()
}