package dao

import (
	"bookstore01/model"
	"bookstore01/utils"
	"fmt"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	fmt.Println(orderItem)
	sql := "INSERT INTO order_items(id, count, amount, title, author, price, img_path, order_id) VALUES(?,?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, orderItem.OrderItemID, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}
