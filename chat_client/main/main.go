package main

import (
	"Go_ChatServer/chat_client/manager"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// messageSend sends the message string 'input' to connection 'c'
func messageSend(c net.Conn, input string) {
	if c == nil {
		log.Println("Connection is closed.")
		return
	}
	if len(input) == 0 {
		return
	}
	c.Write([]byte(input))
	//log.Println("send msg:", input)
}

// messageListen is a goroutine always listen message from connection 'c' and deal with it
func messageListen(c net.Conn) {
	for {
		//log.Println("Listening...")

		buf := make([]byte, 4096)
		cnt, err := c.Read(buf)
		if err != nil {
			log.Printf("Fail to read data, %s\n", err)
			return
		}

		msgMap := make(map[string]interface{})
		err = json.Unmarshal(buf[0:cnt], &msgMap)
		if err != nil {
			log.Println("Decode message error: ", err)
			return
		}

		cmd := fmt.Sprintf("%v", msgMap["Id"])
		deal(cmd, buf[0:cnt])
	}
}

// deal assigns messages to right logic according to their 'id'
func deal(id string, msg []byte) {
	switch strings.ToLower(id) {
	case "rooms":
		manager.ShowRooms(msg)
	case "chat":
		manager.ShowChat(msg)
	case "enter":
		manager.EnterResult(msg)
	case "create":
		manager.CreateResult(msg)
	case "quit":
		manager.QuitResult(msg)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Printf("Fail to connect, %s\n", err)
		return
	}
	defer conn.Close()

	go messageListen(conn)

	log.Println("Input your player ID (int64): ")
	var idStr string
	fmt.Scanln(&idStr)

	fmt.Println("Hello ", idStr, "! You can use the command below: ")
	fmt.Println("'show': show all the chat rooms.")
	fmt.Println("'create': create a chat room.")
	fmt.Println("'enter': enter a existing chat room.")
	fmt.Println("'quit': quit a chat room or the client.")
MAIN:
	for {
		time.Sleep(time.Second)
		log.Println("Input your command: ")

		var input string
		fmt.Scanln(&input)

		var cmd string
	CMD:
		switch input {
		case "quit":
			log.Println("Input 'a' to quit the client, and others to quit the chat room: ")
			var quitType string
			fmt.Scanln(&quitType)
			cmd = "{\"id\":\"quit\",\"playerId\":" + idStr + ",\"quitType\":\"" + quitType + "\"}"
			if quitType == "a" {
				messageSend(conn, cmd)
				break MAIN
			}
			messageSend(conn, cmd)
		case "show":
			cmd = "{\"id\":\"show\",\"playerId\":" + idStr + "}"
			messageSend(conn, cmd)
		case "create":
			cmd = "{\"id\":\"create\",\"playerId\":" + idStr + "}"
			messageSend(conn, cmd)
		case "enter":
			log.Println("Input the room ID which you want to enter: ")
			var enterId string
			fmt.Scanln(&enterId)
			cmd = "{\"id\":\"enter\",\"playerId\":" + idStr + ",\"roomId\":" + enterId + "}"
			messageSend(conn, cmd)
		case "chat":
			log.Println("Input the message('quit' to back to menu): ")
			for {
				var msg string
				var inputReader *bufio.Reader
				inputReader = bufio.NewReader(os.Stdin)
				msg, err = inputReader.ReadString('\n')
				msg = strings.Replace(msg, "\n", "", -1)
				if msg == "quit" {
					break CMD
				}
				cmd = "{\"id\":\"chat\",\"playerId\":" + idStr + ",\"content\":\"" + msg + "\"}"
				fmt.Println("【You】: ", msg, "\n")
				messageSend(conn, cmd)
			}
		}
	}
}
