package defense

type Item struct {
	Name  string
	Price float64
}

func (i *Item) GetPrice() float64 {
	return i.Price
}

func (i *Item) GetName() string {
	return i.Name
}
