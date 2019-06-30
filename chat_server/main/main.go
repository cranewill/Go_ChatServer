package main

import (
	"Go_ChatServer/chat_server/manager"
	"Go_ChatServer/chat_server/netserver"
	"fmt"
	"log"
)

// main
func main() {
	// Init message handlers
	netserver.Pool.Init()
	// Init managers
	manager.Pool.Init()

	go netserver.Start()
	log.Println("Server Started !")

	for {
		var input string
		log.Println("Input 'quit' to close server.")
		fmt.Scanln(&input)
		if input == "quit" {
			break
		}
	}
}
