package origin

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type RespAuth struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

func Auth(remoteAddr string, authKey string) (RespAuth, error) {
	reqAuth := make(map[string]interface{})
	reqAuth["authKey"] = authKey
	reqAuthBody, err := jsoniter.Marshal(&reqAuth)

	resp, err := http.DefaultClient.Post("http://"+remoteAddr+"/auth", "application/json; charset=utf-8", bytes.NewReader(reqAuthBody))
	if err != nil {
		return RespAuth{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespAuth{}, err
	}

	var respAuth RespAuth
	err = jsoniter.Unmarshal(body, &respAuth)
	if err != nil {
		return RespAuth{}, err
	}

	err = CheckErrorCode(respAuth.Code)
	return respAuth, err
}
