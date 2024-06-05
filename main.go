package main

import (
	"chainup.com/node-exchange/routers"
	"fmt"
)

func main() {
	fmt.Println("node中转启动了")
	if err := routers.Init().Run("0.0.0.0:10001"); err != nil {
		fmt.Println(err)
	}
}
