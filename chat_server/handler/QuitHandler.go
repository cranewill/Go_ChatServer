package handler

import (
	"encoding/json"
	"log"

	"Go_ChatServer/chat_server/manager"
	"Go_ChatServer/chat_server/message"
)

type QuitHandler struct {

}

func (handler QuitHandler)Deal(msgMap []byte) {
	msg := message.ReqQuitMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		log.Println("Decode message error: ", err)
		return
	}
	manager.Pool.ChatManager.Quit(msg.PlayerId, msg.QuitType)
}