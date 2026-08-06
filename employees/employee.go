package employees

type Employee struct {
	Name string
	Position string
	salary float64
	Skills []string
}

func (e *Employee) SetBaseSalary(salary float64) bool {
	if salary > 0 {
		e.salary = salary
		return true
	}
	return false
}

func (e *Employee) BaseSalary() float64 {
	return e.salary
}