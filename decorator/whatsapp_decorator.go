package decorator

import (
	"fmt"

	"github.com/Mhakimamransyah/practice-design-pattern/entity"
)

type WhatsAppDecorator struct {
	Wrappee entity.Notifier
}

func (decorator WhatsAppDecorator) SendMessage(employee entity.Employee, messages string) {
	fmt.Println("Sending notifications by whatsapp...")
	decorator.Wrappee.SendMessage(employee, messages)
}

func NewWhatsAppDecorator(Wrappee entity.Notifier) *WhatsAppDecorator {
	return &WhatsAppDecorator{Wrappee}
}
