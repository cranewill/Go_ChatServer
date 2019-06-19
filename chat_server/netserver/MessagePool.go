package netserver

import (
	"fmt"

	handler "chat_server/handler"
)

type MessagePool struct {
	Handlers map[string]handler.Handler
}

var Pool MessagePool

// Init registers all the messages
func (pool *MessagePool)Init() {
	fmt.Println("Init MessagePool ...")
	
	// message register
	pool.Handlers = map[string]handler.Handler{}
	pool.Handlers["create"] = handler.CreateRoomHandler{}
	pool.Handlers["show"] = handler.ShowRoomsHandler{}
}