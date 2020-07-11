package test

import (
	"fmt"
	"github.com/MatchaC/miraiGo/origin"
	"testing"
)

func TestSendTempMessage(t *testing.T) {
	messageChain := []origin.Message{origin.MessagePlain{Type: "Plain", Text: "你好(临时消息测试)"}}
	res, err := origin.SendTempMessage(MiraiHttpApiAddr, SessionKey, TestGroupQQ, TestTargetQQ, messageChain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %+v\n", res)
}
