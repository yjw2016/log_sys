package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"fmt"
)

var (
	producer sarama.SyncProducer
)

func InitKafka(addr string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer([]string{addr}, config)

	if err != nil {
		logs.Error("init kafka producer failed,err:", err)
		return
	}

	logs.Debug("init kafka succ")
	return
}
func Send2kafka(data, topic string) (err error) {
	fmt.Println(data)

	message := &sarama.ProducerMessage{}
	message.Topic = topic
	message.Value = sarama.StringEncoder(data)

	partition, offset, err := producer.SendMessage(message)

	if err != nil {
		logs.Error("send message failed,err:%v data:%v topic:%v", err, data, topic)
		return
	}
	fmt.Printf("send succ,parition:%v offset:%v,topic:%v\n", partition, offset, topic)
	//logs.Debug("send succ,parition:%v offset:%v,topic:%v\n", partition, offset, topic)
	return

}
