package concretevisitor

import (
	"testing"
	"time"
	"visitor-pattern/entity"

	"github.com/stretchr/testify/assert"
)

func TestVisitStaffWarningNotification(t *testing.T) {

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
			result, _ := PerformanceWarningVisitor{}.VisitStaff(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestVisitSupervisorWarningNotification(t *testing.T) {

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
			result, _ := PerformanceWarningVisitor{}.VisitSupervisor(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestVisitManagerWarningNotification(t *testing.T) {

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
			result, _ := PerformanceWarningVisitor{}.VisitManager(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
