package message

type ResChatMessage struct {
	Id       string
	PlayerId string
	Content  string
}

func (r ResChatMessage) MsgName() string {
	return "ResChatMessage"
}
