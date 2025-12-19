package Company

func Example() {
	company := NewCompany()

	ft := FullTimeEmployee{
		ID:     1,
		Name:   "erhan",
		Salary: 200000,
	}

	pt := PartTimeEmployee{
		ID:          2,
		Name:        "amir",
		PerHour:     5000,
		HoursWorked: 80,
	}

	company.AddEmployee(ft.ID, ft)
	company.AddEmployee(pt.ID, pt)

	company.ListEmployees()
}
