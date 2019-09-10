package sender

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/urlooker/alarm/cache"
	"github.com/urlooker/web/model"
)

func BuildMail(event *model.Event) string {
	strategy, _ := cache.StrategyMap.Get(event.StrategyId)
	respTime := fmt.Sprintf("%dms", event.RespTime)
	hostName, localIp := getHostNameAndIp()
	return fmt.Sprintf(
		"hostName:%s\nhostIp:%s\nStatus:%s\nUrl:%s\nIp:%s\nRespCode:%s\nRespTime:%s\nTimestamp:%s\nStep:%d\nNote:%s\n",
		hostName,
		localIp,
		event.Status,
		event.Url,
		event.Ip,
		event.RespCode,
		respTime,
		humanTime(event.EventTime),
		event.CurrentStep,
		strategy.Note,
	)
}

func BuildSms(event *model.Event) string {
	respTime := fmt.Sprintf("%dms", event.RespTime)
	return fmt.Sprintf(
		"[%s][%s %s][%s][%s][%s][O%d]",
		event.Status,
		showSubString(event.Url, 50),
		event.Ip,
		event.RespCode,
		respTime,
		humanTime(event.EventTime),
		event.CurrentStep,
	)
}

func BuildDing(event *model.Event) string {
	respTime := fmt.Sprintf("%dms", event.RespTime)
	hostName, localIp := getHostNameAndIp()
	return fmt.Sprintf(
		"hostName:%s\nhostIp:%s\nStatus:%s\nUrl:%s\nIp:%s\nRespCode:%s\nRespTime:%s\nTimestamp:%s\nStep:%d\nNote:%s\n",
		hostName,
		localIp,
		event.Status,
		event.Url,
		event.Ip,
		event.RespCode,
		respTime,
		humanTime(event.EventTime),
		event.CurrentStep,
	)
}

func humanTime(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

func showSubString(str string, length int) string {
	runeStr := []rune(str)
	s := ""
	if length > len(runeStr) {
		length = len(runeStr)
	}

	for i := 0; i < length; i++ {
		s += string(runeStr[i])
	}
	return s
}

func getHostNameAndIp() (string, string) {
	var (
		hostName = ""
		localIp  = ""
		err      error
	)
	hostName, err = os.Hostname()
	if err != nil {
		log.Println("error get local hostname")
	}
	localIp, err = getLocalhostIp()
	if err != nil {
		log.Println("error get local ip")
	}
	return hostName, localIp
}

func getLocalhostIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx], nil
}
