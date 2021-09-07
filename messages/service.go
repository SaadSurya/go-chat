package messages

import (
	"sort"
	"time"

	"github.com/saadsurya/go-chat/database"
)

func CreateMessage(message *Message) {
	db := database.DBConn
	db.Create(&message)
}

func Retrieve(username string, ofUser string, before time.Time, limit int) []Message {
	db := database.DBConn
	var messages []Message
	db.Where("created_at < ? and ((\"from\" = ? and \"to\" = ?) or (\"from\" = ? and \"to\" = ?))", before, username, ofUser, ofUser, username).Limit(limit).Order("created_at desc").Find(&messages)
	sort.Slice(messages, func(i, j int) bool { return messages[i].CreatedAt.Before(messages[j].CreatedAt) })
	return messages
}
