package entity

import (
	"time"
)

type Supervisor struct {
	Id                    int
	Name                  string
	JoinDate              time.Time
	BirthDate             time.Time
	PerformancePercentage float32
	WhatsAppNumber        int
	Email                 string
	TokenNotifications    string
	Status                bool
}

func (spv Supervisor) Accept(visitor Visitor) {
	visitor.VisitSupervisor(&spv)
}

func (spv Supervisor) GetName() string {
	return spv.Name
}

func (spv Supervisor) GetWhatsAppNumber() int {
	return spv.WhatsAppNumber
}

func (spv Supervisor) GetEmail() string {
	return spv.Email
}

func (spv Supervisor) GetTokenNotifications() string {
	return spv.TokenNotifications
}

func (spv Supervisor) GetStatus() bool {
	return spv.Status
}
