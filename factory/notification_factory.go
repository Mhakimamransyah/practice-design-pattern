package factory

import (
	"visitor-pattern/decorator"
	"visitor-pattern/entity"
)

func StaffNotifier() entity.Notifier {
	return decorator.NewPushNotificationDecorator()
}

func SupervisorNotifier() entity.Notifier {
	pushNotification := decorator.NewPushNotificationDecorator()
	emailNotification := decorator.NewEmailDecorator(pushNotification)
	return emailNotification
}

func ManagerNotifier() entity.Notifier {
	pushNotification := decorator.NewPushNotificationDecorator()
	emailNotification := decorator.NewEmailDecorator(pushNotification)
	whatsAppNotification := decorator.NewWhatsAppDecorator(emailNotification)
	return whatsAppNotification
}
