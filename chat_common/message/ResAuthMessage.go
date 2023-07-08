package message

type ResAuthMessage struct {
	Result bool
}

func (r ResAuthMessage) MsgName() string {
	return "ResAuthMessage"
}
