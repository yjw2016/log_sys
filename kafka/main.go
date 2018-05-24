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

/*
运行，
先启动zk & kafka
zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties & kafka-server-start /usr/local/etc/kafka/server.properties
使用kafka消费者消费队列中的消息,注意地址端口参数为zk的，不是生产者的
kafka-console-consumer --topic nginx_log --zookeeper 127.0.0.1 2181
测试生产者
*/
