package main

import "fmt"

type Human struct {
	name string
	sex string
}

func (p *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (p *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

// SuperMan  ----------------------------
type SuperMan struct {
	Human  //SuperMan类继承了Human类的方法

	level int
}

// Eat 重新定义父类的方法Eat
func (p *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// Fly 子类的新方法
func (p *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (p *SuperMan) Print(){
	fmt.Println("SuperMan:", p)
}

func main()  {
	h := Human{
		name: "zhang3",
		sex: "woman",
	}

	h.Eat()
	h.Walk()

	//s := SuperMan{Human{ "li4", "man"}, 6}
	var s SuperMan
	s.name = "li4"
	s.sex = "man"
	s.level = 90

	s.Walk()
	s.Eat()
	s.Fly()

	s.Print()
}