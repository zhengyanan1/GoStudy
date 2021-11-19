package main

import "fmt"

//本质 是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类
type Cat struct {
	color string // 猫的颜色
}

func (p *Cat) Sleep(){
	fmt.Println("Cat is Sleep")
}

func (p *Cat) GetColor() string{
	return p.color
}

func (p *Cat) GetType() string{
	return "Cat"
}


type Dog struct {
	color string
}

func (p *Dog) Sleep(){
	fmt.Println("Dog is Sleep")
}

func (p *Dog) GetColor() string{
	return p.color
}

func (p *Dog) GetType() string{
	return "Dog"
}

func showAnimal(animal Animal){
	animal.Sleep()  // 多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	//var animal Animal  //接口类型，父类指针
	//animal = &Cat{"Green"}
	//animal.Sleep()
	//
	//animal = &Dog{"Yellow"}
	//animal.Sleep()

	cat := Cat{
		color: "Green",
	}
	dog := Dog{
		color: "Yellow",
	}

	showAnimal(&cat)
	showAnimal(&dog)
}
