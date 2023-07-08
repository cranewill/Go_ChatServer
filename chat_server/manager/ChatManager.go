package manager

import (
	"Go_ChatServer/chat_server/chat"
)

type ChatManager struct {
	ChatRooms map[string]*chat.ChatRoom
}
