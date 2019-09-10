package sender

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/urlooker/alarm/g"
)

func LPUSH(queue, message string) {
	rc := g.RedisConnPool.Get()
	defer rc.Close()
	_, err := rc.Do("LPUSH", queue, message)
	if err != nil {
		log.Println("LPUSH redis", queue, "fail:", err, "message:", message)
	}
}

func WriteSmsModel(sms *g.Sms) {
	if sms == nil {
		return
	}

	bs, err := json.Marshal(sms)
	if err != nil {
		log.Println(err)
		return
	}

	LPUSH(g.Config.Queue.Sms, string(bs))
}

func WriteMailModel(mail *g.Mail) {
	if mail == nil {
		return
	}

	bs, err := json.Marshal(mail)
	if err != nil {
		log.Println(err)
		return
	}

	LPUSH(g.Config.Queue.Mail, string(bs))
}

func WriteDingModel(ding *g.Ding) {
	if ding == nil {
		return
	}

	bs, err := json.Marshal(ding)
	if err != nil {
		log.Println(err)
		return
	}

	LPUSH(g.Config.Queue.Ding, string(bs))
}

func WriteSms(tos []string, content string) {
	if len(tos) == 0 {
		return
	}
	log.Println("Write content to sms")
	sms := &g.Sms{Tos: strings.Join(tos, ","), Content: content}
	WriteSmsModel(sms)
}

func WriteMail(tos []string, subject, content string) {
	if len(tos) == 0 {
		return
	}
	log.Println("Write content to mail")
	mail := &g.Mail{Tos: strings.Join(tos, ","), Subject: subject, Content: content}
	WriteMailModel(mail)
}

func WriteDing(content string) {
	log.Println("Write content to ding")
	ding := &g.Ding{Content: content}
	WriteDingModel(ding)
}
