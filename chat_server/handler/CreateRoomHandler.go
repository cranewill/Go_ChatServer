package handler

import (
	"fmt"
	"encoding/json"

	manager "Go_ChatServer/chat_server/manager"
	message "Go_ChatServer/chat_server/message"
)

type CreateRoomHandler struct {

}

func (handler CreateRoomHandler)Deal(msgMap []byte) {
	msg := message.ReqCreateChatRoomMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		fmt.Println("Decode message error: ", err)
		return
	}
	manager.Pool.ChatManager.CreateChatRoom(msg.PlayerId)
}


