package main

import (
	"flag"
	"fmt"
)

func main() {
	//正常變量聲明,型別這時候宣告
	var name string
	//取得命令列輸入
	flag.StringVar(&name, "name", "everyone", " name of user")
	//取得命令列輸入
	//型別由編譯的時候go幫你確認
	var name1 = flag.String("name1", "everyone", " name of user")
	// parse取得value
	flag.Parse()
	fmt.Println(name)

	//go語言中,若回傳為＊abc表示為該值的指針,記憶體位址,這時變數的值為記憶體地址,要拿到真正的value 需要加上＊
	fmt.Println(*name1)

	fmt.Printf("start up name is %v\n", name)

	//短變量聲明,只能在function內使用
	abc := 20
	fmt.Print(abc)
}
