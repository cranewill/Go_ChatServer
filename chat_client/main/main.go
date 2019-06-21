package main

import (
	"fmt"
	"net"
	"strings"
)

func messageSend(c net.Conn, input string) {
	if c == nil {
		fmt.Println("Connection is closed.")
		return
	}

	input = strings.TrimSpace(input)
	c.Write([]byte(input))
}

func messageListen(c net.Conn) {
	for {
		fmt.Println("Listening...")

		buf := make([]byte, 4096)
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			continue
		}
		fmt.Println(string(buf[0:cnt]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}
	defer conn.Close()

	go messageListen(conn)
	
	for {
		var input string
		fmt.Scanln(&input)
		if input == "quit" {
			break
		}
		messageSend(conn, input)
	}
}
