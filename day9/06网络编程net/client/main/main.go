package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		msg, err := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if err != nil {
			log.Fatal(err)
		}
		if msg == "exit" {
			break
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
	}

}
