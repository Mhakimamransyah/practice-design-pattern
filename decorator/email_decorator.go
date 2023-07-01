package decorator

import (
	"fmt"
	"visitor-pattern/entity"
)

type EmailDecorator struct {
	Wrappee entity.Notifier
}

func (decorator EmailDecorator) SendMessage(employee entity.Employee, messages string) {
	fmt.Println("Sending notifications by email...")
	decorator.Wrappee.SendMessage(employee, messages)
}

func NewEmailDecorator(Wrappee entity.Notifier) *EmailDecorator {
	return &EmailDecorator{Wrappee}
}
