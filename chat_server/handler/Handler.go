package handler

import (
	// "fmt"

)

type Handler interface {
	Deal(msg []byte)
}

// func Deal(msg map[string]interface{}) {
// 	fmt.Println("deal", msg)
// }


