package origin

type RespSendMessage struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	MessageId int    `json:"messageId"`
}

type Message interface {
	JustMark() //这只是一个标记，防止用户乱传类型
}

type MessagePlain struct {
	Message
	Type string `json:"type"`
	Text string `json:"text"`
}
