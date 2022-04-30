package utils

type MessageData struct {
	Message string `json:"Message"`
	From    string `json:"From"`
	To      string `json:"To"`
	Date    int    `json:"Date"`
}

func NewMessageData() *MessageData {
	return &MessageData{}
}
