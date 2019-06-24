package manager

import (
	"fmt"
	"net"
)

type ConnectPool struct {
	Conns map[int64]net.Conn
}

var Pool ConnectPool

// SaveConn saves a player tcp connection in Conns
func (this *ConnectPool) SaveConn (playerId int64, con net.Conn) {
	_, exist := this.Conns[playerId]
	if (exist) {
		//fmt.Println("Player ", playerId, " connection already saved")
		return
	}
	this.Conns[playerId] = con
	fmt.Println("Player ", playerId, " connect!")
}

// RemoveConn called when client disconnects to remove player tcp connection from Conns
func (this *ConnectPool) RemoveConn (playerId int64) {
	delete(this.Conns, playerId)
}