package models

// Item model
type Item struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	OrderID     int64  `json:"order_id"`
}
