package message

type ResEnterResultMessage struct {
	Id     string
	Result string
}

func (r ResEnterResultMessage) MsgName() string {
	return "ResEnterResultMessage"
}
