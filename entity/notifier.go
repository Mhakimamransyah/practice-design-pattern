package entity

type Notifier interface {
	SendMessage(employee Employee, messages string)
}
