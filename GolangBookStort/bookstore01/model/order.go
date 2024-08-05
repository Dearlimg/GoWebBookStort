package model

type Order struct {
	OrderID     string
	CreateTime  string
	TotalCount  int64
	TotalAmount float64
	State       int64
	UserID      int64
}
