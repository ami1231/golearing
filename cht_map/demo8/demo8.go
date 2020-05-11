package main

import "fmt"

//map數據結構
func main() {

	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	key := "two"

	value, ok := map1[key]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("value not found")
	}
}
