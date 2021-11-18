package main

import "fmt"

func main() {
	//slice1 是一个切片，并且初始化，默认值是1，2，3, 长度len是3
	//slice1 := []int{1,2,3}

	//声明是一个切片，但并没有给slice1 分配空间
	var slice1 []int
	//slice1 = make([]int, 3)
	//slice1[0] = 100

	//声明的同时分配空间
	//var slice1 []int = make([]int, 5)

	//声明的同时分配空间，通过:= 自动推倒出slice 是一个切片
	//slice1 := make([]int ,3)


	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)


//	判断slice 是否为0
	if slice1 == nil{
		fmt.Println("空切片")
	} else {
		fmt.Println(" slice 有空间")
	}
}
