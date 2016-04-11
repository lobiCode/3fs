package main

import (
	"encoding/json"
	"fmt"
	"net/rpc/jsonrpc"
)

type MsisdnFormat struct {
	Cc  string // country dialing code
	Ndc string // mno identifier (national destination code)
	Sn  string // sucriber code
	Ci  string // country indentiifier
}

func main() {

	rpcClient, _ := jsonrpc.Dial("tcp", ":1234")

	msisdn := "+38640111111"

	var reply MsisdnFormat

	err := rpcClient.Call("Arith.ParseMsisdn", msisdn, &reply)

	if err != nil {
		fmt.Println(err)
	}

	j, _ := json.Marshal(reply)

	fmt.Println(string(j))

}
