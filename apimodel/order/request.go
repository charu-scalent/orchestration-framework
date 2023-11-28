package apimodel

type GetOrderRequest struct {
	UserID  string `form:"order_id"`
	OrderID string `form:"order_id"`
}
