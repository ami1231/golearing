package main

import "fmt"

//goroutine2
//不透過sleep,讓主線程等到所有GO RUNTUINE做完
// func main() {

// 	num := 100

// 	sign := make(chan struct{}, num)

// 	for i := 0; i < num; i++ {
// 		go func(index int) {
// 			fmt.Println(index)
// 			sign <- struct{}{}
// 		}(i)
// 	}

// 	for i := 0; i < num; i++ {
// 		value := <-sign
// 		fmt.Println(value)
// 	}
// }

func main() {

	num := 100

	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func(index int) {
			fmt.Println(index)
			sign <- struct{}{}
		}(i)
	}

	for i := 0; i < num; i++ {
		value := <-sign
		fmt.Println(value)
	}
}
