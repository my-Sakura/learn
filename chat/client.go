package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1000")
	if err != nil {
		fmt.Println("dialErr: ", err)
		os.Exit(0)
	}

	sendMessage(conn)
}

func sendMessage(conn net.Conn) {
	userName := conn.LocalAddr().String()
	var msg string

	go receive(conn)

	for {
		fmt.Scanln(&msg)

		if msg == "quit" {
			break
		}
		_, err := conn.Write([]byte(userName + ":" + msg))
		if err != nil {
			fmt.Println("messageSendErr: ", err)
			conn.Close()
			break
		}
	}
}

func receive(conn net.Conn) {
	msg := make([]byte, 1024)

	for {
		_, err := conn.Read(msg)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(msg))
	}
}
