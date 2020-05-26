package miraiGo

import (
	"bytes"
	"errors"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type session struct {
	client *http.Client
	miraiAddr string
	sessionKey string
	botqq uint
}

//验证你的身份，并创建一个新的会话
func NewSession(miraiAddr string,authKey string) (*session,error)  {
	req:=AuthReq{AuthKey: authKey}
	data, err := jsoniter.Marshal(&req)
	if err!=nil{
		return nil,err
	}

	client:=&http.Client{}
	resp, err := client.Post("http://"+miraiAddr+"/auth", "application/json; charset=utf-8", bytes.NewBuffer(data))
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
			client:     client,
			miraiAddr:  miraiAddr,
			sessionKey: authResp.SessionKey,
		},nil
	default:
		return nil,errors.New("身份认证失败,Code:"+string(authResp.Code))
	}
}

//校验并激活此会话，并与指定bot绑定
func (this *session) Verify(botqq uint) error  {
	req:=VerifyReq{
		SessionKey: this.sessionKey,
		BotQQ:      botqq,
	}

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := this.client.Post("http://"+this.miraiAddr+"/verify", "application/json; charset=utf-8", bytes.NewReader(data))
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
		this.botqq=botqq
	}
	return err
}

func (this *session) Release() error{
	if this.botqq==0{
		return errors.New("会话未与bot绑定")
	}
	req:=VerifyReq{
		SessionKey: this.sessionKey,
		BotQQ:      this.botqq,
	}
	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := this.client.Post("http://"+this.miraiAddr+"/release", "application/json; charset=utf-8", bytes.NewReader(data))
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