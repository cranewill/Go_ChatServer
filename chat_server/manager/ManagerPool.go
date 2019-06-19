package manager

import (
	"fmt"

	chat "chat_server/chat"
)

type ManagerPool struct {
	ChatManager ChatManager
}

var Pool ManagerPool

// Init registers all the managers
func (pool *ManagerPool)Init() {
	fmt.Println("Init ManagerPool ...")

	// manager register
	pool.ChatManager = ChatManager{ChatRooms:map[int64]chat.ChatRoom{}}
}