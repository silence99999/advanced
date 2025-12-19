package Company

import "fmt"

type Company struct {
	Employees map[uint64]Employee
}

func (c *Company) AddEmployee(id uint64, employee Employee) {
	c.Employees[id] = employee
}

func (c *Company) ListEmployees() {
	if len(c.Employees) == 0 {
		fmt.Println("There is no Employees")
		return
	}
	for _, v := range c.Employees {
		v.GetDetails()
	}
}

func NewCompany() *Company {
	return &Company{
		Employees: make(map[uint64]Employee),
	}
}
