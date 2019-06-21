package manager

import (
	"fmt"

	chat "chat_server/chat"
	utils "chat_server/utils"
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

	manager.ChatRooms[playerId] = chat.ChatRoom{Owner:playerId, Members:[]int64{playerId}}

	fmt.Printf("Create chat room success: %v\n", manager.ChatRooms[playerId])


}

// ShowChatRooms shows all the openning rooms to player
func (manager ChatManager) ShowChatRooms(playerId int64) {
	for roomId, room := range manager.ChatRooms {
		fmt.Println("Room id (", roomId, "), room members:", room.Members)
	}

	utils.TellPlayer(playerId, "hello")
}

// EnterChatRoom allows player enter the specific chat room
func (manager *ChatManager) EnterChatRoom(playerId int64, roomId int64) {
	if playerId == roomId {
		fmt.Println("PlayerId == RoomId")
		return
	}
	room, exist := manager.ChatRooms[roomId]
	if !exist {
		fmt.Println("Room(", roomId, ") not exist!")
		return
	}
	for _, member := range room.Members {
		if member == playerId {
			fmt.Println("Player is in the chat room already")
			return
		}
	}

	room.Members = append(room.Members, playerId)
	manager.ChatRooms[roomId] = room
}