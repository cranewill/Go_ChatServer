package messageout

import ()

type RoomInfo struct {
	RoomId int64
	Members []int64
}

type ResRoomsMessage struct {
	Id string
	Rooms []RoomInfo
}

func (message ResRoomsMessage) ToString() {

}