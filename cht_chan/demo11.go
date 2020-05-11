package main

import (
	"fmt"
	"math/rand"
)

//channel select
func main() {

	//三個通道array
	intChannels := [3]chan int{make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1)}
	index := rand.Intn(3)
	intChannels[index] <- 999
	select {
	case <-intChannels[0]:
		fmt.Println("channel 1 is select")
	case <-intChannels[1]:
		fmt.Println("channel 2 is select")
	case <-intChannels[2]:
		fmt.Println("channel 3 is select")
	default:
		fmt.Println("no channel be select")
	}
}
