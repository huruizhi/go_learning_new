package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	var b [128]byte
	defer conn.Close()
	for {
		n, err := conn.Read(b[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b[:n]))
	}
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
