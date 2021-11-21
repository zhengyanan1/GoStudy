package main

import "fmt"

func fibonacii(c, quit chan int){
	x, y := 1, 1

	for{
		fmt.Println("select....")
		select {
			case c <- x:
				fmt.Println("write to c: data =", x)
				y = x + y
				x = y - x
			case <- quit:
				fmt.Println("quit")
				return
		}
	}
}

func main()  {
	c := make(chan int)
	quit := make(chan int)


	go func() {
		for i:=0; i<6; i++{
			fmt.Println(" try read form c")
			fmt.Println("read:", <-c)
		}

		quit <- 0
	}()


	fibonacii(c, quit)
}
