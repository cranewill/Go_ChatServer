package netserver

import (
	"Go_ChatServer/chat_common/message"
	"Go_ChatServer/chat_server/models"
	"encoding/json"
	"log"
	"net"
	"sync"

	connect "Go_ChatServer/chat_server/connect"
	"Go_ChatServer/chat_server/utils"
)

// Start starts the socket server and deals message receiving and sending
func Start() {
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("Server Listen TCP Error: ", err)
		return
	}

	// Init connect pool
	connect.Pool = connect.PlayerPool{
		Players: new(sync.Map),
	}

	// message sending channel
	utils.MsgSendChan = make(chan utils.MessageTask, 20)
	go utils.SendMsg(utils.MsgSendChan)

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Connect Client Error: ", err)
			continue
		}
		player := models.NewPlayer("", &conn)
		go handle(player)
	}
}

// handle deals message accepted from client
func handle(player *models.Player) {
	if player == nil || player.Conn == nil {
		log.Println("Received NULL Connection")
		return
	}

	for {
		// create buf to accept message
		buf := make([]byte, 4096)
		length, err := (*player.Conn).Read(buf)
		if err != nil {
			log.Println("Read Message Error: ", err)
			connect.Pool.RemovePlayer(player)
			return
		}

		// create a map to decode msg-json
		netMsg := message.NetMessage{}
		err = json.Unmarshal(buf[:length], &netMsg)
		if err != nil {
			log.Println("Net Message Format Error: ", err)
			continue
		}
		msgName := netMsg.MsgName
		// the first message must be "ReqAuthMessage"
		if player.Id == "" && msgName != "ReqAuthMessage" {
			log.Println("Message Illegal: ", msgName)
			return
		}

		//switch msgName {
		//case "ReqAuthMessage":
		//	if player.Id != "" {
		//		log.Println("Repeat Auth")
		//		break
		//	}
		//	msgBody := message.ReqAuthMessage{}
		//	err = json.Unmarshal([]byte(netMsg.Data), &msgBody)
		//	if err != nil {
		//		log.Println("Message Body Format Error: ", err)
		//		continue
		//	}
		//	player.Id = msgBody.PlayerId
		//	connect.Pool.SavePlayer(player)
		//default:
		msgHandler, ok := Pool.Handlers[msgName]
		if !ok {
			log.Println("Cannot find handler for command: ", msgName)
			continue
		}
		msgHandler.Handle(player, []byte(netMsg.Data))
		//}
	}
}
