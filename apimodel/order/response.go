package apimodel

type GetOrderResponse struct {
	OrderID     string  `json:"id"`
	UserID      string  `json:"user_id"`
	ProductName string  `json:"product"`
	Quantity    int     `json:"quantity"`
	SubTotal    float64 `json:"subt_total"`
	GrandTotal  float64 `json:"grand_total"`
	Status      string  `json:"status"`
}
