package main

import "fmt"

func myFunc(arg interface{}){
	fmt.Println("myFunc is called...")
	fmt.Println(arg)


//	interface{} 如何区分此时引用的底层数据类型到底是什么
	value, ok := arg.(string)
	if ok{
		fmt.Println("arg is string value=", value)
		fmt.Printf("%T\n", value)
	} else {
		fmt.Println("arg is not string")
	}
}

type Book struct {
	auth string
}

func main(){
	book := Book{
		auth: "harden",
	}

	myFunc(book)
	myFunc(77)
	myFunc("golang")
}
