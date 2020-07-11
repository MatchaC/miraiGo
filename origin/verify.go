package origin

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type RespVerify struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Verify(remoteAddr string, sessionKey string, qq int) (RespVerify, error) {
	reqVerify := make(map[string]interface{})
	reqVerify["sessionKey"] = sessionKey
	reqVerify["qq"] = qq
	reqVerifyBody, err := jsoniter.Marshal(&reqVerify)

	resp, err := http.DefaultClient.Post("http://"+remoteAddr+"/verify", "application/json; charset=utf-8", bytes.NewReader(reqVerifyBody))
	if err != nil {
		return RespVerify{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespVerify{}, err
	}

	var respVerify RespVerify
	err = jsoniter.Unmarshal(body, &respVerify)
	if err != nil {
		return RespVerify{}, err
	}

	err = CheckErrorCode(respVerify.Code)
	return respVerify, err
}
