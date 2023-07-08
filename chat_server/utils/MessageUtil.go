package utils

import (
	"Go_ChatServer/chat_common/message"
	"encoding/json"
	"log"

	connect "Go_ChatServer/chat_server/connect"
)

type MessageRawTask struct {
	Target string
	Msg    string
}

type MessageTask struct {
	Target string
	Msg    message.IMessage
}

var StrSendChan chan MessageRawTask
var MsgSendChan chan MessageTask

// Tell executes STRING message sending to specific player
func TellRaw(playerId string, msg string) {
	StrSendChan <- MessageRawTask{Target: playerId, Msg: msg}
}

// Send is a goroutine fetching messages from StrSendChan to send to player
func Send(SendChan chan MessageRawTask) {
	for {
		sendTask := <-SendChan
		log.Printf("send message task: %v\n", sendTask)

		target, exist := connect.Pool.GetPlayer(sendTask.Target)
		if !exist {
			log.Println("Cannot find player connection: ", sendTask.Target)
			return
		}

		_, err := (*target.Conn).Write([]byte(sendTask.Msg))
		if err != nil {
			log.Println("Send message failed: ", err)
		}
	}
}

// Tell sends Message type message to channel
func Tell(playerId string, message message.IMessage) {
	MsgSendChan <- MessageTask{playerId, message}
}

// SendMsg is a goroutine worked with Tell() to fetch MessageTasks from MsgSendChan and send them to players
func SendMsg(msgChan chan MessageTask) {
	for {
		msgData := <-msgChan
		target, exist := connect.Pool.GetPlayer(msgData.Target)
		if !exist {
			log.Println("Cannot find player connection: ", msgData.Target)
			return
		}
		netMsg := message.NetMessage{}
		msg, err := json.Marshal(msgData.Msg)
		if err != nil {
			log.Println("Json encode message failed: ", err)
			return
		}
		netMsg.MsgName = msgData.Msg.MsgName()
		netMsg.Data = string(msg)
		data, err := json.Marshal(netMsg)
		if err != nil {
			log.Println("Encode net message error: ", err)
			return
		}
		_, err = (*target.Conn).Write(data)
		if err != nil {
			log.Println("Send message failed: ", err)
		}
	}
}
