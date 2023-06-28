package entity

import "time"

type Supervisor struct {
	Id                    int
	Name                  string
	JoinDate              time.Time
	BirthDate             time.Time
	PerformancePercentage float32
}

func (spv Supervisor) Accept(visitor Visitor) {
	visitor.VisitSupervisor(&spv)
}
