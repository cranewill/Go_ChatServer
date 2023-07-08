package handler

import (
	"Go_ChatServer/chat_common/message"
	"Go_ChatServer/chat_server/models"
	"Go_ChatServer/chat_server/utils"
	"encoding/json"
	"log"

	"Go_ChatServer/chat_server/manager"
)

type ShowRoomsHandler struct {
}

func (handler ShowRoomsHandler) Handle(player *models.Player, msgMap []byte) {
	msg := message.ReqShowRoomsMessage{}
	err := json.Unmarshal(msgMap, &msg)
	if err != nil {
		log.Println("Decode message error: ", err)
		return
	}
	ShowChatRooms(player.Id)
}

// ShowChatRooms shows all the opening rooms to player
func ShowChatRooms(playerId string) {
	msg := message.ResRoomsMessage{
		Id:    "rooms",
		Rooms: make([]message.RoomInfo, 0),
	}

	for roomId, room := range manager.Pool.ChatManager.ChatRooms {
		roomInfo := message.RoomInfo{
			RoomId:  roomId,
			Members: make([]string, len(room.Members)),
		}
		for i, memberId := range room.Members {
			roomInfo.Members[i] = memberId
		}
		msg.Rooms = append(msg.Rooms, roomInfo)
	}
	utils.Tell(playerId, msg)
}
