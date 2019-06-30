package manager

import (
	"container/list"
	"log"
	"strconv"

	"Go_ChatServer/chat_server/chat"
	"Go_ChatServer/chat_server/redis"
	"Go_ChatServer/chat_server/utils"

	connect "Go_ChatServer/chat_server/connect"
	out "Go_ChatServer/chat_server/messageout"
)

type ChatManager struct {
	ChatRooms map[int64]*chat.ChatRoom
}

// CreateChatRoom creates a new chat room
func (manager *ChatManager) CreateChatRoom(playerId int64) {
	message := out.ResCreateResultMessage{Id: "create", Result: 0}

	_, err := redis.Redisdb.Get(strconv.FormatInt(playerId, 10)).Result()
	if err == nil {
		log.Println("Player is in a chat room now, please quit first")
		message.Result = 2
		utils.Tell(playerId, message)
		return
	} else if !redis.IsNil(err) {
		log.Println("Get playerRoomId error: ", err)
		message.Result = 1
		utils.Tell(playerId, message)
		return
	}

	manager.ChatRooms[playerId] = &chat.ChatRoom{Owner: playerId, Members: list.New()}
	manager.ChatRooms[playerId].Members.PushBack(playerId)

	// redis save which room the player is in
	err = redis.Redisdb.Set(strconv.FormatInt(playerId, 10), playerId, 0).Err()
	if err != nil {
		log.Println("Redis save the roomId error: ", err)
		message.Result = 3
		utils.Tell(playerId, message)
		return
	}

	utils.Tell(playerId, message)
}

// ShowChatRooms shows all the opening rooms to player
func (manager *ChatManager) ShowChatRooms(playerId int64) {
	message := out.ResRoomsMessage{Id: "rooms"}

	for roomId, room := range manager.ChatRooms {
		roomInfo := out.RoomInfo{RoomId: roomId, Members: []int64{}}

		var memberStr string
		for memberId := room.Members.Front(); memberId != nil; memberId = memberId.Next() {
			memberStr += strconv.FormatInt(memberId.Value.(int64), 10) + ";"

			roomInfo.Members = append(roomInfo.Members, memberId.Value.(int64))
		}
		//log.Println("Room id (", roomId, "), room members:", memberStr)
		message.Rooms = append(message.Rooms, roomInfo)
	}
	utils.Tell(playerId, message)
}

// EnterChatRoom allows player to enter the specific chat room
func (manager *ChatManager) EnterChatRoom(playerId int64, roomId int64) {
	message := out.ResEnterResultMessage{Id: "enter", Result: 0}

	playerRoomId, err := redis.Redisdb.Get(strconv.FormatInt(playerId, 10)).Result()
	if err == nil {
		log.Println("Player is in a chat room now, please quit first")
		message.Result = 2
		utils.Tell(playerId, message)
		return
	} else if !redis.IsNil(err) {
		log.Println("Get playerRoomId error: ", err)
		message.Result = 1
		utils.Tell(playerId, message)
		return
	}
	room, exist := manager.ChatRooms[roomId]
	if !exist {
		log.Println("Room(", roomId, ") not exist!")
		message.Result = 3
		utils.Tell(playerId, message)
		return
	}
	if playerRoomId == strconv.FormatInt(roomId, 10) {
		log.Println("Player is in this chat room now")
		message.Result = 4
		utils.Tell(playerId, message)
		return
	}

	(*room).Members.PushBack(playerId)
	//manager.ChatRooms[roomId] = room

	// redis save this player and chat room
	err = redis.Redisdb.Set(strconv.FormatInt(playerId, 10), roomId, 0).Err()
	if err != nil {
		log.Println("Redis save the roomId error: ", err)
		message.Result = 5
		utils.Tell(playerId, message)
		return
	}

	utils.Tell(playerId, message)
}

// Chat receives player's words and broadcasts to every player in the chat room
func (manager ChatManager) Chat(playerId int64, content string) {
	roomIdStr, err := redis.Redisdb.Get(strconv.FormatInt(playerId, 10)).Result()
	if err != nil {
		if !redis.IsNil(err) {
			log.Println("Get player roomId error: ", err)
			return
		} else {
			log.Println("Player is not in a chat room")
			return
		}
	}

	roomId, _ := strconv.ParseInt(roomIdStr, 10, 64)
	room := manager.ChatRooms[roomId]

	message := out.ResChatMessage{Id: "chat", PlayerId: playerId, Content: content}
	for memberId := room.Members.Front(); memberId != nil; memberId = memberId.Next() {
		mId, ok := memberId.Value.(int64)
		if mId != playerId && ok {
			utils.Tell(mId, message)
		}
	}
}

// Quit allows player to quit the chat room, and the same time removes its info from redis
func (manager *ChatManager) Quit(playerId int64, quitType string) {
	//if quitType == "a" {
	//	connect.Pool.RemoveConn(playerId)
	//}

	message := out.ResQuitResultMessage{Id: "quit", Result: 0}

	roomIdStr, err := redis.Redisdb.Get(strconv.FormatInt(playerId, 10)).Result()
	if err != nil {
		if !redis.IsNil(err) {
			log.Println("Get player roomId error: ", err)
			message.Result = 1
			utils.Tell(playerId, message)
			return
		} else {
			//log.Println("Player is not in a chat room")
			message.Result = 2
			utils.Tell(playerId, message)
			return
		}
	}
	roomId, _ := strconv.ParseInt(roomIdStr, 10, 64)
	room := manager.ChatRooms[roomId]

	for memberId := room.Members.Front(); memberId != nil; memberId = memberId.Next() {
		mId, ok := memberId.Value.(int64)
		if ok && mId == playerId {
			(*room).Members.Remove(memberId)
		}
	}
	//manager.ChatRooms[roomId] = room
	_, err = redis.Redisdb.Del(strconv.FormatInt(playerId, 10)).Result()
	if err != nil {
		log.Println("Delete player-room info in redis error: ", err)
		message.Result = 3
		utils.Tell(playerId, message)
	}

	if room.Members.Len() == 0 {
		delete(manager.ChatRooms, roomId)
		log.Println("Delete chat room ", roomId)
	}

	utils.Tell(playerId, message)
	if quitType == "a" {
		connect.Pool.RemoveConn(playerId)
	}
}
