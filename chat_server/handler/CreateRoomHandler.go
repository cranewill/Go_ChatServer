package handler

import (
	"fmt"
	"encoding/json"

	manager "chat_server/manager"
	message "chat_server/message"
)

type CreateRoomHandler struct {

}

func (handler CreateRoomHandler)Deal(msgMap []byte) {
	// playerId := int64(msg["playerId"])
	msg := message.ReqCreateChatRoomMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		fmt.Println("Decode message error: ", err)
		return
	}
	manager.Pool.ChatManager.CreateChatRoom(msg.PlayerId)
}


