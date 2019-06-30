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
	pool.Handlers["create"] = handler.CreateRoomHandler{}
	pool.Handlers["show"] = handler.ShowRoomsHandler{}
	pool.Handlers["enter"] = handler.EnterRoomHandler{}
	pool.Handlers["chat"] = handler.ChatHandler{}
	pool.Handlers["quit"] = handler.QuitHandler{}
}

func CallBack(msg []byte, f func([]byte)) {
	f(msg)
}
