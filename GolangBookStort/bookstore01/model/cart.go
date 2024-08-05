package model

type Cart struct {
	CartID      string
	CartItems   []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int
	UserName    string
}

func (Cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range Cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

func (Cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range Cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
