package main

import (
	"Go_ChatServer/chat_client/handler"
	"Go_ChatServer/chat_client/models"
	"Go_ChatServer/chat_common/message"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic("Fail to connect: " + err.Error())
		return
	}
	defer conn.Close()

	// start a goroutine to listen response message
	go messageListen(conn)

	fmt.Println("Input your player name: ")
	var playerId string
	_, _ = fmt.Scanln(&playerId)

	// auth
	player := new(models.Player)
	player.Id = playerId
	player.Conn = &conn
	handler.HandleAuth(player)
	handler.AuthResultChan = make(chan int, 1)

	go mainLogic(player, handler.AuthResultChan)

	signChan := make(chan int, 1)
	select {
	case <-signChan:
		break
	}
	//for {
	//	fmt.Println("Input !quit to close this process ")
	//	var sign string
	//	_, _ = fmt.Scanln(&sign)
	//	if sign == "!quit" {
	//		break
	//	}
	//}

}

// mainLogic blocks till receive the auth result, if auth success, then run the client function
func mainLogic(player *models.Player, authChan chan int) {
	var authSuccess bool
	for {
		select {
		case <-authChan:
			authSuccess = true
			break
		}
		if authSuccess {
			break
		}
	}

	fmt.Println(fmt.Sprintf(`--- Hello %s! You can use the command below: 
'show': show all the chat rooms.
'create': create a chat room.
'enter': enter a existing chat room.
'quit': quit a chat room or the client.`, player.Id))

	var quitProcess bool
	for {
		if quitProcess {
			break
		}
		time.Sleep(time.Second)
		fmt.Println("--- Input your command: ")

		var cmd string
		_, _ = fmt.Scanln(&cmd)
		handle, ok := handler.ReqRoutes[cmd]
		if !ok {
			continue
		}
		_ = handle(player)
	}
}

// messageListen is a goroutine always listen message from chat server and deal with it
func messageListen(c net.Conn) {
	for {
		buf := make([]byte, 4096)
		length, err := c.Read(buf)
		if err != nil {
			log.Printf("Fail to read data, %s\n", err)
			return
		}

		netMsg := message.NetMessage{}
		err = json.Unmarshal(buf[:length], &netMsg)
		if err != nil {
			log.Println("Decode message error: ", err)
			return
		}

		routes(netMsg.MsgName, []byte(netMsg.Data))
	}
}

// routes assigns messages to right logic according to their name
func routes(msgName string, msg []byte) {
	handle, ok := handler.ResRoutes[msgName]
	if !ok {
		return
	}
	handle(msg)
}
