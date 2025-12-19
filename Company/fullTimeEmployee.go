package Company

import (
	"fmt"
)

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary int
}

func (f FullTimeEmployee) GetDetails() {
	fmt.Printf("ID: %d | Name: %s | Salary: %d Tenge \n", f.ID, f.Name, f.Salary)
}
