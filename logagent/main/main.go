package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"log_sys/logagent/tailf"
	"time"
)

func main() {
	filename := "./conf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Println("load config failed : %v\n", err)
		panic("load config failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Println("load logger failed,err: %v\n", err)
		panic("load logger failed")
		return
	}

	logs.Debug("load conf succ ,config:%v", appConfig)

	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed ,err:%v", err)
		return
	}

	logs.Debug("initialize all succ")

	go func() {
		for i:=0;i<60;i++ {
			logs.Debug("test for logger %d,", i)
			time.Sleep(time.Millisecond*800)
		}
	}()

	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed")
		return
	}

	logs.Info("program exited")
}
