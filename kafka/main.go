package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"192.168.1.102:9092"}, config)
	if err != nil {
		fmt.Println("producer close , err : ", err)
		return
	}

	defer client.Close()
	for i := 0; i < 10; i++ {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		content := fmt.Sprintf("this is a good test , my message is cool , No. %d", i)
		//fmt.Println(content)
		msg.Value = sarama.StringEncoder(content)

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}

		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(10 * time.Millisecond)
	}
}
