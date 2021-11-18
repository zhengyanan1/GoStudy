package main

import "fmt"

func printMap(targetMap map[string]string){
	//targetMap 是一个引用传递
	for key, value := range targetMap{
		fmt.Println(key, ":", value)
	}
}

func changeValue(targetMap map[string]string){
	targetMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//visit
	printMap(cityMap)
	fmt.Println("))))))))))))))))))))))))))))))))))))))))))))")
//	delete
	delete(cityMap, "China")
	printMap(cityMap)

//	modify
	cityMap["USA"] = "DC"
	fmt.Println("-------------------------------------------")
	printMap(cityMap)

	fmt.Println("-------------------------------------------")
	changeValue(cityMap)
	printMap(cityMap)

}
