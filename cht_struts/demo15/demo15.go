package main

import "fmt"

type cat struct {
	name string
	age  int
}

//指針類型方法
//接收者為cat物件的指針,因此進入setName為這個cat的原本物件
func (cat *cat) setName(name string) {
	cat.name = name
}

//值方法,接收者為cat物件的副本,因此進入setName1時,這個cat為複製出來的副本,不是原本的cat
func (cat cat) setName1(name string) {
	cat.name = name
}

//指針型方法,與值方法
func main() {

	//此時的cat為值
	cat := cat{name: "222"}

	//method 為指針類型
	//因此呼叫此function時,根據指針拿到原本的cat物件
	//呼叫setName時,golang會自動把cat轉換成 &cat取得該屬性的指針,因此可以調用指針方法
	cat.setName("333")
	fmt.Println(cat)

	// method為value類型
	//呼叫此function時,修正的為複製的cat值,外層物件不受影響
	cat.setName1("444")
	fmt.Println(cat)
}
