package main

import (
	"fmt"

	netServer "chat_server/netserver"
	manager "chat_server/manager"
)

// main
func main() {
	manager.Pool.Init()
	fmt.Println(manager.Pool.Handlers["create"] == nil)

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