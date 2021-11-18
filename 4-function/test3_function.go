package main

import "fmt"

func foo1(a string, b int) int{
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}
//返回多个返回值，匿名
func foo2(a string, b int)(int, int){
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	return 666, 777
}
//返回多个返回值，有行参名称的
func foo3(a string , b int)(r1 int, r2 int){
	fmt.Println("_______foo3_________")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//r1 r2属于foo3的形参，初始化默认值是0
	//r1,r2作用域空间是foo3 整个函数体的{}空间
	fmt.Println("r1 =", r1)
	fmt.Println("r2 =", r2)

	r1 = 2000
	r2 = 1000

	return
}

func foo4(a string, b int)(r1, r2 int)  {
	fmt.Println("_______foo4_________")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	r1 = 2000
	r2 = 1000

	return
}

func main(){
	c := foo1("anc", 555)
	fmt.Println("c=", c)


	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1=",ret1)
	fmt.Println("ret2=",ret2)


	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1=",ret1)
	fmt.Println("ret2=",ret2)

	ret1, ret2 = foo4("foo4", 333)
	fmt.Println("ret1=",ret1)
	fmt.Println("ret2=",ret2)
}
