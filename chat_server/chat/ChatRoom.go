package chat

import "container/list"

type ChatRoom struct {
	Owner   int64
	Members *list.List
}
