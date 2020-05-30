package miraiGo

import (
	"bytes"
	"errors"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

//发送好友消息
func (this *session) SendFriendMessage(req RequestSendFriendMessage) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/sendFriendMessage", "application/json; charset=utf-8", bytes.NewReader(data))
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

//发送临时消息
func (this *session) SendTempMessage(req RequestSendTempMessage) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/sendTempMessage", "application/json; charset=utf-8", bytes.NewReader(data))
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

//发送群消息
func (this *session) SendGroupMessage(req RequestSendTempMessage) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/sendTempMessage", "application/json; charset=utf-8", bytes.NewReader(data))
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