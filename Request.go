package miraiGo

//身份验证请求
type AuthReq struct {
	AuthKey string `json:"authKey"`
}

//校验(或释放)请求
type SessionReq struct {
	SessionKey string `json:"sessionKey"`
	BotQQ uint `json:"qq"`
}

type RequestSendFriendMessage struct {
	SessionKey string `json:"sessionKey"`
	Target uint `json:"target"`
	MessageChain []Message `json:"messageChain"`
}

type RequestSendTempMessage struct {
	SessionKey string `json:"sessionKey"`
	QQ uint `json:"qq"`
	Group uint `json:"group"`
	MessageChain []Message `json:"messageChain"`
}

type RequestSendGroupMessage struct {
	SessionKey string `json:"sessionKey"`
	Target uint `json:"target"`
	MessageChain []Message `json:"messageChain"`
}