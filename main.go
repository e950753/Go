package main

import (
	"./employees"
	"./payroll"
	"fmt"
)

func main() {
	team := []employees.Employee{
		{
			Name: "John Doe",
			Position: "Developer",
			Skills: []string{"Go", "Python", "JavaScript"},
		},
		{
			Name: "Jane Doe",
			Position: "Designer",
			Skills: []string{"Photoshop", "Illustrator", "InDesign"},
		},
		{
			Name: "Jim Doe",
			Position: "Manager",
			Skills: []string{"Leadership", "Management", "Communication"},
		},
	}

	if !team[0].SetBaseSalary(1000) {
		fmt.Println("Failed to set base salary for", team[0].Name)
	}
	if !team[1].SetBaseSalary(1500) {
		fmt.Println("Failed to set base salary for", team[1].Name)
	}
	if !team[2].SetBaseSalary(2000) {
		fmt.Println("Failed to set base salary for", team[2].Name)
	}

	bonus := 0.1

	if bonus < 0 {
		fmt.Println("Bonus cannot be negative")
	}
	tax := 0.2
	if tax < 0 {
		fmt.Println("Tax cannot be negative")
	}

	for _, employee := range team {
		payroll.CalcGross(employee, bonus)
		payroll.CalcNet(employee, tax)
		fmt.Println("Employee:", employee.Name, "Gross:", payroll.CalcGross(employee, bonus), "Net:", payroll.CalcNet(employee, tax))
	}
	// employee.salary = 1000 // employee.salary undefined (cannot refer to unexported field salary)

	// Output:
	// Employee: John Doe Gross: 1100 Net: 960
	//Employee: Jane Doe Gross: 1650.0000000000002 Net: 1440
	// Employee: Jim Doe Gross: 2200 Net: 1920
}