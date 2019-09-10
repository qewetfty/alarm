package sender

import (
	"bytes"
	"encoding/json"
	"github.com/urlooker/alarm/g"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DingMsg struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

type DingResponse struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
}

func ConsumeDing() {
	queue := g.Config.Queue.Ding
	for {
		L := PopAllDing(queue)
		if len(L) == 0 {
			time.Sleep(200 * time.Millisecond)
			continue
		}
		SendDingList(L)
	}
}

func SendDingList(L []*g.Ding) {
	log.Println("Message to send: ", len(L))
	for _, ding := range L {
		if ding.Content == "" {
			continue
		}

		DingWorkerChan <- 1
		go SendDing(ding)
	}
}

func SendDing(ding *g.Ding) {
	defer func() {
		<-DingWorkerChan
	}()

	dingMsg := new(DingMsg)
	dingMsg.MsgType = "text"
	text := new(Text)
	text.Content = ding.Content
	dingMsg.Text = *text

	sendText, err := json.Marshal(dingMsg)
	if err != nil {
		log.Println("json marshal error", err, "sendText: ", sendText)
		return
	}

	token := g.Config.Ding.Token
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + token
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(sendText))
	if err != nil {
		log.Println("Dingding post err.", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read response err.", err)
		return
	}
	var response DingResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("json unmarchal error.", err)
		return
	}
	if response.ErrorCode == 0 {
		log.Println("Ding msg success")
	} else {
		log.Println("Ding msg failed. errMsg: ", response.ErrorMsg)
	}
}
