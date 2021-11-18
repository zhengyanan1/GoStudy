package main

import "fmt"

//const 定义枚举类型
const (
	//可以再const()添加一个关键字，每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING = 10 * iota   //iota = 0
	SHNAGHAI		//iota = 1
	SHENXHEN		//iota = 2
)

const(
	a, b = iota + 1, iota + 2	//iota = 0, a=1, b=2
	c, d						//iota = 1, c=2, d=3
	e, f						//iota = 2, e=3, d=4
	g, h = iota*2, iota * 3		//iota = 3, g=6, h=9
	i, k						//iota = 4, i=8, k=12
)

func main() {
	//常量
	const length int = 10

	fmt.Println("length = ", length)
	//length = 1000 // 常量不允许修改


	fmt.Println(BEIJING, SHNAGHAI, SHENXHEN)

	//iota 只能配合const() 一起使用，iota只有在const进行累加效果
	//var a int = iota

	const b = iota
	fmt.Println(b)
}
