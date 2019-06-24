package handler

import (
	"fmt"
	"encoding/json"

	manager "Go_ChatServer/chat_server/manager"
	message "Go_ChatServer/chat_server/message"
)

type ChatHandler struct {

}

func (handler ChatHandler)Deal(msgMap []byte) {
	msg := message.ReqChatMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		fmt.Println("Decode message error: ", err)
		return
	}
	manager.Pool.ChatManager.Chat(msg.PlayerId, msg.Content)
}