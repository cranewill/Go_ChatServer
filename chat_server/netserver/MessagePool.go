package netserver

import (
	"Go_ChatServer/chat_server/handler"
	"log"
)

type MessagePool struct {
	Handlers map[string]handler.Handler
}

var Pool MessagePool

// Init registers all the messages
func (pool *MessagePool) Init() {
	log.Println("Init MessagePool ...")

	// message register
	pool.Handlers = map[string]handler.Handler{}
	pool.Handlers["ReqAuthMessage"] = handler.AuthHandler{}
	pool.Handlers["ReqCreateChatRoomMessage"] = handler.CreateRoomHandler{}
	pool.Handlers["ReqShowRoomsMessage"] = handler.ShowRoomsHandler{}
	pool.Handlers["ReqEnterChatRoomMessage"] = handler.EnterRoomHandler{}
	pool.Handlers["ReqChatMessage"] = handler.ChatHandler{}
	pool.Handlers["ReqQuitMessage"] = handler.QuitHandler{}
}
