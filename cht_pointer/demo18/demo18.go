package main

import "fmt"

type Dog struct {
	name string
}

//Dog的指針方法
func (dog *Dog) setName(name string) {
	dog.name = name
}

func main() {

	var dog Dog

	fmt.Println(dog)
	fmt.Println(&dog)

}
