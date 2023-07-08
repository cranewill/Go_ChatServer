package message

type RoomInfo struct {
	RoomId  string
	Members []string
}

type ResRoomsMessage struct {
	Id    string
	Rooms []RoomInfo
}

func (r ResRoomsMessage) MsgName() string {
	return "ResRoomsMessage"
}
