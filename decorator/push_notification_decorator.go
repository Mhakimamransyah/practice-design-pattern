package decorator

import (
	"fmt"
	"visitor-pattern/entity"
)

type PushNotificationDecorator struct{}

func (decorator PushNotificationDecorator) SendMessage(employee entity.Employee, messages string) {
	fmt.Println("Sending notifications by push notification... ")
	fmt.Println(messages)
}

func NewPushNotificationDecorator() *PushNotificationDecorator {
	return &PushNotificationDecorator{}
}
