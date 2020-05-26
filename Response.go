package miraiGo

import "errors"

//身份验证响应
type AuthResp struct {
	Code uint `json:"code"`
	SessionKey string `json:"session"`
}

//统一响应
type UniResp struct {
	Code uint `json:"code"`
	Msg string `json:"msg"`
}

func GetErrByCode(code uint) error {
	switch code {
	case 0:
		return nil
	case 1:
		return errors.New("错误的auth key")
	case 2:
		return errors.New("指定的Bot不存在")
	case 3:
		return errors.New("Session失效或不存在")
	case 4:
		return errors.New("Session未认证(未激活)")
	case 5:
		return errors.New("发送消息目标不存在(指定对象不存在)")
	case 6:
		return errors.New("指定文件不存在，出现于发送本地图片")
	case 10:
		return errors.New("无操作权限，指Bot没有对应操作的限权")
	case 20:
		return errors.New("Bot被禁言，指Bot当前无法向指定群发送消息")
	case 30:
		return errors.New("消息过长")
	case 400:
		return errors.New("错误的访问，如参数错误等")
	default:
		return errors.New("未知错误,Code:"+string(code))
	}
}