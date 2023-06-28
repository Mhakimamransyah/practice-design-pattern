package main

import (
	"fmt"
	"time"
	visitor "visitor-pattern/concretevisitor"
	entities "visitor-pattern/entity"
)

func main() {

	fmt.Println("\n")

	employees := make([]entities.Employee, 0)

	performanceNotification := visitor.PerformanceWarningVisitor{}
	bonusIncome := visitor.BonusIncomeVisitor{}

	employees = append(employees,
		entities.Staff{
			Id:                    12345,
			Name:                  "Odonkor",
			JoinDate:              time.Date(2023, time.January, 10, 0, 0, 0, 0, time.UTC),
			BirthDate:             time.Date(1990, time.August, 02, 0, 0, 0, 0, time.UTC),
			PerformancePercentage: 85,
		}, entities.Staff{
			Id:                    12344,
			Name:                  "Deryl Salvio",
			JoinDate:              time.Date(2010, time.July, 07, 0, 0, 0, 0, time.UTC),
			BirthDate:             time.Date(1991, time.March, 23, 0, 0, 0, 0, time.UTC),
			PerformancePercentage: 27,
		},
		entities.Staff{
			Id:                    12343,
			Name:                  "Pedro Pacquita",
			JoinDate:              time.Date(2011, time.February, 10, 0, 0, 0, 0, time.UTC),
			BirthDate:             time.Date(1990, time.January, 07, 0, 0, 0, 0, time.UTC),
			PerformancePercentage: 47.5,
		},
		entities.Supervisor{
			Id:                    12342,
			Name:                  "Omar Maulo Atarez",
			JoinDate:              time.Date(2005, time.December, 15, 0, 0, 0, 0, time.UTC),
			BirthDate:             time.Date(1987, time.December, 02, 0, 0, 0, 0, time.UTC),
			PerformancePercentage: 75,
		},
		entities.Supervisor{
			Id:        12340,
			Name:      "Alexis Otario",
			JoinDate:  time.Date(2013, time.December, 15, 0, 0, 0, 0, time.UTC),
			BirthDate: time.Date(1996, time.January, 02, 0, 0, 0, 0, time.UTC),

			PerformancePercentage: 30,
		},
		entities.Manager{
			Id:                    12312,
			Name:                  "Zakaria Owele",
			JoinDate:              time.Date(1993, time.December, 04, 0, 0, 0, 0, time.UTC),
			BirthDate:             time.Date(1971, time.July, 15, 0, 0, 0, 0, time.UTC),
			PerformancePercentage: 50,
		},
	)

	for _, employee := range employees {

		// accept performance notification visitor
		employee.Accept(performanceNotification)

		// accept bonus income visitor
		employee.Accept(bonusIncome)

	}
}
