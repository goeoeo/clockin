package test

import (
	"github.com/phpdi/clockin/core"
	"testing"
)

func TestMail(t *testing.T)  {

	err:=core.SendMail("打卡通知","/home/yu/code/clockin/data/screen.png")
	if err != nil {
		t.Fatal(err)
	}
}