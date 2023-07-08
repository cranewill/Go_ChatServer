package message

type ResCreateResultMessage struct {
	Id     string
	Result string
}

func (r ResCreateResultMessage) MsgName() string {
	return "ResCreateResultMessage"
}
