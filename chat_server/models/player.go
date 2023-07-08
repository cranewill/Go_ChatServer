package models

import "net"

type Player struct {
	Id   string
	Conn *net.Conn
}

func NewPlayer(id string, conn *net.Conn) *Player {
	return &Player{
		Id:   id,
		Conn: conn,
	}
}
