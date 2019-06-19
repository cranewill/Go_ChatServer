package manager

import (
	"fmt"

	chat "chat_server/chat"
)

type ChatManager struct {
	ChatRooms map[int64]chat.ChatRoom
}

// CreateChatRoom creates a new chat room
func (manager *ChatManager) CreateChatRoom(playerId int64) {
	fmt.Println("Create chat room...")

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

// ShowChatRooms returns all the openning rooms to player
func (manager ChatManager) ShowChatRooms(playerId int64) {
	fmt.Println("Show chat rooms...")
	for roomId, room := range manager.ChatRooms {
		fmt.Println("Room id (", roomId, "), room members:", room.Members)
	}
}