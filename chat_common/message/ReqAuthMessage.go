package message

type ReqAuthMessage struct {
	PlayerId string
}

func (r ReqAuthMessage) MsgName() string {
	return "ReqAuthMessage"
}
