package miraiGo

//身份验证请求
type AuthReq struct {
	AuthKey string `json:"authKey"`
}

//校验(或释放)请求
type VerifyReq struct {
	SessionKey string `json:"sessionKey"`
	BotQQ uint `json:"qq"`
}
