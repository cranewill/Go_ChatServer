package main

import (
	"fmt"

	"Go_ChatServer/chat_server/manager"
	"Go_ChatServer/chat_server/netserver"
)

// main
func main() {
	// Init message handlers
	netserver.Pool.Init()
	// Init managers
	manager.Pool.Init()

	go netserver.Start()
	fmt.Println("Server Started !")

	for {
		var input string
		fmt.Println("Input 'quit' to close server.")
		fmt.Scanln(&input)
		if input == "quit" {
			break
		}
	}
}
