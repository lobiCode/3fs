package main

import (
	"github.com/lobiCode/3fs/mylib"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	arith := new(mylib.Arith)
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
