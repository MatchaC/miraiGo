package test

import (
	"fmt"
	"github.com/MatchaC/miraiGo/origin"
	"testing"
)

func TestSendFriendMessage(t *testing.T) {
	messageChain := []origin.Message{origin.MessagePlain{Type: "Plain", Text: "你好(好友消息测试)"}}
	res, err := origin.SendFriendMessage(MiraiHttpApiAddr, SessionKey, TestTargetQQ, messageChain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %+v\n", res)
}
