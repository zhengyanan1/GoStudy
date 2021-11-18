package main

import "fmt"

//announce one kind of new data type, another name for int
type myint int

// Book define a struct
type Book struct {
	title string
	auth string
}

func changeBook(book Book){
	//transfer a copy of book
	book.auth = "666"
}

func changeBook2(book *Book){
//	transfer a pointer
	book.auth = "777"
}

func main() {
	//var a myint = 10
	//fmt.Println(a)
	//fmt.Printf("%T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"
	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
