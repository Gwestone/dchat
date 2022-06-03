package utils

type MessageData struct {
	MessageText string `json:"MessageText"`
	From        string `json:"From"`
	To          string `json:"To"`
	Date        int    `json:"Date"`
}

func NewMessageData() *MessageData {
	return &MessageData{}
}
