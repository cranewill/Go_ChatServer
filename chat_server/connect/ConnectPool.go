package manager

import (
	"log"
	"net"
)

type ConnectPool struct {
	Conns map[int64]net.Conn
}

var Pool ConnectPool

// SaveConn saves a player tcp connection in Conns
func (this *ConnectPool) SaveConn(playerId int64, con net.Conn) {
	_, exist := this.Conns[playerId]
	if exist {
		//fmt.Println("Player ", playerId, " connection already saved")
		return
	}
	this.Conns[playerId] = con
	log.Println("Player ", playerId, " connects!")
}

// RemoveConn called when client disconnects to remove player tcp connection from Conns
func (this *ConnectPool) RemoveConn(playerId int64) {
	delete(this.Conns, playerId)
	log.Println("Player ", playerId, " disconnects!")
}

// RemoveTheConn execute player connection delete by giving the connection
func (this *ConnectPool) RemoveTheConn(con net.Conn) {
	for playerId, conn := range this.Conns {
		if conn == con {
			delete(this.Conns, playerId)
			log.Println("Player ", playerId, " disconnects!")
		}
	}
}
