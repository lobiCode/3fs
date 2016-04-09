package main

import (
	"net"
	"net/rpc"

	"./mylib"
)

func main() {

	arith := new(methods.Arith)
	rpc.Register(arith)

	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
	// TODO err

	listener, _ := net.ListenTCP("tcp", tcpAddr)
	// TODO err

	for {
		conn, _ := listener.Accept()
		// TODO err
		rpc.ServeConn(conn)
	}

}
