package main

import "fmt"

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	dog := Dog{name: "111"}
	//複製一個值到pet變數上
	var pet Pet = dog
	//為指針方法,因此會變更dog的value
	dog.SetName("222")
	//pet本身的name沒改變
	fmt.Println(pet)
	//dog本身的name為222
	fmt.Println(dog)

}
