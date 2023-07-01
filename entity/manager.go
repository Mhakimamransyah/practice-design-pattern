package entity

import "time"

type Manager struct {
	Id                    int
	Name                  string
	JoinDate              time.Time
	BirthDate             time.Time
	PerformancePercentage float32
	WhatsAppNumber        int
	Email                 string
	TokenNotifications    string
}

func (manager Manager) Accept(visitor Visitor) {
	visitor.VisitManager(&manager)
}

func (manager Manager) GetName() string {
	return manager.Name
}

func (manager Manager) GetWhatsAppNumber() int {
	return manager.WhatsAppNumber
}

func (manager Manager) GetEmail() string {
	return manager.Email
}

func (manager Manager) GetTokenNotifications() string {
	return manager.TokenNotifications
}
