package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	//取得命令列輸入
	flag.StringVar(&name, "name", "everyone", " name of user")
	//取得命令列輸入
	var name1 = flag.String("name1", "everyone", " name of user")
	// parse取得value
	flag.Parse()
	fmt.Println(name)

	//go語言中,若回傳為＊abc表示為該值的指針,記憶體位址,這時變數的值為記憶體地址,要拿到真正的value 需要加上＊
	fmt.Println(*name1)

}
