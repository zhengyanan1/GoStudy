package main

import "fmt"

func main() {
	s := []int{1,2,3}

	//[0, 2)
	s1 := s[0:2]

	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s)


//	copy 可以将底层数组一起拷贝
	s2 := make([]int, 2)
	//将s中的值，依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2, "===")
}
