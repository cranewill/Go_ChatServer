package manager

import (
	"Go_ChatServer/chat_server/chat"
	"log"
)

type ManagerPool struct {
	// ConnectManager ConnectManager
	ChatManager ChatManager
}

var Pool ManagerPool

// Init registers all the managers
func (pool *ManagerPool) Init() {
	log.Println("Init ManagerPool ...")

	// manager register
	// pool.ConnectManager = ConnectManager{Conns:map[int64]net.Conn{}}
	pool.ChatManager = ChatManager{ChatRooms: map[int64]*chat.ChatRoom{}}
}
