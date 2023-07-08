package handler

import (
	"Go_ChatServer/chat_common/message"
	manager "Go_ChatServer/chat_server/connect"
	"Go_ChatServer/chat_server/models"
	"Go_ChatServer/chat_server/utils"
	"encoding/json"
	"log"
)

type AuthHandler struct {
}

func (handler AuthHandler) Handle(player *models.Player, msgMap []byte) {
	msg := message.ReqAuthMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		log.Println("Decode message error: ", err)
		return
	}
	Auth(player, msg.PlayerId)
}

// Auth check if a new player can login
func Auth(player *models.Player, playerId string) {
	msg := message.ResAuthMessage{
		Result: true,
	}
	if player.Id != "" {
		log.Println("repeat auth: ", playerId)
		msg.Result = false
	}
	_, ok := manager.Pool.GetPlayer(playerId)
	if ok {
		log.Println("player is already login: ", playerId)
		msg.Result = false
	}
	player.Id = playerId
	manager.Pool.SavePlayer(player)
	utils.Tell(playerId, msg)
}
