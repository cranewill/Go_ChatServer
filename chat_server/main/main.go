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
	// Init managers
	manager.Pool.Init()

	go netServer.Start()

	for {
		var input string
		fmt.Println("Input 'quit' to close server.")
		fmt.Scanln(&input)
		if input == "quit" {
			break
		}
	}
}