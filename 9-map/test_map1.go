package main

import "fmt"

func main() {
	//first style：
	//声明myMap1 is type of map. key is string, value is also string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is empty")
	}

	//before use map, need use 'make' to assign space to map
	myMap1 = make(map[string]string, 10)
	fmt.Println(len(myMap1))
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	fmt.Println(len(myMap1))
	fmt.Println(myMap1)


//	second style
	myMap2 := make(map[int]string)
	fmt.Println(len(myMap2))
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"
	fmt.Println(myMap2)
	fmt.Println(len(myMap2))


	//	three style
	myMap3 := map[string]string{
		"one": "php",
		"two": "c++",
		"three": "python",
	}
	fmt.Println(myMap3)
}
