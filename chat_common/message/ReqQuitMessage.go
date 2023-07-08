package message

type ReqQuitMessage struct {
	Id string
}

func (r ReqQuitMessage) MsgName() string {
	return "ReqQuitMessage"
}
