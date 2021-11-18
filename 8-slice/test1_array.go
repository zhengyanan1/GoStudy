package main

import "fmt"

func printArray(myArray [4]int){ //传参是值拷贝
	for index, value := range myArray{
		fmt.Println(index, ": = ", value)
	}
	myArray[0] = 99
	fmt.Println("in func:", myArray[0])
}

//func main(){
//	//固定长度的数组
//	var myArray1 [10]int
//	myArray2 := [10]int{1,3,4,5}
//	myArray3 := [4]int{11, 22, 33,44}
//	//for i := 0; i < len(myArray1);i++ {
//	//	fmt.Println(myArray1[i])
//	//}
//	//
//	//for index, value := range myArray2{
//	//	fmt.Println(index, ": = ", value)
//	//}
//
////	查看数组的数据类型
//	fmt.Printf("type1  %T\n", myArray1)
//	fmt.Printf("type2  %T\n", myArray2)
//	fmt.Printf("type3  %T\n", myArray3)
//
//
//	printArray(myArray3)
//	fmt.Println(myArray3[0])
//}
