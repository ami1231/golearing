package main

import "fmt"

//动物分类
type animalCategory struct {
	kingdom string //界
	phylum  string //门
	class   string //纲
	order   string //目
	family  string //科
	genus   string //属
	species string //种
}

type animal struct {
	scientificName string //学名
	animalCategory        //嵌入一个类型
}

type cat struct {
	animal
	name string
}

//宣告 AniamlCategory类型得string方法
func (ac animalCategory) Sring() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", ac.kingdom, ac.phylum, ac.class, ac.order, ac.family,
		ac.genus, ac.species)
}

func (ac animalCategory) changeSpecies() animalCategory {
	ac.species = "123"
	return ac
}

func (a animal) category() string {
	return a.animalCategory.Sring()
}

// struts 结构体篇
// function 本身可以匿名,可以被当成value传递
// 但是宣告在结构体内的方法,不可匿名,不可以被当成值传递
// 重点是他归属在某一个类型下
func main() {

	//初始化结构体方法1
	category := animalCategory{species: "cat"} //结构体初始化直接返回值
	fmt.Printf("the animal category : %s\n", category)
	category2 := category.changeSpecies()

	fmt.Println(category)
	fmt.Println(category2)
	fmt.Println("==============================")
	//初始化结构体方法2
	category1 := new(animalCategory) // new function 返回是一个指针
	category1.species = "dog"
	fmt.Printf("the animal category : %s\n", category1)
	category3 := category1.changeSpecies()

	fmt.Println(category1)
	fmt.Println(category3)
	fmt.Println("==============================")
	//创建animal
	animal := animal{}
	//嵌入字段使用方式
	animal.animalCategory.species = "Lion"
	animal.scientificName = "John"

	fmt.Printf("The animal: %s\n", animal)

	fmt.Println("==============================")
	fmt.Println(animal.category())

	//由此可证明,当透过组合structs,animal继承了animalGateory得method,属性,只有当两者同名时,才有需要透过a.b.属性去设置
	animal.class = "334"
	animal.animalCategory = animal.changeSpecies()

	fmt.Println("==============================")
	fmt.Println(animal.category())

	//多重struct

}

// struct本身的method 是具有overwrite得特性
func (cat1 cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)", cat1.scientificName, cat1.animal.animalCategory, cat1.name)
}

//在goalng中 并不像java具有所谓的继承的特性跟设计,但透过嵌入字段｜组合的方式,让外层物件拥有内层物件的属性及方法
//因此当我在本身物件组合了3个物件,该物件就拥有了这三个物件的特性,不像oop语言是要强制继承具有侵入性的
//因为透过组合模式 物件本身的定义还是主体,嵌入的只是附属功能,但当我透过继承时,物件本身主体变成继承的类型
//而golang在方法上本身支持overwrite,但不支持overload
