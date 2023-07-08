package message

import ()

type ReqQuitMessage struct {
	Id       string
	PlayerId int64
	QuitType string
}
