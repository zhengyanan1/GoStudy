package main

import "fmt"

/*
四种变量声明方式
 */

// 方法1、2、3可以声明全局变量
var gA int = 300
var gB = 500
//方法4只能够再函数内部声明使用
//gC := 600

func main() {
//	1。声明一个变量,默认值是0
	var a int
	fmt.Println(a)
	fmt.Printf("type of a = %T\n", a)

//	2.声明一个变量，初始化一个值
	var b int = 100
	fmt.Println(b)
	fmt.Printf("type of b = %T\n", b)


	//	3.初始化的时候，省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println(c)
	fmt.Printf("type of c = %T\n", c)

//	4. 省去var关键字，直接自动匹配
	e := 100
	fmt.Println(e)
	fmt.Printf("type of e is %T\n", e)


//	声明多个变量
	var xx, yy int = 100, 200
	fmt.Println(xx, ":", yy)

	var kk, ll = 100, "abcd"
	fmt.Println(kk, ":", ll)

	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println(vv, ":", jj)
}