package handler

import "Go_ChatServer/chat_server/models"

type Handler interface {
	Handle(player *models.Player, msg []byte)
}
