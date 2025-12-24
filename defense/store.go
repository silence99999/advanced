package defense

import "fmt"

type Store struct {
	Orders map[uint64]Order
}

var ErrNotFound = "order not found with this id"

func (s *Store) AddOrder(order Order) error {
	var maxID uint64
	if len(s.Orders) == 0 {
		maxID = 1
	}

	for key := range s.Orders {
		if key > maxID {
			maxID = key
		}
	}

	order.CreateOrder()

	s.Orders[maxID] = order

	return nil
}

func (s *Store) RemoveItem(orderID uint64, itemName string) error {
	if orderID <= 0 {
		fmt.Println(ErrNotFound)
		return nil
	}

	//
	//order := Order{}
	//
	//for _, value := range s.Orders {
	//
	//	if s.Orders[orderID].ID == value.ID {
	//		break
	//	}
	//	for i := 0;i < len(value.Items);i++{
	//		if value.Items[i] ==
	//	}
	//}
	//if s.Orders[orderID].Items

	orders

	delete(s.Orders, orderID)

	return nil

}

func (s *Store) ListOrders() []Order {
	slc := []Order{}

	for _, value := range s.Orders {
		slc = append(slc, value)
	}

	return slc
}

func (s *Store) TotalRevenue() float64 {
	var total float64
	for _, value := range s.Orders {
		for i := 0; i < len(value.Items); i++ {
			total += value.Items[i].GetPrice()
		}
	}
	return total
}

func NewStore() *Store {
	return &Store{
		Orders: make(map[uint64]Order),
	}
}
