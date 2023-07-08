package handler

import (
	"Go_ChatServer/chat_common/common_utils"
	"Go_ChatServer/chat_common/message"
	"Go_ChatServer/chat_server/chat"
	"Go_ChatServer/chat_server/db"
	"Go_ChatServer/chat_server/models"
	"Go_ChatServer/chat_server/utils"
	"encoding/json"
	"log"

	"Go_ChatServer/chat_server/manager"
)

type CreateRoomHandler struct {
}

func (handler CreateRoomHandler) Handle(player *models.Player, msgMap []byte) {
	msg := message.ReqCreateChatRoomMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		log.Println("Decode message error: ", err)
		return
	}
	CreateChatRoom(player.Id)
}

// CreateChatRoom creates a new chat room
func CreateChatRoom(playerId string) {
	msg := message.ResCreateResultMessage{
		Id:     "create",
		Result: message.SUCCESS,
	}

	_, err := db.Redisdb.Get(db.PlayerRoomKey(playerId)).Result()
	if err == nil {
		log.Println("Player is in a chat room now, please quit first")
		msg.Result = message.PLAYER_NOT_IN_CHAT_ROOM
		utils.Tell(playerId, msg)
		return
	} else if !db.IsNil(err) {
		log.Println("Get playerRoomId error: ", err)
		msg.Result = message.REDIS_ERROR
		utils.Tell(playerId, msg)
		return
	}
	mgr := manager.Pool.ChatManager
	roomId := genRoomUid()
	mgr.ChatRooms[roomId] = &chat.ChatRoom{Owner: playerId, Members: make([]string, 0)}
	mgr.ChatRooms[roomId].Members = append(mgr.ChatRooms[roomId].Members, playerId)

	// db save which room the player is in
	err = db.Redisdb.Set(db.PlayerRoomKey(playerId), roomId, 0).Err()
	if err != nil {
		log.Println("Redis save the roomId error: ", err)
		msg.Result = message.FAILED_TO_SAVE_PLAYER_ROOM_INFO
		utils.Tell(playerId, msg)
		return
	}

	utils.Tell(playerId, msg)
}

func genRoomUid() string {
	uid, _ := db.Redisdb.Incr(db.RoomUidKey()).Result()
	return "room:" + common_utils.ToString(uid)
}
