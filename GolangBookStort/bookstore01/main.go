package main

import (
	"bookstore01/controller"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	//http.HandleFunc("/main", controller.IndexHandler)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/login", controller.Login)
	//http.HandleFunc("/AddBook", controller.AddBook)
	//http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/toUpdataBookPage", controller.ToUpdataBookPage)
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	http.HandleFunc("/checkout", controller.Checkout)

	http.ListenAndServe(":8090", nil)
}
