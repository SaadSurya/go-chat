package messages

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}
