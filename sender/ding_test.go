package sender

import (
	"fmt"
	"github.com/urlooker/alarm/g"
	"os"
	"testing"
)

func TestSendDing(t *testing.T) {
	ding := new(g.Ding)
	ding.Content = "测试钉钉"
	SendDing(ding)
}

func TestConsume(t *testing.T) {
	name, _ := os.Hostname()
	fmt.Println(name)
}
