package main

import (
	"ctpgo/ctpgo"
	"fmt"
)

func main() {
	pq := ctpgo.InitQuote()
	fmt.Println(pq.GetApiVersion())

	pt := ctpgo.InitTrade()
	fmt.Println(pt.GetApiVersion())

}
