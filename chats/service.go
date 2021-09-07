package chats

import "github.com/saadsurya/go-chat/database"

var (
	getChatsQuery = `with chat_users AS (
		select id, text, "from", "to", created_at,
			case "from" when ? then "to" else "from" end chat_user
		from messages
		where ("from" = ? or "to" = ?)
	), last_messages AS (
		select chat_user, max(created_at) last_messaged_at from chat_users group by chat_user
	)
	select c.id, text, "from", "to", username, first_name, last_name, l.last_messaged_at
	from chat_users c
	inner join last_messages l on c.chat_user = l.chat_user and c.created_at = l.last_messaged_at
	inner join users u on u.username = c.chat_user`
)

func GetChats(username string) map[string]Chat {
	db := database.DBConn
	var chats []Chat
	db.Raw(getChatsQuery, username, username, username).Scan(&chats)
	chatsMap := make(map[string]Chat, len(chats))
	for _, chat := range chats {
		chatsMap[chat.Username] = chat
	}
	return chatsMap
}
