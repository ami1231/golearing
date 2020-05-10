package main

import "fmt"

var container = []string{"zero", "one", "two"}

//如何判斷一個變數的型別
//透過斷言表達式 interface{}(container).([]string)
func main() {

	//區域變數覆蓋包級變數
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	fmt.Printf("the emelent is %q.\n", container[1])

	//空介面轉型物件,介面可判斷類型 .(類型)
	//assertType 若不正確,value = nil
	//若正確,則value會為正確的值
	//型別判斷false
	value, ok := interface{}(container).([]string)
	fmt.Println(value)
	fmt.Println(ok)

	//型別判斷true
	value1, ok1 := interface{}(container).(map[int]string)
	fmt.Println(value1)
	fmt.Println(ok1)

	//型別判斷t中間擺入表達式
	var value2, ok2 = interface{}(1 + 2).(int)
	if ok2 {
		fmt.Println(value2)
	}

}
