package factory

import (
	"github.com/Mhakimamransyah/practice-design-pattern/entity"

	"github.com/stretchr/testify/mock"
)

type NotifierMock struct {
	Mock mock.Mock
}

func (notifier *NotifierMock) StaffNotifier() entity.Notifier {
	arguments := notifier.Mock.Called()
	return arguments.Get(0).(entity.Notifier)
}

func (notifier *NotifierMock) SupervisorNotifier() entity.Notifier {
	arguments := notifier.Mock.Called()
	return arguments.Get(0).(entity.Notifier)
}

func (notifier *NotifierMock) ManagerNotifier() entity.Notifier {
	arguments := notifier.Mock.Called()
	return arguments.Get(0).(entity.Notifier)
}

func NewNotifierMock() *NotifierMock {
	return &NotifierMock{}
}
