package test

import (
	"fmt"
	"github.com/MatchaC/miraiGo/origin"
	"testing"
)

func TestAuth(t *testing.T) {
	res, err := origin.Auth(MiraiHttpApiAddr, AuthKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %+v\n", res)
}
