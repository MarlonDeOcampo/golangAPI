package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(messages string, topicName string) {
	fmt.Printf("Starting producer...")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := topicName
	data := []byte(messages)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

//Send data to Kafka topic
// type Data struct {
// 	Id          int
// 	Message     string
// 	Source      string
// 	Destination string
// }

// func main() {
// 	topicName := "myTopic"

// 	data := []Data{
// 		{Id: 2, Message: "World", Source: "1", Destination: "A"},
// 		{Id: 2, Message: "Earth", Source: "1", Destination: "B"},
// 		{Id: 2, Message: "Planets", Source: "2", Destination: "C"},
// 	}
// 	stringJson, _ := json.Marshal(data)

// 	fmt.Println(string(stringJson))

// 	produce(string(stringJson), topicName)
// }
