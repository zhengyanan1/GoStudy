package main

import "fmt"

//如果类名首字母大写，表示其他包也能访问
type Hero struct {
	// 如果类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name string
	Ad int
	Level int
	name2 string
}
/* version1
func (this Hero) Show(){
	fmt.Println("hero = ", this)
}

func (this Hero) GetName(){
	fmt.Println("Name = ", this.Name)
}

func (this Hero) SetName(newName string){
	// this 是调用该方法的对象的一个副本
	this.Name = newName
}
*/

func (this *Hero) Show(){
	fmt.Println("hero = ", this)
}

func (this *Hero) GetName() string{
	return this.Name
}

func (this *Hero) SetName(newName string){
	// this 是调用该方法的对象的一个副本
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{
		Name: "zhang3",
		Ad: 100,
		Level:1}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}