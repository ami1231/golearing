package main

import "fmt"

//channel ,併發安全
func main() {

	intChannel := make(chan int)
	intChannel <- 0
	intChannel <- 1
	intChannel <- 2

	value := <-intChannel
	fmt.Println(value)

}
