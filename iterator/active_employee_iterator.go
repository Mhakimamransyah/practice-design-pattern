package iterator

import "github.com/Mhakimamransyah/practice-design-pattern/entity"

type ActiveEmployeeIterator struct {
	Employees       []entity.Employee
	CurrentPosition int
	Status          bool
}

func (iter *ActiveEmployeeIterator) GetNext() entity.Employee {

	if !iter.HasMore() {
		return nil
	}

	employee := iter.Employees[iter.CurrentPosition]

	iter.CurrentPosition++

	if employee.GetStatus() {
		return employee
	}

	return iter.GetNext()

}
func (iter *ActiveEmployeeIterator) HasMore() bool {

	if iter.CurrentPosition >= len(iter.Employees) {
		return false
	}

	employee := iter.Employees[iter.CurrentPosition]

	if employee.GetStatus() {
		return iter.CurrentPosition < len(iter.Employees)
	}

	iter.CurrentPosition++

	return iter.HasMore()

}

func NewActiveEmployeeIterator(employees []entity.Employee, status bool) *ActiveEmployeeIterator {
	return &ActiveEmployeeIterator{Employees: employees, Status: status}
}
