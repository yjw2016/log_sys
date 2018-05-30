package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"log_sys/logagent/tailf"
	"time"
	"log_sys/logagent/kafka"
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

func sendToKafka(msg *tailf.TextMsg) (e error) {

	e = kafka.Send2kafka(msg.Msg, msg.Topic)
	if e != nil {
		fmt.Println("发送kafka失败")
	}
	//fmt.Printf("read msg:%v,topic:%v\n", msg.Msg, msg.Topic)
	return
}
