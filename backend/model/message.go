package model

import "time"

type MessageType int

const (
	MessageTypeText  = 1
	MessageTypeImage = 2
	MessageTypeStamp = 3
)

var MessageTypeStrings = map[string]MessageType{
	"text":  MessageTypeText,
	"image": MessageTypeImage,
	"stamp": MessageTypeStamp,
}

func (m MessageType) String() string {
	for k, v := range MessageTypeStrings {
		if v == m {
			return k
		}
	}

	return ""
}

type Message struct {
	ID       string
	UserID   string
	Type     string
	Content  string
	PostedAt time.Time
}
