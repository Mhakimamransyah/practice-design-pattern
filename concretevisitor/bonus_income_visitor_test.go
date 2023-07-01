package concretevisitor

import (
	"testing"

	"github.com/Mhakimamransyah/practice-design-pattern/entity"

	"github.com/stretchr/testify/assert"
)

func TestVisitStaffBonusIncome(t *testing.T) {

	testSuites := []struct {
		name     string
		input    entity.Staff
		expected interface{}
	}{
		{
			name: "Test Staff Bonus Income",
			input: entity.Staff{
				Id:                    1111,
				Name:                  "TestSuites 1",
				PerformancePercentage: 88,
			},
			expected: STAFF_BONUS_INCOME,
		},
		{
			name: "Test Staff With no Bonus Income",
			input: entity.Staff{
				Id:                    2222,
				Name:                  "TestSuites 2",
				PerformancePercentage: 50,
			},
			expected: 0,
		},
	}

	for _, test := range testSuites {
		t.Run(test.name, func(t *testing.T) {
			result, _ := BonusIncomeVisitor{}.VisitStaff(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestVisitSupervisorBonusIncome(t *testing.T) {

	testSuites := []struct {
		name     string
		input    entity.Supervisor
		expected interface{}
	}{
		{
			name: "Test Supervisor Bonus Income",
			input: entity.Supervisor{
				Id:                    1111,
				Name:                  "TestSuites 1",
				PerformancePercentage: 85,
			},
			expected: SUPERVISOR_BONUS_INCOME,
		},
		{
			name: "Test Supervisor With no Bonus Income",
			input: entity.Supervisor{
				Id:                    2222,
				Name:                  "TestSuites 2",
				PerformancePercentage: 69,
			},
			expected: 0,
		},
	}

	for _, test := range testSuites {
		t.Run(test.name, func(t *testing.T) {
			result, _ := BonusIncomeVisitor{}.VisitSupervisor(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestVisitManagerBonusIncome(t *testing.T) {

	testSuites := []struct {
		name     string
		input    entity.Manager
		expected interface{}
	}{
		{
			name: "Test Manager Bonus Income",
			input: entity.Manager{
				Id:                    1111,
				Name:                  "TestSuites 1",
				PerformancePercentage: 90,
			},
			expected: MANAGER_BONUS_INCOME,
		},
		{
			name: "Test Manager With no Bonus Income",
			input: entity.Manager{
				Id:                    2222,
				Name:                  "TestSuites 2",
				PerformancePercentage: 59,
			},
			expected: 0,
		},
	}

	for _, test := range testSuites {
		t.Run(test.name, func(t *testing.T) {
			result, _ := BonusIncomeVisitor{}.VisitManager(&test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
