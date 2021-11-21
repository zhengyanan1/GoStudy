package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i:=0; i< 5; i++{
			c <- i
		}

		//"close" can close a channel
		close(c)
	}()

	//for {
	//	//ok=true, channel is open;false=ok, channel is closed
	//	if data, ok := <-c; ok{
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//}

	//can use range to get data form channel
	for data := range c{
		fmt.Println(data)
	}


	fmt.Println("Main Finished...")
}
