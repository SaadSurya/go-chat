package chats

import (
	"time"

	"github.com/saadsurya/go-chat/messages"
	"github.com/saadsurya/go-chat/users"
)

type Chat struct {
	messages.Message
	users.User
	LastMessagedAt time.Time `json:"lastMessagedAt"`
}
