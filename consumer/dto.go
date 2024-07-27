package consumer

import "fmt"

type NotificationMessage struct {
	Topic string
	Key   string
}

func GetMessages() []NotificationMessage {
	var messages []NotificationMessage = make([]NotificationMessage, 100)

	for i := 0; i < 1000; i++ {
		key := i % 10
		messages = append(messages, NotificationMessage{
			Topic: fmt.Sprintf("topic-%d", key),
			Key:   fmt.Sprintf("key-%d", key),
		})
	}
	return messages
}
