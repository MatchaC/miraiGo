package miraiGo

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

//发送好友消息
func (this *Session) SendFriendMessage(req RequestSendFriendMessage) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}
	req.SessionKey=this.sessionKey

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

	fmt.Println("msg:",uniResp.Msg)
	err = GetErrByCode(uniResp.Code)
	return err
}

//发送临时消息
func (this *Session) SendTempMessage(req RequestSendTempMessage) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}
	req.SessionKey=this.sessionKey

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

	fmt.Println("msg:",uniResp.Msg)
	err = GetErrByCode(uniResp.Code)
	return err
}

//发送群消息
func (this *Session) SendGroupMessage(req RequestSendGroupMessage) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}
	req.SessionKey=this.sessionKey

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/sendGroupMessage", "application/json; charset=utf-8", bytes.NewReader(data))
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

	fmt.Println("msg:",uniResp.Msg)
	err = GetErrByCode(uniResp.Code)
	return err
}

//发送图片消息(通过URL)
func (this *Session) SendImageMessage(req RequestSendImageMessage) ([]string,error){
	if this.qq==0{
		return nil,errors.New("会话未与bot绑定")
	}
	req.SessionKey=this.sessionKey

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return nil,err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/sendImageMessage", "application/json; charset=utf-8", bytes.NewReader(data))
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil,err
	}

	var imglist []string
	err = jsoniter.Unmarshal(body, &imglist)
	if err!=nil{
		return nil,err
	}

	return imglist,nil
}

//上传图片到服务器，暂无


//撤回消息
func (this *Session) ReCall(req RequestReCall) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}
	req.SessionKey=this.sessionKey

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/recall", "application/json; charset=utf-8", bytes.NewReader(data))
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

func (this *Session) HandleEventNewFriendRequest(req ResponseEventInvited) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}
	req.SessionKey=this.sessionKey

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/resp/newFriendRequestEvent", "application/json; charset=utf-8", bytes.NewReader(data))
	if err!=nil{
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (this *Session) HandleEventBotInvitedJoinGroupRequest(req ResponseEventInvited) error{
	if this.qq==0{
		return errors.New("会话未与bot绑定")
	}
	req.SessionKey=this.sessionKey

	data,err:=jsoniter.Marshal(&req)
	if err!=nil{
		return err
	}

	resp, err := http.DefaultClient.Post("http://"+this.miraiAddr+"/resp/memberJoinRequestEvent", "application/json; charset=utf-8", bytes.NewReader(data))
	if err!=nil{
		return err
	}
	defer resp.Body.Close()

	return nil
}