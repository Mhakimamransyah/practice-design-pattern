package entity

import "time"

type Manager struct {
	Id                    int
	Name                  string
	JoinDate              time.Time
	BirthDate             time.Time
	PerformancePercentage float32
}

func (manager Manager) Accept(visitor Visitor) {
	visitor.VisitManager(&manager)
}
