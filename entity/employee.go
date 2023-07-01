package entity

type Employee interface {
	Accept(v Visitor)
	GetName() string
	GetWhatsAppNumber() int
	GetEmail() string
	GetTokenNotifications() string
}
