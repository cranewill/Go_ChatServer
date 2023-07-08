package handler

import (
	"Go_ChatServer/chat_client/models"
	"Go_ChatServer/chat_client/utils"
	"Go_ChatServer/chat_common/message"
	"encoding/json"
	"fmt"
	"log"
)

func HandleCreate(player *models.Player) bool {
	utils.MessageSend(player, message.ReqCreateChatRoomMessage{
		Id: "create",
	})
	return true
}

// CreateResult returns the result of creating chat room
func CreateResult(msgData []byte) {
	var msg message.ResCreateResultMessage
	err := json.Unmarshal(msgData, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}

	fmt.Println(msg.Result)
}
