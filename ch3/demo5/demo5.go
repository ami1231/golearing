package main

import "fmt"

var container = []string{"zero", "one", "two"}

//如何判斷一個變數的型別
//透過斷言表達式 interface{}(container).([]string)
func main() {

	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	fmt.Printf("the emelent is %q.\n", container[1])
	//空介面轉型物件,介面可判斷類型 .(類型)
	_, ok := interface{}(container).([]string)

	if ok {
		fmt.Println("container is slice")
	} else {
		fmt.Println("container is map")
	}
}
