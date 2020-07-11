package test

import (
	"fmt"
	"github.com/MatchaC/miraiGo/origin"
	"testing"
)

func TestAuth(t *testing.T) {
	res, err := origin.Auth(MiraiHttpApiAddr, "1qaz2wsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %+v\n", res)
}
