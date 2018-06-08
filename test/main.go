package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	fmt.Println("real main")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
	}
	fmt.Println("real main end")
	return 0
}
