package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	rpcClient, _ := rpc.Dial("tcp", ":8080")

	msisdn := "number"
	var reply string

	rpcClient.Call("Arith.ParseMsisdn", msisdn, &reply)

	fmt.Println(reply)

}
