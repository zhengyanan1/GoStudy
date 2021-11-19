package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string			`info:"name" doc:"我的名字"`
	Sex string			`info:"sex"`
}

func findTag(str interface{})  {
	type1 := reflect.TypeOf(str)
	fmt.Println("type1:", type1)

	fmt.Println(type1.NumField())


	t := reflect.TypeOf(str).Elem()
	for i:= 0; i < t.NumField(); i++{
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", taginfo, "; doc:", tagdoc)
	}
}

func main(){
	re := resume{"Harden", "man"}
	findTag(re)
}
