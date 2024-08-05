package dao

import (
	"bookstore01/model"
	"bookstore01/utils"
)

func AddOrder(order *model.Order) error {
	sql := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}
