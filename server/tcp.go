package main

import (
	"fmt"
	"net"
	"os"
)

func startTcp() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Printf("init listener: %s", err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept connection: %s", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Remote client %s just connected\n", conn.RemoteAddr())
}
