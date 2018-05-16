package main

import (
  "fmt"
  "time"
  "strings"
  "flag"
  "encoding/json"
  "github.com/confluentinc/confluent-kafka-go/kafka"
)

var viewed int
var interacted int
var clickThrough int

type Message struct {
	Type string `json:"type"`
	Data
}

type Data struct {
	ViewId string      `json:"viewId"`
	EventDateTime string `json:"eventDateTime"`
}

func init()  {
  viewed = 0
  interacted = 0
  clickThrough = 0
}

func main() {
  topic := flag.String("topic","/data/","Where to write batch files")
  flag.Parse()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka",
		"group.id":          "counterApp",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{*topic, "^aRegex.*[Tt]opic"}, nil)

  go func ()  {
    for {
      fmt.Printf("Viewed: %d,\nInteracted: %d,\nClick-Through: %d\n",viewed,interacted,clickThrough)
      fmt.Println("\n-----------------------------------------------------------------------------\n")
      time.Sleep(5 * time.Second)
    }
  }()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
      var message Message
      jsonErr := json.Unmarshal(msg.Value,&message)
      if jsonErr == nil {
        switch strings.TrimSpace(strings.ToLower(message.Type)) {
        case "viewed":
          viewed++
        case "click-through":
          clickThrough++
        case "interacted":
          interacted++
        default:
          fmt.Printf("Unknown event type: %s\n",message.Type)
        }
      }
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
