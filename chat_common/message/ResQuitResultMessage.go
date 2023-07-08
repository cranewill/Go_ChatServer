package message

type ResQuitResultMessage struct {
	Id     string
	Result string
}

func (r ResQuitResultMessage) MsgName() string {
	return "ResQuitResultMessage"
}
