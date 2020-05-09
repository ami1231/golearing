package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getFlagName()
	//透過flag.parse才會把command line value讀入
	flag.Parse()
	fmt.Println(name)

}

//回傳＊string
func getFlagName() *string {
	return flag.String("name", "everyone", "name to start")
}
