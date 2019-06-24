package manager

import (
	"container/list"
	"fmt"
	"strconv"

	"Go_ChatServer/chat_server/chat"
	"Go_ChatServer/chat_server/utils"
	"Go_ChatServer/chat_server/redis"
)

type ChatManager struct {
	ChatRooms map[int64]chat.ChatRoom
}

// CreateChatRoom creates a new chat room
func (manager *ChatManager) CreateChatRoom(playerId int64) {
	if playerId == 0 {
		fmt.Println("PlayerId is nil when create chat room")
		return
	}

	_,exist := manager.ChatRooms[playerId]
	if exist {
		fmt.Println("Player has created a chat room yet")
		return
	}

	manager.ChatRooms[playerId] = chat.ChatRoom{Owner:playerId, Members:list.New()}
	manager.ChatRooms[playerId].Members.PushBack(playerId)

	// redis save which room the player is in
	err := redis.Redisdb.Set(strconv.FormatInt(playerId,10),playerId,0).Err()
	if err != nil {
		fmt.Println("Redis save the roomId error: ", err)
		return
	}

	fmt.Println("Create chat room success")
	utils.TellPlayer(playerId, "create chat room success!")


}

// ShowChatRooms shows all the opening rooms to player
func (manager *ChatManager) ShowChatRooms(playerId int64) {
	var result string
	for roomId, room := range manager.ChatRooms {
		var memberStr string
		for memberId := room.Members.Front(); memberId != nil; memberId = memberId.Next() {
			memberStr += strconv.FormatInt(memberId.Value.(int64),10) + ";"
		}
		fmt.Println("Room id (", roomId, "), room members:", memberStr)
		result += "Room ID (" + strconv.FormatInt(roomId, 10) + "), Members: " + memberStr + "\n"
	}
	utils.TellPlayer(playerId, result)
}

// EnterChatRoom allows player to enter the specific chat room
func (manager *ChatManager) EnterChatRoom(playerId int64, roomId int64) {
	playerRoomId, err:= redis.Redisdb.Get(strconv.FormatInt(playerId, 10)).Result()
	if err == nil {
		fmt.Println("Player is in a chat room now, please quit first")
		utils.TellPlayer(playerId, "you are in a room now, please quit first!")
		return
	} else if !redis.IsNil(err) {
		fmt.Println("Get playerRoomId error: ", err)
		return
	}
	room, exist := manager.ChatRooms[roomId]
	if !exist {
		fmt.Println("Room(", roomId, ") not exist!")
		utils.TellPlayer(playerId, "room not exist!")
		return
	}
	if playerRoomId == strconv.FormatInt(roomId,10) {
		fmt.Println("Player is in this chat room now")
		utils.TellPlayer(playerId, "you are int this room now!")
		return
	}

	room.Members.PushBack(playerId)
	manager.ChatRooms[roomId] = room

	// redis save this player and chat room
	err = redis.Redisdb.Set(strconv.FormatInt(playerId,10),roomId,0).Err()
	if err != nil {
		fmt.Println("Redis save the roomId error: ", err)
		return
	}

}

// Chat receives player's words and broadcasts to every player in the chat room
func (manager ChatManager) Chat(playerId int64, content string) {
	roomIdStr, err:= redis.Redisdb.Get(strconv.FormatInt(playerId, 10)).Result()
	if err != nil {
		if !redis.IsNil(err) {
			fmt.Println("Get player roomId error: ", err)
			return
		} else {
			fmt.Println("Player is not in a chat room")
			return
		}
	}

	roomId,_ := strconv.ParseInt(roomIdStr,10,64)
	room := manager.ChatRooms[roomId]

	for memberId := room.Members.Front(); memberId != nil; memberId = memberId.Next() {
		mId, ok := memberId.Value.(int64)
		if ok {
			utils.TellPlayer(mId, content)
		}
	}
}

// Quit allows player quit the chat room, and the same time remove its info from redis
func (manager *ChatManager) Quit(playerId int64) {
	roomIdStr, err:= redis.Redisdb.Get(strconv.FormatInt(playerId, 10)).Result()
	if err != nil {
		if !redis.IsNil(err) {
			fmt.Println("Get player roomId error: ", err)
			return
		} else {
			fmt.Println("Player is not in a chat room")
			return
		}
	}
	roomId,_ := strconv.ParseInt(roomIdStr,10,64)
	room := manager.ChatRooms[roomId]

	for memberId := room.Members.Front(); memberId != nil; memberId = memberId.Next() {
		mId, ok := memberId.Value.(int64)
		if ok && mId == playerId{
			room.Members.Remove(memberId)
		}
	}
	manager.ChatRooms[playerId] = room

	if room.Members.Len() == 0 {
		delete(manager.ChatRooms, roomId)
		fmt.Println("delete chat room ", roomId)
	}
}
