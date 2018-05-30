package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"log_sys/logagent/tailf"
	"github.com/astaxie/beego/logs"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath  string
	chanSize int
	kafkaAddr string
	collectConf []tailf.CollectConf
}

func loadCollectConf(conf config.Configer) (err error) {

	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		errors.New("invalid collect::log_path")
		return
	}
	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		errors.New("invalid collect::topic")
		return
	}

	//fmt.Printf("cc is %v\n", cc)
	appConfig.collectConf = append(appConfig.collectConf, cc)

	return
}

func loadConf(confType, filename string) (err error) {

	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed,err : ", err)
		return
	}

	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}
	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "./logs"
	}
	appConfig.chanSize, err = conf.Int("logs::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}
	appConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.kafkaAddr) == 0 {
		logs.Error("invalid kafka address")
		appConfig.kafkaAddr = "192.168.14.7:9092"
	}
	err = loadCollectConf(conf)
	if err != nil {
		fmt.Println("load collect conf failed :%v\n", err)
	}

	return
}
