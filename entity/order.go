package entity

// order entity to represent order
type Order struct {
	OrderID     string
	UserID      string
	ProductName string
	Quantity    int
	SubTotal    float64
	GrandTotal  float64
	Status      string
}

var DummyOrders = []Order{
	{
		OrderID:     "abcd123",
		UserID:      "user1",
		ProductName: "shirt",
		Quantity:    2,
		SubTotal:    100,
		GrandTotal:  200,
		Status:      "Placed",
	},
	{
		OrderID:     "abcd124",
		UserID:      "user2",
		ProductName: "top",
		Quantity:    1,
		SubTotal:    100,
		GrandTotal:  100,
		Status:      "Shipped",
	},
}
