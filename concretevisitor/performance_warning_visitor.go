package concretevisitor

import (
	"fmt"
	"time"
	entities "visitor-pattern/entity"
)

const WARNING = "warning"
const STERN_WARNING = "stern-warning"

type PerformanceWarningVisitor struct{}

func (visitor PerformanceWarningVisitor) VisitStaff(staff *entities.Staff) (interface{}, error) {

	var warningType string

	if staff.PerformancePercentage < 30 {
		warningType = STERN_WARNING
	}

	if staff.PerformancePercentage >= 30 && staff.PerformancePercentage < 67 {
		warningType = WARNING
	}

	if warningType != "" {
		fmt.Printf("Notification %s send to staff : %s \n", warningType, staff.Name)
	}

	return warningType, nil
}

func (visitor PerformanceWarningVisitor) VisitSupervisor(spv *entities.Supervisor) (interface{}, error) {

	var warningType string
	durationInYear := (time.Now()).Sub(spv.JoinDate).Hours() / 8760

	if spv.PerformancePercentage < 30 {
		warningType = STERN_WARNING
	}

	if durationInYear >= 10 && (spv.PerformancePercentage >= 30 && spv.PerformancePercentage < 60) {
		warningType = WARNING
	}

	if durationInYear < 10 && (spv.PerformancePercentage >= 30 && spv.PerformancePercentage < 70) {
		warningType = WARNING
	}

	if warningType != "" {
		fmt.Printf("Notification %s send to supervisor : %s \n", warningType, spv.Name)
	}

	return warningType, nil
}

func (visitor PerformanceWarningVisitor) VisitManager(manager *entities.Manager) (interface{}, error) {

	var warningType string
	Age := (time.Now()).Sub(manager.BirthDate).Hours() / 8760

	if manager.PerformancePercentage < 40 {
		warningType = STERN_WARNING
	}

	if Age >= 40 && (manager.PerformancePercentage >= 40 && manager.PerformancePercentage < 60) {
		warningType = WARNING
	}

	if Age < 40 && (manager.PerformancePercentage >= 40 && manager.PerformancePercentage < 70) {
		warningType = WARNING
	}

	if warningType != "" {
		fmt.Printf("Notification %s send to manager : %s \n", warningType, manager.Name)
	}

	return warningType, nil
}
