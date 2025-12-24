package defense

import (
	"fmt"
	"slices"
)

type Order struct {
	ID    uint64
	Items []Item
}

func (o *Order) CreateOrder() {
	name := readLine("Please enter the item name")
	fmt.Println("Enter the price")
	var price float64
	_, err := fmt.Scanln(&price)
	if err != nil {
		fmt.Println(err)
		return
	}
	o.Items = append(o.Items, Item{
		Name:  name,
		Price: price,
	})

}

func (o *Order) DeleteItem(itemName string) {
	for i := 0; i < len(o.Items); i++ {
		if o.Items[i].GetName() == itemName {
			o.Items = slices.Delete(o.Items, i, i)
		}
	}
}
