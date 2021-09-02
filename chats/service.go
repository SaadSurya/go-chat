package chats

import "github.com/saadsurya/go-chat/database"

var (
	getChatsQuery = `with chat_users AS (
		select id, text, "from", "to", created_at,
			case "from" when ? then "to" else "from" end chat_user
		from messages
		where ("from" = ? or "to" = ?)
	), last_messages AS (
		select chat_user, max(created_at) created_at from chat_users group by chat_user
	)
	select c.*, username, first_name, last_name
	from chat_users c
	inner join last_messages l on c.chat_user = l.chat_user and c.created_at = l.created_at
	inner join users u on u.username = c.chat_user`
)

func GetChats(username string) []Chat {
	db := database.DBConn
	var chats []Chat
	db.Raw(getChatsQuery, username, username, username).Scan(&chats)
	return chats
}
