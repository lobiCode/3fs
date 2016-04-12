package main

import (
	"errors"
	"fmt"
	"github.com/lobiCode/3fs/mylib"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

var (
	rpcClient *rpc.Client
	addr      string = ":1234"
	method    string = "Arith.ParseMsisdn"
)

func init() {

	var err error
	rpcClient, err = jsonrpc.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
}

func TestSI(t *testing.T) {

	test := &mylib.MsisdnFormat{"386", "40", "111111", "SI"}
	err := HelpTest(test, "+38640111111")

	if err != nil {
		t.Error(err)
	}
}

func TestFR(t *testing.T) {

	test := &mylib.MsisdnFormat{"33", "1", "23456789", "FR"}
	err := HelpTest(test, "+33123456789")

	if err != nil {
		t.Error(err)
	}
}

func HelpTest(t *mylib.MsisdnFormat, msisdn string) error {

	var reply mylib.MsisdnFormat

	err := rpcClient.Call(method, msisdn, &reply)

	if err != nil {
		return err
	}

	if reply.Cc != t.Cc {
		return errors.New("wrong country dialing code")
	}

	if reply.Ndc != "" && reply.Ndc != t.Ndc {
		return errors.New("wrong mno identifier")
	}

	if reply.Sn != (t.Ndc+t.Sn) && reply.Sn != t.Sn {
		return errors.New("wrong sucriber identifier")
	}

	if reply.Ci != t.Ci {
		return errors.New("wrong country identifier")
	}

	return nil
}
