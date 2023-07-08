package handler

import (
	"Go_ChatServer/chat_common/message"
	"Go_ChatServer/chat_server/db"
	"Go_ChatServer/chat_server/models"
	"Go_ChatServer/chat_server/utils"
	"encoding/json"
	"fmt"
	"log"

	manager "Go_ChatServer/chat_server/manager"
)

type EnterRoomHandler struct {
}

func (handler EnterRoomHandler) Handle(player *models.Player, msgMap []byte) {
	msg := message.ReqEnterChatRoomMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		fmt.Println("Decode message error: ", err)
		return
	}
	EnterChatRoom(player.Id, msg.RoomId)
}

// EnterChatRoom allows player to enter the specific chat room
func EnterChatRoom(playerId, roomId string) {
	msg := message.ResEnterResultMessage{
		Id:     "enter",
		Result: message.SUCCESS,
	}

	playerRoomId, err := db.Redisdb.Get(db.PlayerRoomKey(playerId)).Result()
	if playerRoomId != "" {
		log.Println("Player is in a chat room now, please quit first: ", playerRoomId)
		msg.Result = message.PLAYER_ALREADY_IN_CHAT_ROOM
		utils.Tell(playerId, msg)
		return
	} else if !db.IsNil(err) {
		log.Println("Get playerRoomId error: ", err)
		msg.Result = message.FAILED_TO_GET_ROOM_PLAYER_IN
		utils.Tell(playerId, msg)
		return
	}
	room, exist := manager.Pool.ChatManager.ChatRooms[roomId]
	if !exist {
		log.Println("Room(", roomId, ") not exist!")
		msg.Result = message.ROOM_NOT_EXIST
		utils.Tell(playerId, msg)
		return
	}

	room.Members = append(room.Members, playerId)
	manager.Pool.ChatManager.ChatRooms[roomId] = room

	// db save this player and chat room
	err = db.Redisdb.Set(db.PlayerRoomKey(playerId), roomId, 0).Err()
	if err != nil {
		log.Println("Redis save the roomId error: ", err)
		msg.Result = message.FAILED_TO_SAVE_PLAYER_ROOM_INFO
		utils.Tell(playerId, msg)
		return
	}

	utils.Tell(playerId, msg)
}
