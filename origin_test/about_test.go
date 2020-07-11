package test

import (
	"fmt"
	"github.com/MatchaC/miraiGo/origin"
	"testing"
)

func TestAbout(t *testing.T) {
	about, err := origin.About(MiraiHttpApiAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %+v\n", about)
}
