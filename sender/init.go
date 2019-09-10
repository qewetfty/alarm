package sender

import (
	"github.com/urlooker/alarm/g"
)

var (
	SmsWorkerChan  chan int
	MailWorkerChan chan int
	DingWorkerChan chan int
)

func Init() {
	workerConfig := g.Config.Worker
	SmsWorkerChan = make(chan int, workerConfig.Sms)
	MailWorkerChan = make(chan int, workerConfig.Mail)
	DingWorkerChan = make(chan int, workerConfig.Ding)

	Consume()
}

func Consume() {
	go ConsumeMail()
	go ConsumeSms()
	go ConsumeDing()
}
