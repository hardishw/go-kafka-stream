package generator

import (
  "math/rand"
  "time"
  "crypto/md5"
  "encoding/hex"
  "encoding/json"
)

type Message struct {
	Type string `json:"type"`
	Data
}

type Data struct {
	ViewId string      `json:"viewId"`
	EventDateTime string `json:"eventDateTime"`
}

func CreateEvents(evntGrps int) []string {
	var messages []string
	for i := 0; i < evntGrps; i++ {
		messages = append(messages,getEvents()...)
	}

	return messages
}

func getEvents() []string {
	currentTime := time.Now().UTC()
  rand.Seed(currentTime.UnixNano())
  hash := getMD5Hash(currentTime.String())
  events := []string{"viewed"}
  var messages []string
  random := rand.Intn(100)

  switch {
  case random > 94:
    events = append(events,"interacted","click-Through")
	break
  case random > 89:
    events = append(events,"click-Through")
    break
  case random > 84:
    events = append(events,"interacted")
  }

  for _, event := range events{
	time.Sleep(time.Second / 100)

	message, _ := json.Marshal(Message{
		Type: event,
		Data: Data{
			ViewId: hash,
			EventDateTime: time.Now().UTC().String(),
		},
	})

	messages = append(messages,string(message))
  }

  return messages
}

func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
