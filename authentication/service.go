package authentication

import (
	"github.com/saadsurya/go-chat/users"
)

func Authenticate(username string, password string) *users.User {
	user := users.FindByUsernamePassword(username, password)
	if user.Username != username {
		return nil
	}
	return user
}
