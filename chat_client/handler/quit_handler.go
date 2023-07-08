package handler

import (
	"Go_ChatServer/chat_client/models"
	"Go_ChatServer/chat_client/utils"
	"Go_ChatServer/chat_common/message"
	"encoding/json"
	"fmt"
	"log"
)

func HandleQuit(player *models.Player) bool {
	utils.MessageSend(player, message.ReqQuitMessage{
		Id: "quit",
	})
	return true
}

// QuitResult returns the result of quiting chat room
func QuitResult(msgData []byte) {
	var msg message.ResQuitResultMessage
	err := json.Unmarshal(msgData, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}

	fmt.Println(msg.Result)
}
