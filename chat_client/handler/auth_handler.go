package handler

import (
	"Go_ChatServer/chat_client/models"
	"Go_ChatServer/chat_client/utils"
	"Go_ChatServer/chat_common/message"
	"encoding/json"
	"fmt"
	"log"
)

var AuthResultChan chan int

func HandleAuth(player *models.Player) bool {
	utils.MessageSend(player, message.ReqAuthMessage{
		PlayerId: player.Id,
	})
	return true
}

// AuthResult deals with the auth result
func AuthResult(msgData []byte) {
	var msg message.ResAuthMessage
	err := json.Unmarshal(msgData, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	if msg.Result {
		AuthResultChan <- 1
	} else {
		fmt.Println("Auth failed!")
	}
}
