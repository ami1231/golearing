package main

import "fmt"

//數組,切片 array slice
func main() {
	//宣告數組1
	var array1 [3]int
	fmt.Println(array1)
	array1[0] = 1

	//宣告數組2
	var array2 = [3]int{0, 1, 2}
	fmt.Println(array2)

	//宣告切片1
	var slice1 []int
	slice1 = append(slice1, 1)
	fmt.Println(slice1[0])

	//宣告切片2
	//先宣告array,在從array中提取切片
	var sliceArray [10]int
	var slice2 = sliceArray[0:5]
	fmt.Println(slice2)

	//宣告切片3
	//透過make製作切片
	var slice3 = make([]int, 3, 10)

	//切片長度 --3
	fmt.Println(len(slice3))
	//切片容量 --10
	fmt.Println(cap(slice3))

	fmt.Println(&slice3[0])
	fmt.Println(&slice3[1])

	//塞入20個元素
	// for i := 0; i < 100; i++ {
	// 	slice3 = append(slice3, i)
	// 	fmt.Printf("address of slice %p add of Arr \n", slice3)
	// }

	//切片長度為23
	fmt.Println(len(slice3))
	//切片容量為40,當切片到其容量上限時,會將其底層的array複製一倍容量,所以目前容量上限若為50,當數量要往51增加時,容量會擴充為100
	//當數量為100,繼續增加時,會擴充為200
	//但當原容量大於1024時,會用1.25倍增長
	fmt.Println(cap(slice3))
	fmt.Println(slice3)

	//golang 基礎類型,struct array = 值類型(傳值)
	//slice map chan function 為引用,存的為指標的value
}
