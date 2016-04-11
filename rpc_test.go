package main

import (
	"errors"
	"fmt"
	"github.com/lobiCode/3fs/mylib"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

var rpcClient *rpc.Client
var addr string = ":1234"
var method string = "Arith.ParseMsisdn"

func init() {

	var err error
	rpcClient, err = jsonrpc.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
}

func TestSI(t *testing.T) {

	var reply mylib.MsisdnFormat

	test := "TestSI "
	msisdn := "+38640111111"
	cc := "386"
	ndc := ""
	ndc1 := "40"
	sn := "111111"
	sn1 := "04111111"
	ci := "SI"

	err := rpcClient.Call(method, msisdn, &reply)

	if err != nil {
		t.Error(err)
	}

	if reply.Cc != cc {
		t.Error(errors.New(test + "wrong country dialing code"))
	}

	if reply.Ndc != ndc && reply.Ndc != ndc1 {
		t.Error(errors.New(test + "wrong mno identifier"))
	}

	if reply.Sn != sn && reply.Sn != sn1 {
		t.Error(errors.New(test + "wrong sucriber identifier"))
	}

	if reply.Ci != ci {
		t.Error(errors.New(test + "wrong country identifier"))
	}
}
