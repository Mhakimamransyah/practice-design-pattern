package entity

import "time"

type Staff struct {
	Id                    int
	Name                  string
	JoinDate              time.Time
	BirthDate             time.Time
	PerformancePercentage float32
}

func (staff Staff) Accept(v Visitor) {
	v.VisitStaff(&staff)
}
