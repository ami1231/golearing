package main

import "fmt"

//channel ,併發安全
func main() {

	// 數字位表示channel的容量
	// 容量=0時,為unbuffer channel ,> 0 為buffer channel
	intChannel := make(chan int, 3)
	var a1 = 0
	intChannel <- a1
	intChannel <- 1
	intChannel <- 2

	go func() {
		fmt.Println("channel is full up , block channel")
		//可以看出因为chan容量为3,所以这时候send阻塞
		intChannel <- 3
		fmt.Println("channel is empty , send channel")
	}()

	for i := 1; i < 10000000000; i++ {
		// if i > 9000000000 {
		// 	//因为目前这写法,会block main-goroutine,所以会导致main-goroutine爆出 deadlock
		// 	value := <-intChannel
		// 	fmt.Printf("value = %d\n", value)
		// }

		if i > 9000000000 && i < 9000000010 {
			//此写法不会抛错,因为是由其他的goroutine
			go func() {
				value := <-intChannel
				fmt.Printf("value = %d\n", value)
			}()
		}
	}
}
