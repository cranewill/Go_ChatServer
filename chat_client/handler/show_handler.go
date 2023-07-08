package handler

import (
	"Go_ChatServer/chat_client/models"
	"Go_ChatServer/chat_client/utils"
	"Go_ChatServer/chat_common/message"
	"encoding/json"
	"fmt"
	"log"
)

func HandleShow(player *models.Player) bool {
	utils.MessageSend(player, message.ReqShowRoomsMessage{
		Id: "show",
	})
	return true
}

// ShowRooms displays all the existing chat room info
func ShowRooms(msgData []byte) {
	var msg message.ResRoomsMessage
	err := json.Unmarshal(msgData, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	if len(msg.Rooms) == 0 {
		fmt.Println("There is no chat room now. You can create a new chat room with 'create' command")
		return
	}
	for _, room := range msg.Rooms {
		fmt.Println("id:", room.RoomId, "\tmembers:", room.Members)
	}
}
