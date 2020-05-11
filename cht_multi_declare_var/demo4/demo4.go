package main

import "fmt"

var block = "123"

/**
在golang中不同的區塊重複聲明變數
區塊內的會覆蓋區塊外的值,在java也有這種作法
*/
func main() {
	fmt.Println(block)
	var block = "223"
	{
		block := 323
		fmt.Println(block)
	}
	fmt.Println(block)
}
