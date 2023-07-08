package message

type ReqChatMessage struct {
	Id      string
	Content string
}

func (r ReqChatMessage) MsgName() string {
	return "ReqChatMessage"
}
