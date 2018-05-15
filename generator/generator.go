package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
  "flag"
  "math/rand"
  "time"
  "crypto/md5"
  "encoding/hex"
  "encoding/json"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka"})
	if err != nil {
		panic(err)
	}

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
	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value: []byte(word),
		}, nil)
	}

	// Wait for message deliveries
	p.Flush(15 * 1000)
}

func getEvents() []string {
  rand.Seed(time.Now().UTC().UnixNano())
  hash := getMD5Hash(time.Now().UTC().String())
  events := []string{"viewed"}
  //var messages []Message
  ev := rand.Intn(100)

  switch {
  case ev > 94:
    events = append(events,"interacted","click-Through")
	break
  case ev > 89:
    events = append(events,"click-Through")
    break
  case ev > 84:
    events = append(events,"interacted")
  }

  for _,event := range events{
	time.Sleep(time.Second / 100)

	message, _ := json.Marshal(Message{
		Type: event,
		Data: Data{
			ViewId: hash,
			EventDateTime: time.Now().UTC().String(),
		},
	})

	fmt.Println(string(message))
  }

  
}

func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
