package miraiGo

import (
	"bytes"
	"errors"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type session struct {
	miraiAddr string
	sessionKey string
	qq uint
}

func (this *session) GetBindQQ() uint {
	return this.qq
}

//验证你的身份，并创建一个新的会话
func NewSession(miraiAddr string,authKey string) (*session,error)  {
	req:=AuthReq{AuthKey: authKey}
	data, err := jsoniter.Marshal(&req)
	if err!=nil{
		return nil,err
	}

	resp, err := http.DefaultClient.Post("http://"+miraiAddr+"/auth", "application/json; charset=utf-8", bytes.NewBuffer(data))
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil,err
	}

	var authResp AuthResp
	err = jsoniter.Unmarshal(body, &authResp)
	if err!=nil{
		return nil,err
	}

	switch authResp.Code {
	case 0:
		return &session{
			miraiAddr:  miraiAddr,
			sessionKey: authResp.SessionKey,
		},nil
	default:
		return nil,errors.New("身份认证失败,Code:"+string(authResp.Code))
	}
}

//校验并激活此会话，并与指定bot绑定
func (this *session) Verify(botqq uint) error  {
	req:=SessionReq{
		SessionKey: this.sessionKey,
		BotQQ:      botqq,
	}

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/verify", "application/json; charset=utf-8", bytes.NewReader(data))
	if err!=nil{
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return err
	}

	var uniResp UniResp
	err = jsoniter.Unmarshal(body, &uniResp)
	if err!=nil{
		return err
	}

	err = GetErrByCode(uniResp.Code)
	if err==nil{
		this.qq=botqq
	}
	return err
}

//释放这个Session
func (this *session) Release() error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}
	req:=SessionReq{
		SessionKey: this.sessionKey,
		BotQQ:      this.qq,
	}
	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/release", "application/json; charset=utf-8", bytes.NewReader(data))
	if err!=nil{
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return err
	}

	var uniResp UniResp
	err = jsoniter.Unmarshal(body, &uniResp)
	if err!=nil{
		return err
	}

	err = GetErrByCode(uniResp.Code)
	return err
}