package Bank

import "fmt"

func CLIBank() {
	number := -1
	account := BankAccount{
		Owner: "Amir",
	}
	str := `
		--------------------------
		Bank System
		1.Deposit
		2.Withdraw
		3.Get balance
		4.Get transactions
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
			account.Deposit()
		case 2:
			account.Withdraw()
		case 3:
			account.GetBalance()
		case 4:
			account.GetTransactions()
		case 0:
			fmt.Println("Thank you for using bank account system")
			return
		default:
			fmt.Println("You need to choose between 1-4")
		}
	}
}
