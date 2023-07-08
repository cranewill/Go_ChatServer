package handler

import (
	"Go_ChatServer/chat_client/models"
	"Go_ChatServer/chat_client/utils"
	"Go_ChatServer/chat_common/message"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func HandleChat(player *models.Player) bool {
	fmt.Println("--- Input the message('quit' to back to menu): ")
	for {
		var msg string
		var inputReader *bufio.Reader
		inputReader = bufio.NewReader(os.Stdin)
		msg, err := inputReader.ReadString('\n')
		if err != nil {
			log.Println("Fail to read data, ", err)
		}
		msg = strings.Replace(msg, "\n", "", -1)
		if msg == "quit" {
			break
		}
		fmt.Println("【You】: ", msg, "\n")
		utils.MessageSend(player, message.ReqChatMessage{
			Id:      "chat",
			Content: msg,
		})
	}
	return true
}

// ShowChat shows the chat content of the members in the same chat room
func ShowChat(msgData []byte) {
	var msg message.ResChatMessage
	err := json.Unmarshal(msgData, &msg)
	if err != nil {
		log.Println("Json decode error: ", err)
		return
	}
	fmt.Println("【", msg.PlayerId, "】: ", msg.Content, "\n")
}
