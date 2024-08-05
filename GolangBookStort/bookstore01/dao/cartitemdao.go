package dao

import (
	"bookstore01/model"
	"bookstore01/utils"
	"fmt"
)

func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}

	return nil
}

func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select * from cart_items where book_id=? AND cart_id=?"
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	cartItem := &model.CartItem{}
	var bookId int
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartID)
	//fmt.Println(cartItem)
	if err != nil {
		return nil, err
	}
	book, _ := GetBookByID(bookID)
	cartItem.Book = book
	//fmt.Println("数据库查询到的数据", cartItem)
	return cartItem, nil
}

func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select * from cart_items where cart_id = ?"
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		fmt.Println(err)
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID string
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err2 != nil {
			return nil, err
		}
		book, _ := GetBookByID(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//func GetCartByUserID(bookID int) (*model.Cart, error) {
//	sqlStr := "select * from cart_items where book_id = ?"
//	row := utils.Db.QueryRow(sqlStr, bookID)
//	cart := &model.Cart{}
//	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.CartID)
//	if err != nil {
//		return nil, err
//	}
//	//获取购物想
//	cartItems, _ := GetCartItemsByCartID(cart.CartID)
//	cart.CartItems = cartItems
//	return cart, nil
//}

func UpdateBookCount(cartItem *model.CartItem) error {
	//sql := "update cart_items set count = ? where book_id = ?and cart_id=?"
	sql := "update cart_items set count = ? , amount = ? where book_id = ? and cart_id = ?"
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartItemByCartID(cartID string) error {
	sql := "delete from cart_items where cart_id = ?"
	_, err := utils.Db.Exec(sql, cartID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartItemByID(cartItemID string) error {
	sql := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sql, cartItemID)
	if err != nil {
		return err
	}
	return nil
}
