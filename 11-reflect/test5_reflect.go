package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (p *User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", p)
}

func main(){
	user := User{1, "Harden", 18}

	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}){
//	获取type 和value
	inputType := reflect.TypeOf(input)
	fmt.Println("input type:", inputType.Name())
	inputValue :=reflect.ValueOf(input)
	fmt.Println("inputValue is ", inputValue)

//	通过type 获取里面的字段
	for i:=0; i < inputType.NumField(); i++{
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", filed.Name, filed.Type, value)
	}

//通过type 获取里面的方法
	fmt.Println(inputType.NumMethod(), "******")
	for i:= 0; i < inputType.NumMethod(); i++{
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}