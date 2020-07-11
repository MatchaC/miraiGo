package test

import (
	"fmt"
	"github.com/MatchaC/miraiGo/origin"
	"testing"
)

func TestSendGroupMessage(t *testing.T) {
	messageChain := []origin.Message{origin.MessagePlain{Type: "Plain", Text: "你好(群组消息测试)"}}
	res, err := origin.SendGroupMessage(MiraiHttpApiAddr, SessionKey, TestGroupQQ, messageChain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %+v\n", res)
}
