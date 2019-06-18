package main

import (
	"fmt"
	"net"
	"strings"
)

func connHandler(c net.Conn, input string) {
	if c == nil {
		fmt.Println("Connection is closed.")
		return
	}

	defer c.Close()

	buf := make([]byte, 4096)
	for {
		input = strings.TrimSpace(input)
		c.Write([]byte(input))
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			continue
		}
		fmt.Print(string(buf[0:cnt]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:7777")
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}
	for {
		var input string
		fmt.Scanln(&input)
		if input == "quit" {
			break;
		}
		connHandler(conn, input)
	}
}
