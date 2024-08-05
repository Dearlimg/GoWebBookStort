package controller

import (
	"bookstore01/dao"
	"bookstore01/model"
	"bookstore01/utils"
	"html/template"
	"net/http"
	"strconv"
)

//func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
//	bookId := r.FormValue("bookId")
//	book, _ := dao.GetBookByID(bookId)
//	//fmt.Println(bookId)
//	_, session := dao.IsLogin(r)
//	//fmt.Println(session)
//	userID := session.UserID
//	//fmt.Println(userID)
//	cart, _ := dao.GetCartByUserID(userID)
//	//fmt.Println("督导的cart", cart)
//	if cart.CartItems != nil {
//		fmt.Println("cart不为空", bookId, cart.CartID)
//		cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookId, cart.CartID)
//		fmt.Println(cartItem)
//		if cartItem != nil {
//			fmt.Println("测试这里")
//			//购物车中已经有东西了,只要加就好
//			cts := cart.CartItems
//			for _, v := range cts {
//				if cartItem.Book.ID == v.Book.ID {
//					v.Count = v.Count + 1
//					fmt.Println(v.Count)
//					dao.UpdateBookCount(v.Count, v.Book.ID, cart.CartID)
//				}
//			}
//		} else {
//			//购物车中没有这个商品,需要创造购物下个,提交到数据库中
//			cartItem := &model.CartItem{
//				Book:   book,
//				Count:  1,
//				CartID: cart.CartID,
//			}
//			cart.CartItems = append(cart.CartItems, cartItem)
//			dao.AddCartItem(cartItem)
//		}
//		dao.UpdateCart(cart)
//	} else {
//		fmt.Println("没有购物车")
//		cartID := utils.CreatUUID()
//		cart = &model.Cart{
//			CartID: cartID,
//			UserID: userID,
//		}
//		var cartItems []*model.CartItem
//		cartItem := &model.CartItem{
//			Book:   book,
//			Count:  1,
//			CartID: cartID,
//		}
//		cartItems = append(cartItems, cartItem)
//		cart.CartItems = cartItems
//		dao.AddCart(cart)
//		fmt.Println("成功加入购物车")
//	}
//}

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(bookId)
	flag, session := dao.IsLogin(r)
	if flag {
		bookId := r.FormValue("bookId")
		book, _ := dao.GetBookByID(bookId)
		//fmt.Println(session)
		userID := session.UserID
		//fmt.Println(userID)
		cart, _ := dao.GetCartByUserID(userID)
		//fmt.Println("督导的cart", cart)
		if cart.CartItems != nil {
			//fmt.Println("cart不为空", bookId, cart.CartID)
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookId, cart.CartID)
			//fmt.Println(cartItem)
			if cartItem != nil {
				//fmt.Println("测试这里")
				//购物车中已经有东西了,只要加就好
				cts := cart.CartItems
				for _, v := range cts {
					if cartItem.Book.ID == v.Book.ID {
						v.Count = v.Count + 1
						//fmt.Println(v.Count, v.Book.ID, cart.CartID)
						dao.UpdateBookCount(v)
					}
				}
			} else {
				//购物车中没有这个商品,需要创造购物下个,提交到数据库中
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				cart.CartItems = append(cart.CartItems, cartItem)
				dao.AddCartItem(cartItem)
			}
			dao.UpdateCart(cart)
		} else {
			//fmt.Println("没有购物车")
			cartID := utils.CreatUUID()
			cart = &model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}
			cartItems = append(cartItems, cartItem)
			cart.CartItems = cartItems
			dao.AddCart(cart)
			//fmt.Println("成功加入购物车")
		}
		w.Write([]byte("刚刚将" + book.Title + "加入到了购物车"))
	} else {
		w.Write([]byte("请先登录"))
	}
}

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	//fmt.Println(cart)
	if cart != nil {
		cart.UserName = session.UserName
		session.Cart = cart
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//fmt.Println("success", session)
		t.Execute(w, session)
	} else {
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//fmt.Println("lose", session)
		t.Execute(w, session)
	}
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartId")
	dao.DeleteCartByCartID(cartID)
	GetCartInfo(w, r)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemID == iCartItemID {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	dao.UpdateCart(cart)
	GetCartInfo(w, r)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	bookCount := r.FormValue("bookCount")
	ibookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.CartItemID == iCartItemID {
			v.Count = ibookCount
			dao.UpdateBookCount(v)
		}
	}
	dao.UpdateCart(cart)
	GetCartInfo(w, r)
}
