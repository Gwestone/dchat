package utils

type MessageData struct {
	Message string `json:"message"`
	From    int    `json:"from"`
	To      int    `json:"to"`
	Date    int    `json:"date"`
}

func NewMessageData() *MessageData {
	return &MessageData{}
}
