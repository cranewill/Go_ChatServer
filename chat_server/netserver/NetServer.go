package netserver

import (
	"fmt"
	"net"
	"encoding/json"
)

// NetListen starts the socket server and deals messages
func NetListen() {
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Start Server Error: ", err)
		return
	}
	fmt.Println("Server Started !")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Connect Client Error: ", err)
			continue
		}
		handle(conn)
	}
}

// handler deals message accepted from client
func handle(conn net.Conn) {
	if conn == nil {
		fmt.Println("Received NULL Connection")
		return
	}

	// create buf to accept message
	buf := make([]byte, 4096)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read Message ERROR: ", err)
		return
	}
	fmt.Printf("Message is %s\n", buf)
	
	// create map to decode msg-json
	msgMap := make(map[string]interface{})
	err = json.Unmarshal(buf[:len], &msgMap)
	if err != nil {
		fmt.Println("Message Format Error: ", err)
		return
	}
	
	// find the right handler to deal logic
	cmd := fmt.Sprintf("%v", msgMap["id"])
	Pool.Handlers[cmd].Deal(buf[:len])
}