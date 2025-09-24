package main

import (
	"ctpgo/ctpgo"
	"fmt"
)

func main() {
	pq := ctpgo.InitQuote("./log_quote111/", false, false)
	fmt.Println(pq.GetApiVersion())

	pt := ctpgo.InitTrade("./log_trade111/")
	fmt.Println(pt.GetApiVersion())

}
