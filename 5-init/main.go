package main

import (
	//. "GolangStudy/5-init/lib1"		// 所有方法被注入到本包作用域中，调用时无需前缀
	_ "GolangStudy/5-init/lib1"      //引入不调用不报错
	mylib2 "GolangStudy/5-init/lib2" //别名
)

func main() {
	//lib1.Lib1Test()
	mylib2.Lib2Test()
	//Lib1Test()


}
