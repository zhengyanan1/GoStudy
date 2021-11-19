package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}


type Book struct {
}

func (p *Book) ReadBook()  {
	fmt.Println("read a book")
}

func (p *Book) WriteBook() {
	fmt.Println("write a book")
}

func main() {
	b := &Book{}

	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = r.(Writer)  //此处为什么可以转型成功，因为w r具体的type是一致的
	w.WriteBook()
}
