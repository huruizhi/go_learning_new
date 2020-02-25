package main

import (
	"go_learning_new/day9/08tcp粘包解决/proto"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 20; i++ { // 粘包
		msg := "hello this is a client message!"
		msgByte, err := proto.Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
		if msg == "exit" {
			break
		}
		_, err = conn.Write(msgByte)
		if err != nil {
			log.Fatal(err)
		}
	}

}
