package Company

import "fmt"

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	PerHour     int
	HoursWorked int
}

func (p PartTimeEmployee) GetDetails() {
	fmt.Printf("ID: %d | Name: %s | Per hour: %d Tenge | Hours worked: %d \n", p.ID, p.Name, p.PerHour, p.HoursWorked)
}
