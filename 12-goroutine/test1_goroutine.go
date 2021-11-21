package main

import (
	"fmt"
	"time"
)

func newTask(){
	i := 0
	for{
		i++
		fmt.Printf("new Goroutine: i= %d\n", i)
		time.Sleep(time.Second)
	}
}

func main(){
	go newTask()

	fmt.Println("main goroutine exit")
	//i := 0
	//for{
	//	i++
	//	fmt.Printf("main goroutine i=%d\n", i)
	//	time.Sleep(time.Second)
	//}
}
