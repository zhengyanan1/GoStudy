package main

import (
	"fmt"
	"time"
)

func main() {
	//用go创建一个匿名函数
	//go func() {
	//	defer fmt.Println("A.defer")
	//
	//	func(){
	//		defer  fmt.Println("B.defer")
	//
	//		//exit current goroutine
	//		runtime.Goexit()
	//
	//		fmt.Println("B")
	//	}()
	//
	//	fmt.Println("A")
	//}()

	go func(a int, b int) bool{
		fmt.Println("a=,", a, ";b=", b)
		return true
	}(10, 20)

//	dead cycle
	for{
		time.Sleep(time.Second)
	}
}
