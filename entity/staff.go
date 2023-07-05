package entity

import (
	"time"
)

type Staff struct {
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

func (staff Staff) Accept(v Visitor) {
	v.VisitStaff(&staff)
}

func (staff Staff) GetName() string {
	return staff.Name
}

func (staff Staff) GetWhatsAppNumber() int {
	return staff.WhatsAppNumber
}

func (staff Staff) GetEmail() string {
	return staff.Email
}

func (staff Staff) GetTokenNotifications() string {
	return staff.TokenNotifications
}

func (staff Staff) GetStatus() bool {
	return staff.Status
}
