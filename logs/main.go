package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logs/logcollect.log"
	config["level"] = logs.LevelDebug

	configString, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed,err:",err)
		return
	}

	logs.SetLogger(logs.AdapterFile,string(configString))

	logs.Debug("this is a test,my name is %s","stu01")
	logs.Trace("this is a test,my name is %s","stu02")
	logs.Warn("this is a test,my name is %s","stu03")
}
