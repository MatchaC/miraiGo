package test

import (
	"fmt"
	"github.com/MatchaC/miraiGo/origin"
	"testing"
)

func TestRelease(t *testing.T) {
	res, err := origin.Release(MiraiHttpApiAddr, SessionKey, QQ)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %+v\n", res)
}
