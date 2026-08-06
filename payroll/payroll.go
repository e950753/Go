package payroll

import "../employees"

func CalcGross(e employees.Employee, bonusPct float64) float64 {
	return e.BaseSalary() * (1 + bonusPct)
}

func CalcNet(e employees.Employee, taxPct float64) float64 {
	return CalcGross(e, taxPct) * (1 - taxPct)
}