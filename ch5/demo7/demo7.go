package main

import (
	"container/list"
	"fmt"
)

//list結構
func main() {

	//list內部為element,push進去之後會轉為element

	list := list.New()
	for i := 0; i < 10; i++ {
		list.PushBack(i)
	}

	//透過Front取出element,需要.Value才可取得值
	// 因為element結構為{
	// 	Value interface{}
	// }
	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	//輸出第一個元素
	fmt.Println(list.Front().Value)

	//輸出最後一個元素
	fmt.Println(list.Back().Value)

	//list 當成queue使用
	//ring 保存固定數量的
	//heap可以用来排序。游戏编程中是一种高效的定时器实现方案

	//go container有list ring heap需要學習

}
