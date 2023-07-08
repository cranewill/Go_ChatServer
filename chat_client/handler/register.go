package handler

import (
	"Go_ChatServer/chat_client/models"
)

func init() {
	RegisterReqRoutes()
	RegisterResRoutes()
}

// request routes
var ReqRoutes map[string]func(player *models.Player) bool

// response routes
var ResRoutes map[string]func([]byte)

func RegisterReqRoutes() {
	ReqRoutes = make(map[string]func(player *models.Player) bool)
	ReqRoutes["create"] = HandleCreate
	ReqRoutes["show"] = HandleShow
	ReqRoutes["quit"] = HandleQuit
	ReqRoutes["enter"] = HandleEnter
	ReqRoutes["chat"] = HandleChat
}

func RegisterResRoutes() {
	ResRoutes = make(map[string]func([]byte))
	ResRoutes["ResRoomsMessage"] = ShowRooms
	ResRoutes["ResCreateResultMessage"] = CreateResult
	ResRoutes["ResQuitResultMessage"] = QuitResult
	ResRoutes["ResEnterResultMessage"] = EnterResult
	ResRoutes["ResChatMessage"] = ShowChat
	ResRoutes["ResAuthMessage"] = AuthResult
}
