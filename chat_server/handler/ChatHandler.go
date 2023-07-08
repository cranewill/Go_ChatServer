package handler

import (
	"Go_ChatServer/chat_common/message"
	"Go_ChatServer/chat_server/db"
	"Go_ChatServer/chat_server/models"
	"Go_ChatServer/chat_server/utils"
	"encoding/json"
	"log"

	"Go_ChatServer/chat_server/manager"
)

type ChatHandler struct {
}

func (handler ChatHandler) Handle(player *models.Player, msgMap []byte) {
	msg := message.ReqChatMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		log.Println("Decode message error: ", err)
		return
	}
	Chat(player.Id, msg.Content)
}

// Chat receives player's words and broadcasts to every player in the chat room
func Chat(playerId string, content string) {
	roomId, err := db.Redisdb.Get(db.PlayerRoomKey(playerId)).Result()
	if err != nil {
		if !db.IsNil(err) {
			log.Println("Get player roomId error: ", err)
			return
		} else {
			log.Println("Player is not in a chat room")
			return
		}
	}

	room := manager.Pool.ChatManager.ChatRooms[roomId]

	msg := message.ResChatMessage{Id: "chat", PlayerId: playerId, Content: content}
	for _, memberId := range room.Members {
		if memberId == playerId {
			continue
		}
		utils.Tell(memberId, msg)
	}
}
