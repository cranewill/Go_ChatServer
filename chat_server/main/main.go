package main

import (
	"fmt"

	netServer "chat_server/netserver"
	manager "chat_server/manager"
)

// main
func main() {
	// Init message handlers
	netServer.Pool.Init()
	// fmt.Println(netServer.Pool.Handlers["create"] == nil)
	// Init managers
	manager.Pool.Init()
	// fmt.Println(manager.Pool.ChatManager)

	go netServer.NetListen()

	for {
		var input string
		fmt.Println("Input 'quit' to close server.")
		fmt.Scanln(&input)
		if input == "quit" {
			break
		}
	}
}