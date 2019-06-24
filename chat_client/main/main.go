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
	if len(input) == 0 {
		return
	}
	input = strings.TrimSpace(input)
	c.Write([]byte(input))
	//fmt.Println("send msg: ",input)
}

func messageListen(c net.Conn) {
	for {
		//fmt.Println("Listening...")

		buf := make([]byte, 4096)
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			return
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

	fmt.Println("Input your player ID (int64): ")
	var idStr string
	fmt.Scanln(&idStr)
	//playerId, err := strconv.ParseInt(idStr,10,64)
	//if err != nil {
	//	fmt.Println("Input error: ", err)
	//	return
	//}
	MAIN:for {
		fmt.Println("Input your command: ")

		var input string
		fmt.Scanln(&input)
		//if input == "quit" {
		//	break
		//}
		var cmd string
		CMD:switch input {
		case "quit":
			break MAIN
		case "show":
			cmd = "{\"id\":\"show\",\"playerId\":" + idStr + "}"
		case "create":
			cmd = "{\"id\":\"create\",\"playerId\":" + idStr + "}"
		case "enter":
			fmt.Println("Input the room ID which you want to enter: ")
			var enterId string
			fmt.Scanln(&enterId)
			cmd = "{\"id\":\"enter\",\"playerId\":" + idStr + ",\"roomId\":" + enterId + "}"
		case "chat":
			fmt.Println("Input the message('quit' to back to menu): ")
			for {
				var msg string
				fmt.Scanln(&msg)
				if msg == "quit" {
					break CMD
				}
				cmd = "{\"id\":\"chat\",\"playerId\":" + idStr + ",\"content\":\"" + msg + "\"}"
				messageSend(conn, cmd)
			}
		}
		messageSend(conn, cmd)
	}
}
