package main

import (
	"./mylib"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	arith := new(methods.Arith)
	rpc.Register(arith)

	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":1234")
	// TODO err

	listener, _ := net.ListenTCP("tcp", tcpAddr)
	// TODO err

	for {

		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}

}
