package controller

import (
	"bookstore01/dao"
	"bookstore01/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	price := r.FormValue("price")
	sales := r.FormValue("sales")
	stock := r.FormValue("stock")

	fPrice, _ := strconv.ParseFloat(price, 32)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	istock, _ := strconv.ParseInt(stock, 10, 0)

	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(istock),
		ImgPath: "static/img/default.jpg",
	}

	dao.AddBook(book)

	GetPageBooks(w, r)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	dao.DeleteBook(bookID)
	GetPageBooks(w, r)
}

func ToUpdataBookPage(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
	t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
	t.Execute(w, book)
}

func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	BookID := r.FormValue("bookId")
	title := r.FormValue("title")
	author := r.FormValue("author")
	price := r.FormValue("price")
	sales := r.FormValue("sales")
	stock := r.FormValue("stock")

	fPrice, _ := strconv.ParseFloat(price, 32)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	istock, _ := strconv.ParseInt(stock, 10, 0)
	ibookID, _ := strconv.ParseInt(BookID, 10, 0)

	book := &model.Book{
		ID:      int(ibookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(istock),
		ImgPath: "static/img/default.jpg",
	}

	if book.ID > 0 {
		dao.UpdateBook(book)
	} else {
		dao.AddBook(book)
	}
	GetPageBooks(w, r)
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	fmt.Println(page)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	fmt.Println(page)
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.UserName
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	fmt.Println(page)
	t.Execute(w, page)
}
