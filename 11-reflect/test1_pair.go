package main

import "fmt"

func main() {
	var a string = "aceld"
//	pair<type:string, value:"aceld> 这个pair会一直传递下去


	var allType interface{}
	allType = a

	value, ok :=allType.(string)
	if ok {
		fmt.Println(value)
	}
}
