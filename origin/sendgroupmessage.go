package origin

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

func SendGroupMessage(remoteAddr string, sessionKey string, targetGroupId int, messageChain []Message) (RespSendMessage, error) {
	reqSendGroupMessage := make(map[string]interface{})
	reqSendGroupMessage["sessionKey"] = sessionKey
	reqSendGroupMessage["target"] = targetGroupId
	reqSendGroupMessage["messageChain"] = messageChain
	reqSendGroupMessageBody, err := jsoniter.Marshal(&reqSendGroupMessage)

	resp, err := http.DefaultClient.Post("http://"+remoteAddr+"/sendGroupMessage", "application/json; charset=utf-8", bytes.NewReader(reqSendGroupMessageBody))
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
