package entity

type Employee interface {
	Accept(v Visitor)
}
