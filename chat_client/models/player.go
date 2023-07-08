package models

import "net"

type Player struct {
	Id   string
	Conn *net.Conn
}
