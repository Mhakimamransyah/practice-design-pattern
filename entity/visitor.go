package entity

type Visitor interface {
	VisitStaff(staff *Staff) (interface{}, error)
	VisitSupervisor(spv *Supervisor) (interface{}, error)
	VisitManager(manager *Manager) (interface{}, error)
}
