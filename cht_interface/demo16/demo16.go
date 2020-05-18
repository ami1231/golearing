package main

import "fmt"

type pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string
}

// 本案例可以见到,Dog包含了 Category两个function
//但是若为＊Dog 其包含了三个function
//因此＊Dog指针类型实作 pet介面,但是Dog类型没有实作pet介面
func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

//go interface
// 在golang中实作介面,不需要implemets 只需要把setName,name,category三个method都实作,系统会自动认定属于pet介面
func main() {
	var dog pet = &Dog{name: "111"}
	//  这行不会过,因为是＊Dog实作了pet介面,并不是dog,一个struct若method为值的接收者,那值就只有这些function,但引用类型的接收者会自动包含值的function
	//	var dog pet = Dog{name: "111"}
	fmt.Println(dog.Name())
	//因为setName为指针方法,因此set后原变数会一起连动
	dog.SetName("2222")
	fmt.Println(dog.Name())
}
