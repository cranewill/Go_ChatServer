package manager

import (
	"Go_ChatServer/chat_server/chat"
	"log"
)

type ManagerPool struct {
	ChatManager ChatManager
}

var Pool ManagerPool

// Init registers all the managers
func (pool *ManagerPool) Init() {
	log.Println("Init ManagerPool ...")

	// manager register
	pool.ChatManager = ChatManager{ChatRooms: map[string]*chat.ChatRoom{}}
}
