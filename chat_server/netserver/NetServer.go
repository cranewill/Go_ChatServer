package netserver

import (
	"encoding/json"
	"fmt"
	"net"

	connect "Go_ChatServer/chat_server/connect"
	"Go_ChatServer/chat_server/utils"
)

// Start starts the socket server and deals message receiving and sending
func Start() {
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Server Listen TCP Error: ", err)
		return
	}

	// Init connect pool
	connect.Pool = connect.ConnectPool{map[int64]net.Conn{}}

	// message sending channel
	utils.SendChan = make(chan utils.MessageSendTask, 20)
	go utils.Send(utils.SendChan)

	for {
		conn, err := server.Accept()
		//fmt.Println("server accepts msg!!!!!!!!!!!!!")
		defer conn.Close()
		if err != nil {
			fmt.Println("Connect Client Error: ", err)
			continue
		}
		go handle(conn)
	}
}

// handle deals message accepted from client
func handle(conn net.Conn) {
	if conn == nil {
		fmt.Println("Received NULL Connection")
		return
	}

	for {
		// create buf to accept message
		buf := make([]byte, 4096)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read Message Error: ", err)
			// TODO should we remove connect from connectPool here?
			return
		}
		fmt.Printf("Message is %s\n", buf)

		// create a map to decode msg-json
		msgMap := make(map[string]interface{})
		err = json.Unmarshal(buf[:len], &msgMap)
		if err != nil {
			fmt.Println("Message Format Error: ", err)
			continue
		}

		// save player's connection and find the right handler to deal logic
		cmd := fmt.Sprintf("%v", msgMap["id"])
		playerId := int64(msgMap["playerId"].(float64))

		connect.Pool.SaveConn(playerId, conn)
		Pool.Handlers[cmd].Deal(buf[:len])
	}
}
