package message

import (

)

type ReqEnterChatRoomMessage struct {
	Id string
	PlayerId int64
	RoomId int64
}