package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conns := make(map[string]net.Conn)
	messageCh := make(chan []byte, 10)
	var userName string
	l, err := net.Listen("tcp", ":1000")
	checkErr((err))

	go sendMessage(&conns, messageCh)

	for {
		conn, err := l.Accept()
		checkErr(err)
		userName = conn.RemoteAddr().String()
		conns[userName] = conn

		go Handle(conn, userName, messageCh)
	}
}

func Handle(conn net.Conn, userName string, messageCh chan []byte) {
	msg := make([]byte, 1024)
	for {
		_, err := conn.Read(msg)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s login out\n", userName)
				conn.Close()

				break
			} else {
				fmt.Println("messageReadErr :", err)
				conn.Close()
				break
			}
		}
		messageCh <- msg
	}
}

func sendMessage(conns *map[string]net.Conn, messageCh chan []byte) {
	for {
		msg := <-messageCh
		for key, c := range *conns {
			_, err := c.Write(msg)
			if err != nil {
				delete(*conns, key)
			}
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
	return
}
