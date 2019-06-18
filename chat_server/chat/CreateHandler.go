package chat

import (
	"fmt"
	// netserver "chat_server/netserver"
)

type CreateHandler struct {

}

func (handler CreateHandler)Deal(msg map[string]interface{}) {
	fmt.Println("CreateHandler deal...", msg)
}


