package main

import (
	"fmt"
	"time"
)

func main()  {
	c := make(chan int, 3) // channl with buffer

	//fmt.Println(len(c), ":", cap(c))

	go func() {
		defer fmt.Println("child goroutine finished")

		for i := 0; i < 4; i++{
			c <- i
			fmt.Println("child goroutine is running, send data",i, " size:", len(c), ":", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++{
		num := <- c
		fmt.Println("num=",num)
	}

	fmt.Println("main finished")
}
