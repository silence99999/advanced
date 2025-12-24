package defense

import "fmt"

func orderCLI() {
	number := -1
	store := NewStore()
	str := `
		--------------------------
		Order system
		1.Create Order
		2.Add Item to Order
		3.Remove Item from Order
		4.List Orders
		5.Show Total Revenue
		6.Exit
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
			store.AddOrder()
		case 2:
			store.AddItem()
		case 3:
			store.RemoveItem()
		case 4:
			store.ListOrders()
		case 5:
			store.TotalRevenue()
		case 6:
			fmt.Println("Thank you for using library system")
			return
		default:
			fmt.Println("You need to choose between 1-4")
		}

	}
}
