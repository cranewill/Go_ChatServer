package manager

import (
	// "fmt"

)

type Handler interface {
	Deal(msg map[string]interface{})
}

// func Deal(msg map[string]interface{}) {
// 	fmt.Println("deal", msg)
// }


