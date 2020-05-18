package main

import (
	"fmt"
	"time"
)

//goroutine
func main() {
	fmt.Println("go routine start")
	//goroutine需要傳入參數方式
	//因為有傳入i,所以進行value copy
	for i := 0; i < 10; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}

	//無傳入i,所以當要執行goroutine時,i已經完成都是10
	//會時10的原因是,迴圈跑完時,i=10,才觸發迴圈結束條件,
	//每個goroutine要去抓取i的值,這時候已經是10
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("no param = %d\n", i)
		}()
	}

	time.Sleep(time.Millisecond * 1000)
}
