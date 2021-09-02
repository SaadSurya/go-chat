package chats

import (
	"github.com/saadsurya/go-chat/messages"
	"github.com/saadsurya/go-chat/users"
)

type Chat struct {
	messages.Message
	users.User
}
