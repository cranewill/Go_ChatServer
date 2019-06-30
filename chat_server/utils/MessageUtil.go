package utils

import (
	"Go_ChatServer/chat_server/messageout"
	"encoding/json"
	"log"

	connect "Go_ChatServer/chat_server/connect"
)

type MessageSendTask struct {
	Target int64
	Msg    string
}

type MessageTask struct {
	Target int64
	Msg    messageout.Message
}

var StrSendChan chan MessageSendTask
var MsgSendChan chan MessageTask

// Tell executes STRING message sending to specific player
func TellPlayer(playerId int64, msg string) {
	StrSendChan <- MessageSendTask{Target: playerId, Msg: msg}
}

// Send is a goroutine fetching messages from StrSendChan to send to player
func Send(SendChan chan MessageSendTask) {
	for {
		sendTask := <-SendChan
		log.Printf("send message task: %v\n", sendTask)
		con, exist := connect.Pool.Conns[sendTask.Target]
		if !exist {
			log.Println("Cannot find player connection: ", sendTask.Target)
			return
		}

		_, err := con.Write([]byte(sendTask.Msg))
		if err != nil {
			log.Println("Send message failed: ", err)
		}
	}
}

// Tell sends Message type message to channel
func Tell(playerId int64, message messageout.Message) {
	MsgSendChan <- MessageTask{playerId, message}
}

// SendMsg is a goroutine worked with Tell() to fetch MessageTasks from MsgSendChan and send them to players
func SendMsg(msgChan chan MessageTask) {
	for {
		message := <-msgChan
		con, exist := connect.Pool.Conns[message.Target]
		if !exist {
			log.Println("Cannot find player connection: ", message.Target)
			return
		}

		msg, err := json.Marshal(message.Msg)
		if err != nil {
			log.Println("Json encode message failed: ", err)
			return
		}
		//log.Printf("Send message %s\n", msg)
		_, err = con.Write(msg)
		if err != nil {
			log.Println("Send message failed: ", err)
		}

	}
}
