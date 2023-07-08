package handler

import (
	"Go_ChatServer/chat_client/models"
	"Go_ChatServer/chat_client/utils"
	"Go_ChatServer/chat_common/message"
	"encoding/json"
	"fmt"
	"log"
)

func HandleEnter(player *models.Player) bool {
	fmt.Println("--- Input the room ID which you want to enter: ")
	var roomId string
	_, _ = fmt.Scanln(&roomId)
	utils.MessageSend(player, message.ReqEnterChatRoomMessage{
		Id:     "enter",
		RoomId: roomId,
	})
	return true
}

// EnterResult returns the result of entering chat room
func EnterResult(msgData []byte) {
	var msg message.ResEnterResultMessage
	err := json.Unmarshal(msgData, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	fmt.Println(msg.Result)
}
