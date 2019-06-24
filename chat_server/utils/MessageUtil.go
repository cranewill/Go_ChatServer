package utils

import (
	"fmt"

	connect "Go_ChatServer/chat_server/connect"
)

type MessageSendTask struct {
	Target int64
	Msg string
}

var SendChan chan MessageSendTask

// TellPlayer executes message sending to specific player
func TellPlayer(playerId int64, msg string) {
	SendChan <- MessageSendTask{Target:playerId, Msg:msg}
}

// Send is a goroutine fetching messages from SendChan to send to player
func Send(SendChan chan MessageSendTask) {
	for {
		sendTask := <-SendChan
		fmt.Printf("send message task: %v\n", sendTask)
		con, exist := connect.Pool.Conns[sendTask.Target]
		if !exist {
			fmt.Println("Cannot find player connection: ", sendTask.Target)
			return
		}

		_, err := con.Write([]byte(sendTask.Msg))
		if err != nil {
			fmt.Println("Send message failed: ", err)
		}
	}
}
