package Bank

import "fmt"

type BankAccount struct {
	Owner        string
	Balance      int
	Transactions []string
}

func (b *BankAccount) Deposit() {
	amount := 0
	fmt.Println("Please enter amount of money you want to deposit")
	_, err := fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Amount of money to deposit need to be more than 0")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	b.Balance += amount

	b.Transactions = append(b.Transactions, fmt.Sprintf("deposited %d Tenge", amount))
	fmt.Printf("You have successfully deposited %d Tenge", amount)
}

func (b *BankAccount) Withdraw() {
	amount := 0
	fmt.Println("Please enter amount of money you want to withdraw")
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println(err)
		return
	}

	if amount <= 0 || amount > b.Balance {
		fmt.Println("Amount of money to withdraw need to be more than 0 and less or equal to your balance")
		return
	}
	b.Balance -= amount

	b.Transactions = append(b.Transactions, fmt.Sprintf("withdrawed %d Tenge", amount))
	fmt.Printf("You have successfully withdrawed %d Tenge", amount)
}

func (b *BankAccount) GetBalance() {
	fmt.Printf("Your balance is %d Tenge", b.Balance)
}

func (b *BankAccount) GetTransactions() {
	if len(b.Transactions) == 0 {
		fmt.Println("You have no transactions")
		return
	}
	for i := 0; i < len(b.Transactions); i++ {
		fmt.Printf("%d.%s\n", i+1, b.Transactions[i])
	}
}
