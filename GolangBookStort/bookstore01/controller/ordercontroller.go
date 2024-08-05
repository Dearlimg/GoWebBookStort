package controller

import (
	"bookstore01/dao"
	"bookstore01/model"
	"bookstore01/utils"
	"html/template"
	"net/http"
	"time"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	orderID := utils.CreatUUID()
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	dao.AddOrder(order)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderID,
		}
		dao.AddOrderItem(orderItem)
		book := v.Book
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		dao.UpdateBook(book)

	}
	dao.DeleteCartByCartID(cart.CartID)
	session.OrderID = orderID
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, session)
}
