package utils

import (
	"Go_ChatServer/chat_client/models"
	"Go_ChatServer/chat_common/message"
	"encoding/json"
	"log"
)

// messageSend sends the message string 'input' to connection 'c'
func MessageSend(player *models.Player, msg message.IMessage) {
	if player == nil || player.Conn == nil {
		log.Println("Connection is closed.")
		return
	}
	msgData, err := json.Marshal(msg)
	if err != nil {
		log.Println("Encode message error: ", err)
		return
	}
	netMsg := message.NetMessage{
		MsgName: msg.MsgName(),
		Data:    string(msgData),
	}
	data, err := json.Marshal(netMsg)
	if err != nil {
		log.Println("Encode net message error: ", err)
		return
	}
	_, _ = (*player.Conn).Write(data)
}
