package Library

import "fmt"

func CLILibrary() {
	number := -1
	library := NewLibrary()
	str := `
		--------------------------
		Library System
		1.Add book
		2.Borrow book
		3.Return book
		4.List all available books
		0.Exit
		--------------------------
`
	for {
		fmt.Println(str)
		fmt.Println("Please enter a number")
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch number {
		case 1:
			library.AddBook()
		case 2:
			library.BorrowBook()
		case 3:
			library.ReturnBook()
		case 4:
			library.ListAvailableBooks()
		case 0:
			fmt.Println("Thank you for using library system")
			return
		default:
			fmt.Println("You need to choose between 1-4")
		}

	}
}
