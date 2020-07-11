package origin

type RespSendMessage struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	MessageId int    `json:"messageId"`
}

type Message interface {
	JustMark()
}

type MessagePlain struct {
	Message
	Type string `json:"type"`
	Text string `json:"text"`
}
