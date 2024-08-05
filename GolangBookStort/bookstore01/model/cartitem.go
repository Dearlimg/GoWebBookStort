package model

type CartItem struct {
	CartItemID int64
	Book       *Book
	Count      int64
	Amount     float64
	CartID     string
}

func (CartItem *CartItem) GetAmount() float64 {
	price := CartItem.Book.Price
	return float64(CartItem.Count) * price
}
