package manager

import (
	"Go_ChatServer/chat_server/models"
	"log"
	"sync"
)

type PlayerPool struct {
	Players *sync.Map
}

var Pool PlayerPool

// GetConn returns the connection by playerId
func (this *PlayerPool) GetPlayer(playerId string) (*models.Player, bool) {
	conn, ok := this.Players.Load(playerId)
	if !ok {
		return nil, false
	}
	return conn.(*models.Player), ok
}

// SaveConn saves a player tcp connection in Conns
func (this *PlayerPool) SavePlayer(player *models.Player) {
	this.Players.Store(player.Id, player)
	log.Println("Player ", player.Id, " connects!")
}

// RemoveConn called when client disconnects to remove player tcp connection from Conns
func (this *PlayerPool) RemovePlayer(player *models.Player) {
	this.Players.Delete(player.Id)
	log.Println("Player ", player.Id, " disconnects!")
}
