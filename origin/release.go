package origin

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type RespRelease struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Release(remoteAddr string, sessionKey string, qq int) (RespRelease, error) {
	reqRelease := make(map[string]interface{})
	reqRelease["sessionKey"] = sessionKey
	reqRelease["qq"] = qq
	reqReleaseBody, err := jsoniter.Marshal(&reqRelease)

	resp, err := http.DefaultClient.Post("http://"+remoteAddr+"/release", "application/json; charset=utf-8", bytes.NewReader(reqReleaseBody))
	if err != nil {
		return RespRelease{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespRelease{}, err
	}

	var respRelease RespRelease
	err = jsoniter.Unmarshal(body, &respRelease)
	if err != nil {
		return RespRelease{}, err
	}

	err = CheckErrorCode(respRelease.Code)
	return respRelease, err
}
