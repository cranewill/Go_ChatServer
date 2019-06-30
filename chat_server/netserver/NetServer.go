package netserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

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
	connect.Pool = connect.ConnectPool{map[int64]net.Conn{}}

	// message sending channel
	utils.StrSendChan = make(chan utils.MessageSendTask, 20)
	utils.MsgSendChan = make(chan utils.MessageTask, 20)
	go utils.Send(utils.StrSendChan)
	go utils.SendMsg(utils.MsgSendChan)

	for {
		conn, err := server.Accept()
		//fmt.Println("server accepts msg!!!!!!!!!!!!!")
		if err != nil {
			log.Println("Connect Client Error: ", err)
			continue
		}
		go handle(conn)
	}
}

// handle deals message accepted from client
func handle(conn net.Conn) {

	if conn == nil {
		log.Println("Received NULL Connection")
		return
	}

	for {
		// create buf to accept message
		buf := make([]byte, 4096)
		len, err := conn.Read(buf)
		if err != nil {
			log.Println("Read Message Error: ", err)
			connect.Pool.RemoveTheConn(conn)
			return
		}
		//fmt.Printf("Message is %s\n", buf)

		// create a map to decode msg-json
		msgMap := make(map[string]interface{})
		err = json.Unmarshal(buf[:len], &msgMap)
		if err != nil {
			log.Println("Message Format Error: ", err)
			continue
		}

		// save player's connection and find the right handler to deal logic
		cmd := fmt.Sprintf("%v", msgMap["id"])
		playerId := int64(msgMap["playerId"].(float64))

		connect.Pool.SaveConn(playerId, conn)
		Pool.Handlers[cmd].Deal(buf[:len])

	}
}
