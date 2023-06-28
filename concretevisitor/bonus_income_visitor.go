package concretevisitor

import (
	"fmt"
	entities "visitor-pattern/entity"
)

const STAFF_BONUS_INCOME = 1000000
const SUPERVISOR_BONUS_INCOME = 2000000
const MANAGER_BONUS_INCOME = 5000000

type BonusIncomeVisitor struct{}

func (visitor BonusIncomeVisitor) VisitStaff(staff *entities.Staff) (interface{}, error) {

	var income int

	if staff.PerformancePercentage >= 87 {
		income = STAFF_BONUS_INCOME
	}

	if income != 0 {
		fmt.Printf("Bonus income %d for staff %s \n", income, staff.Name)
	}

	return income, nil
}

func (visitor BonusIncomeVisitor) VisitSupervisor(spv *entities.Supervisor) (interface{}, error) {

	var income int

	if spv.PerformancePercentage >= 85 {
		income = SUPERVISOR_BONUS_INCOME
	}

	if income != 0 {
		fmt.Printf("Bonus income %d for supervisor  %s \n", income, spv.Name)
	}

	return income, nil

}

func (visitor BonusIncomeVisitor) VisitManager(manager *entities.Manager) (interface{}, error) {

	var income int

	if manager.PerformancePercentage >= 80 {
		income = MANAGER_BONUS_INCOME
	}

	if income != 0 {
		fmt.Printf("Bonus income %d for manager  %s \n", income, manager.Name)
	}

	return income, nil

}
