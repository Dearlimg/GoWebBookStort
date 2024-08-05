package dao

import (
	"bookstore01/model"
	"bookstore01/utils"
)

func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts (id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		AddCartItem(cartItem)
	}
	return nil
}

func GetCartByUserID(userID int) (cart *model.Cart, err error) {
	sqlStr := "select * from carts where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, userID)
	cart = &model.Cart{}
	row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	cart.CartItems = cartItems
	return cart, nil
}

func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count = ?, total_amount = ? where id = ?"
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartByCartID(cartID string) error {
	err := DeleteCartItemByCartID(cartID)
	if err != nil {
		return err
	}
	sqlStr := "delete from carts where id = ?"
	_, err2 := utils.Db.Exec(sqlStr, cartID)
	if err2 != nil {
		return err
	}
	return nil
}
