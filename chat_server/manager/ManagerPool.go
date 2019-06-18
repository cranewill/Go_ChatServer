package manager

import (
	"fmt"

	chat "chat_server/chat"
)

type ManagerPool struct {
	Handlers map[string]Handler
}

var Pool ManagerPool

// Init assigns all the entries of messages and handlers
func (managerPool *ManagerPool)Init() {
	fmt.Println("Init ManagerPool ...")
	
	*managerPool.Handlers = map[string]Handler{}
	*managerPool.Handlers["create"] = chat.CreateHandler{}
}