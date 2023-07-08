package message

type ReqCreateChatRoomMessage struct {
	Id string
}

func (r ReqCreateChatRoomMessage) MsgName() string {
	return "ReqCreateChatRoomMessage"
}
