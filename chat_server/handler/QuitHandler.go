package handler

import (
	"Go_ChatServer/chat_common/message"
	"Go_ChatServer/chat_server/db"
	"Go_ChatServer/chat_server/manager"
	"Go_ChatServer/chat_server/models"
	"Go_ChatServer/chat_server/utils"
	"encoding/json"
	"log"
)

type QuitHandler struct {
}

func (handler QuitHandler) Handle(player *models.Player, msgMap []byte) {
	msg := message.ReqQuitMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		log.Println("Decode message error: ", err)
		return
	}
	Quit(player.Id)
}

// Quit allows player to quit the chat room, and the same time removes its info from db
func Quit(playerId string) {
	msg := message.ResQuitResultMessage{
		Id:     "quit",
		Result: message.SUCCESS,
	}

	roomId, err := db.Redisdb.Get(db.PlayerRoomKey(playerId)).Result()
	if err != nil {
		if !db.IsNil(err) {
			log.Println("Get player roomId error: ", err)
			msg.Result = message.FAILED_TO_GET_ROOM_PLAYER_IN
			utils.Tell(playerId, msg)
			return
		} else { // db.Nil : player is not in any room
			log.Println("Player is not in a chat room")
			msg.Result = message.PLAYER_NOT_IN_CHAT_ROOM
			utils.Tell(playerId, msg)
			return
		}
	}
	mgr := manager.Pool.ChatManager
	room := mgr.ChatRooms[roomId]
	for i, memberId := range room.Members {
		if memberId == playerId {
			room.Members = append(room.Members[:i], room.Members[i+1:]...)
			break
		}
	}
	mgr.ChatRooms[roomId] = room

	_, err = db.Redisdb.Del(db.PlayerRoomKey(playerId)).Result()
	if err != nil {
		log.Println("Delete player-room info in db error: ", err)
		msg.Result = message.FAILED_TO_DEL_PLAYER_ROOM_INFO
		utils.Tell(playerId, msg)
		return
	}

	if len(room.Members) == 0 {
		delete(mgr.ChatRooms, roomId)
		log.Println("Delete chat room ", roomId)
	}

	utils.Tell(playerId, msg)
}
