package entity

type Iterator interface {
	GetNext() Employee
	HasMore() bool
}
