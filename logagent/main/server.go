package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"log_sys/logagent/tailf"
	"time"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed")
			time.Sleep(time.Second)
			continue
		}
	}

	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {

	fmt.Printf("read msg:%v,topic:%v\n", msg.Msg, msg.Topic)
	return
}
