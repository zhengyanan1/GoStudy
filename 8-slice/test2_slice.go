package main

import "fmt"

func printArray2(myArray []int){
	//引用传递
	for _, value := range myArray{
		fmt.Println("value:",value)
	}
	myArray[0] = 1000
}

func main() {
	//动态数组
	myArray := []int{1, 2,3,4}

	fmt.Printf("%T\n", myArray)
	printArray2(myArray)

	for _, value := range myArray{
		fmt.Println("value:",value)
	}
}
