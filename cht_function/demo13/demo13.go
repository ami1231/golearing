package main

import (
	"errors"
	"fmt"
)

//将function定义成型别
type printer func(contents string) (n int, err error)

func printToStd(contents string) (byteNum int, err error) {
	return fmt.Println(contents)
}

func main() {

	//宣告一个printer变数
	var p printer
	//将printTostd指到p变数上,因为这个function的传入参数,回传参数类型与printer一样,因此是同型别
	p = printToStd
	//直接执行p变数
	var n, _ = p("1234")
	fmt.Println(n)

	//透过自定义function传入,改变calcute运行结果
	//因此calcute为一高阶函数
	fmt.Println(calcute(2, 10, func(x int, y int) int {
		return x + y
	}))

	fmt.Println(calcute(2, 10, func(x int, y int) int {
		return x - y
	}))

	fmt.Println(calcute(2, 10, func(x int, y int) int {
		return x * y
	}))

	//也可以透过function 传出function也是高阶函数表示是
	op := func(x, y int) int {
		return x + y*2
	}
	//传入function ,传出function
	var cal = genCalculator(op)
	value, err := cal(10, 20)
	if err == nil {
		fmt.Println(value)
	}

}

//高阶函数
//1.接收其他function传入 or
//2.把其他function返回

type opreate func(x, y int) int

func calcute(x, y int, opt opreate) (int, error) {
	if opt == nil {
		return 0, errors.New("opreate is nil")
	}
	return opt(x, y), nil
}

type calculateFunc func(x int, y int) (int, error)

func genCalculator(op opreate) calculateFunc {

	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid opreaction")
		}
		return op(x, y), nil
	}
}
