package miraiGo

type Group struct {
	Id uint `json:"id"` //消息来源群号
	Name string `json:"name"` //消息来源群名
	Permisson string `json:"permisson"` //bot在群中的角色
}

type GroupSender struct {
	Id uint `json:"id"` //发送者QQ号
	MemberName string `json:"memberName"` //发送者群昵称
	Permission string `json:"permission"` //发送者在群中的角色
	Group Group `json:"group"` //消息来源群信息
}