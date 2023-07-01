package factory

import (
	"github.com/Mhakimamransyah/practice-design-pattern/decorator"
	"github.com/Mhakimamransyah/practice-design-pattern/entity"
)

type NotificatorContract interface {
	StaffNotifier() entity.Notifier
	SupervisorNotifier() entity.Notifier
	ManagerNotifier() entity.Notifier
}

type Notificator struct{}

func (notifier *Notificator) StaffNotifier() entity.Notifier {
	return decorator.NewPushNotificationDecorator()
}

func (notifier *Notificator) SupervisorNotifier() entity.Notifier {
	pushNotification := decorator.NewPushNotificationDecorator()
	emailNotification := decorator.NewEmailDecorator(pushNotification)
	return emailNotification
}

func (notifier *Notificator) ManagerNotifier() entity.Notifier {
	pushNotification := decorator.NewPushNotificationDecorator()
	emailNotification := decorator.NewEmailDecorator(pushNotification)
	whatsAppNotification := decorator.NewWhatsAppDecorator(emailNotification)
	return whatsAppNotification
}

func NewNotificator() *Notificator {
	return &Notificator{}
}
