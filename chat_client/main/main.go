package main

import (
	"Go_ChatServer/chat_client/manager"
	"Go_ChatServer/chat_client/message"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// messageSend sends the message string 'input' to connection 'c'
func messageSend(c net.Conn, msg message.IMessage) {
	if c == nil {
		log.Println("Connection is closed.")
		return
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Println("Encode message error: ", err)
		return
	}
	c.Write(data)
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
		panic("Fail to connect: " + err.Error())
		return
	}
	defer conn.Close()

	go messageListen(conn)

	fmt.Println("Input your player ID (int64): ")
	var idStr string
	_, _ = fmt.Scanln(&idStr)

	playerId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Invalid player ID: ", err)
	}

	fmt.Println(fmt.Sprintf(`Hello %s! You can use the command below: 
'show': show all the chat rooms.
'create': create a chat room.
'enter': enter a existing chat room.
'quit': quit a chat room or the client.`, idStr))

	var quitProcess bool
	for {
		if quitProcess {
			break
		}
		log.Println("Input your command: ")

		var input string
		_, _ = fmt.Scanln(&input)

		var cmd string
		switch input {
		case "quit":
			log.Println("Input 'a' to quit the client, and others to quit the chat room: ")
			var quitType string
			_, _ = fmt.Scanln(&quitType)
			if quitType == "a" {
				messageSend(conn, message.ReqQuitMessage{
					Id:       "quit",
					PlayerId: playerId,
					QuitType: quitType,
				})
				quitProcess = true
				break
			}
			messageSend(conn, cmd)
		case "show":
			messageSend(conn, message.ReqShowRoomsMessage{
				Id:       "show",
				PlayerId: playerId,
			})
		case "create":
			messageSend(conn, message.ReqCreateChatRoomMessage{
				Id:       "create",
				PlayerId: playerId,
			})
		case "enter":
			log.Println("Input the room ID which you want to enter: ")
			var enterId string
			_, _ = fmt.Scanln(&enterId)
			roomId, err := strconv.ParseInt(enterId, 10, 64)
			if err != nil {
				log.Println("Invalid room ID: ", err)
				break
			}
			messageSend(conn, message.ReqEnterChatRoomMessage{
				Id:       "enter",
				PlayerId: playerId,
				RoomId:   roomId,
			})
		case "chat":
			log.Println("Input the message('quit' to back to menu): ")
			for {
				var msg string
				var inputReader *bufio.Reader
				inputReader = bufio.NewReader(os.Stdin)
				msg, err = inputReader.ReadString('\n')
				msg = strings.Replace(msg, "\n", "", -1)
				if msg == "quit" {
					break
				}
				fmt.Println("【You】: ", msg, "\n")
				messageSend(conn, message.ReqChatMessage{
					Id:       "enter",
					PlayerId: playerId,
					Content:  msg,
				})
			}
		}
	}
}
