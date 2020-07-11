package origin

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

func SendTempMessage(remoteAddr string, sessionKey string, targetGroupId int, targetId int, messageChain []Message) (RespSendMessage, error) {
	reqSendTempMessage := make(map[string]interface{})
	reqSendTempMessage["sessionKey"] = sessionKey
	reqSendTempMessage["group"] = targetGroupId
	reqSendTempMessage["qq"] = targetId
	reqSendTempMessage["messageChain"] = messageChain
	reqSendTempMessageBody, err := jsoniter.Marshal(&reqSendTempMessage)

	resp, err := http.DefaultClient.Post("http://"+remoteAddr+"/sendTempMessage", "application/json; charset=utf-8", bytes.NewReader(reqSendTempMessageBody))
	if err != nil {
		return RespSendMessage{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespSendMessage{}, err
	}

	var respSendMessage RespSendMessage
	err = jsoniter.Unmarshal(body, &respSendMessage)
	if err != nil {
		return RespSendMessage{}, err
	}

	err = CheckErrorCode(respSendMessage.Code)
	return respSendMessage, err
}
