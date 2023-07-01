package concretevisitor

import (
	"testing"
	"time"
	"visitor-pattern/entity"
	"visitor-pattern/factory"

	"github.com/stretchr/testify/assert"
)

var notifier *factory.NotifierMock
var performanceWarningVisitor PerformanceWarningVisitor

type MockNotifierDecorator struct {
}

func (mock MockNotifierDecorator) SendMessage(employee entity.Employee, messages string) {
	//mocking notifications library
}
func NewMockNotifierDecorator() *MockNotifierDecorator {
	return &MockNotifierDecorator{}
}

func TestVisitStaffWarningNotification(t *testing.T) {

	notifier.Mock.On("StaffNotifier").Return(NewMockNotifierDecorator())

	testSuites := []struct {
		name     string
		input    entity.Staff
		expected interface{}
	}{
		{
			name: "Test Staff Warning Notifications",
			input: entity.Staff{
				Id:                    1111,
				Name:                  "TestSuites 1",
				PerformancePercentage: 35,
			},
			expected: WARNING,
		},
		{
			name: "Test Staff Stern-warning Notifications",
			input: entity.Staff{
				Id:                    2222,
				Name:                  "TestSuites 2",
				PerformancePercentage: 20,
			},
			expected: STERN_WARNING,
		},
	}

	for _, test := range testSuites {
		t.Run(test.name, func(t *testing.T) {
			result, _ := performanceWarningVisitor.VisitStaff(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestVisitSupervisorWarningNotification(t *testing.T) {

	notifier.Mock.On("SupervisorNotifier").Return(NewMockNotifierDecorator())

	testSuites := []struct {
		name     string
		input    entity.Supervisor
		expected interface{}
	}{
		{
			name: "Test Supervisor Stern-Warning Notifications",
			input: entity.Supervisor{
				Id:                    1111,
				Name:                  "TestSuites 1",
				JoinDate:              time.Date(2013, time.December, 15, 0, 0, 0, 0, time.UTC),
				PerformancePercentage: 13,
			},
			expected: STERN_WARNING,
		},
		{
			name: "Test Staff Warning Notifications",
			input: entity.Supervisor{
				Id:                    2222,
				Name:                  "TestSuites 2",
				JoinDate:              time.Date(2013, time.January, 01, 0, 0, 0, 0, time.UTC),
				PerformancePercentage: 35,
			},
			expected: WARNING,
		},
	}

	for _, test := range testSuites {
		t.Run(test.name, func(t *testing.T) {
			result, _ := performanceWarningVisitor.VisitSupervisor(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestVisitManagerWarningNotification(t *testing.T) {

	notifier.Mock.On("ManagerNotifier").Return(NewMockNotifierDecorator())

	testSuites := []struct {
		name     string
		input    entity.Manager
		expected interface{}
	}{
		{
			name: "Test Manager Stern-Warning Notifications",
			input: entity.Manager{
				Id:                    1111,
				Name:                  "TestSuites 1",
				BirthDate:             time.Date(1990, time.December, 15, 0, 0, 0, 0, time.UTC),
				PerformancePercentage: 13,
			},
			expected: STERN_WARNING,
		},
		{
			name: "Test Manager Warning Notifications",
			input: entity.Manager{
				Id:                    2222,
				Name:                  "TestSuites 2",
				JoinDate:              time.Date(1970, time.January, 01, 0, 0, 0, 0, time.UTC),
				PerformancePercentage: 50,
			},
			expected: WARNING,
		},
	}

	for _, test := range testSuites {
		t.Run(test.name, func(t *testing.T) {
			result, _ := performanceWarningVisitor.VisitManager(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	notifier = factory.NewNotifierMock()
	performanceWarningVisitor = PerformanceWarningVisitor{Notif: notifier}
	m.Run()
}
