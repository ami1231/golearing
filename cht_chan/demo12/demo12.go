package main

import (
	"fmt"
	"strconv"
)

//透过range 迭带 channel
// 没宣告容量,所以为非缓冲cahnnel,
// sendCake roroutine 阻塞
// 等到receive接收才会放入下一个
func main() {
	cs := make(chan string)
	go sendCake(cs, "", 10)
	go receiveCakeAndPack(cs)

	csStrbry := make(chan string)
	csChoco := make(chan string)

	go sendCake(csStrbry, "strbry", 10)
	go sendCake(csChoco, "choco", 10)

	go receiveCakeAndPack1(csChoco, csStrbry)

	//	time.Sleep(3 * 1e9)
}

func sendCake(cs chan<- string, name string, count int) {

	for i := 0; i < count; i++ {
		fmt.Println(i)
		cakeName := "Cake " + name + strconv.Itoa(i)
		cs <- cakeName
	}
	close(cs)
}

func receiveCakeAndPack(cs <-chan string) {
	for s := range cs {
		fmt.Println(s)
	}
}

func receiveCakeAndPack1(strbryCs <-chan string, chocoCs <-chan string) {
	strbryCsClose, chocoCsClose := false, false

	for {
		if strbryCsClose && chocoCsClose {
			return
		}
		fmt.Println("waiting new cacke")
		select {
		case cakeName, strbryOk := <-strbryCs:
			if strbryOk {
				fmt.Println("StrbyCake is new " + cakeName)
			} else {
				strbryCsClose = true
				fmt.Println("strby cahnnel is close")
			}
		case cakeName, chocoOk := <-chocoCs:
			if chocoOk {
				fmt.Println("ChocoCake is new " + cakeName)
			} else {
				strbryCsClose = true
				fmt.Println("Choco cahnnel is close")
			}
		}
	}

}
