package lib1

import "fmt"

//lib1提供的api
func Lib1Test() {
	fmt.Println("Lib1Test")
}

func lib1Test2(){
	fmt.Println("lib1Test2")
}

func init(){
	fmt.Println("lib1 init()...")
}