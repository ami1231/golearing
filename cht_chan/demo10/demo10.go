package main

import "fmt"

//通道高级用
//透过型别宣告,限制channel在每个地方的使用条件
func main() {

	var channel = make(chan int, 2)
	sendItem(channel)
	getItem(channel)
}

//透过宣告型别 chan <- 表示这个通道在这个function 只可send
func sendItem(ch chan<- int) {
	ch <- 3
}

//透过宣告型别 <-chan 表示这个通道在这个function 只可get
func getItem(ch <-chan int) {
	fmt.Println(<-ch)
}
