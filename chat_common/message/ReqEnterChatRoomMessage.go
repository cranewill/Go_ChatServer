package message

type ReqEnterChatRoomMessage struct {
	Id     string
	RoomId string
}

func (r ReqEnterChatRoomMessage) MsgName() string {
	return "ReqEnterChatRoomMessage"
}
