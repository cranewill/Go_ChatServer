package handler

import (
	"encoding/json"
	"log"

	"Go_ChatServer/chat_server/manager"
	"Go_ChatServer/chat_server/message"
)

type ChatHandler struct {
}

func (handler ChatHandler) Deal(msgMap []byte) {
	msg := message.ReqChatMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		log.Println("Decode message error: ", err)
		return
	}
	manager.Pool.ChatManager.Chat(msg.PlayerId, msg.Content)
}
