package manager

import (
	"encoding/json"
	"fmt"
	"log"

	client_message "Go_ChatServer/chat_client/message"
)

// ShowRooms displays all the existing chat room info
func ShowRooms(message []byte) {
	var msg client_message.ResRoomsMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	if len(msg.Rooms) == 0 {
		log.Println("There is no chat room now. You can create a new chat room with 'create' command")
		return
	}
	for _, room := range msg.Rooms {
		log.Println("id:", room.RoomId, "\tmembers:", room.Members)
	}
}

// CreateResult returns the result of creating chat room
func CreateResult(message []byte) {
	var msg client_message.ResCreateResultMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	switch msg.Result {
	case 0:
		log.Println("Create chat room success !")
	case 1:
		log.Println("Cannot find the room !")
	case 2:
		log.Println("You are in a chat room now !")
	case 3:
		log.Println("Save room info error !")
	}
}

// EnterResult returns the result of entering chat room
func EnterResult(message []byte) {
	var msg client_message.ResEnterResultMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	switch msg.Result {
	case 0:
		log.Println("Enter chat room success !")
	case 1:
		log.Println("Cannot find the room !")
	case 2:
		log.Println("You are in a chat room now !")
	case 3:
		log.Println("Room not exist !")
	case 4:
		log.Println("You have been in this room now !")
	case 5:
		log.Println("Save room info error !")
	}
}

// QuitResult returns the result of quiting chat room
func QuitResult(message []byte) {
	var msg client_message.ResQuitResultMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	switch msg.Result {
	case 0:
		log.Println("Quit chat room success !")
	case 1:
		log.Println("Cannot find the room !")
	case 2:
		log.Println("You are not in a chat room now !")
	case 3:
		log.Println("Save room info error !")
	}
}

// ShowChat shows the chat content of the members in the same chat room
func ShowChat(message []byte) {
	var msg client_message.ResChatMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	fmt.Println("【", msg.PlayerId, "】: ", msg.Content, "\n")
}
