package main

import "fmt"

func main() {
//	define a channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行")
		c <- 666  // send 666 to c
	}()

	fmt.Println(" 111")
	num :=  <- c
	fmt.Println(" 222")

	fmt.Println("num =", num)
	fmt.Println("main goroutine finished")
}