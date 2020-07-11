package origin

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

func SendFriendMessage(remoteAddr string, sessionKey string, targetId int, messageChain []Message) (RespSendMessage, error) {
	reqSendFriendMessage := make(map[string]interface{})
	reqSendFriendMessage["sessionKey"] = sessionKey
	reqSendFriendMessage["target"] = targetId
	reqSendFriendMessage["messageChain"] = messageChain
	reqSendFriendMessageBody, err := jsoniter.Marshal(&reqSendFriendMessage)

	resp, err := http.DefaultClient.Post("http://"+remoteAddr+"/sendFriendMessage", "application/json; charset=utf-8", bytes.NewReader(reqSendFriendMessageBody))
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
