package message

type ReqShowRoomsMessage struct {
	Id string
}

func (r ReqShowRoomsMessage) MsgName() string {
	return "ReqShowRoomsMessage"
}
