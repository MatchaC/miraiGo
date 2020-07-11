package origin

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type RespAbout struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
	Data         struct {
		Version string `json:"version"`
	} `json:"data"`
}

func About(remoteAddr string) (RespAbout, error) {
	resp, err := http.DefaultClient.Get("http://" + remoteAddr + "/about")
	if err != nil {
		return RespAbout{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespAbout{}, err
	}

	var respAbout RespAbout
	err = jsoniter.Unmarshal(body, &respAbout)
	if err != nil {
		return RespAbout{}, err
	}

	err = CheckErrorCode(respAbout.Code)
	return respAbout, err
}
