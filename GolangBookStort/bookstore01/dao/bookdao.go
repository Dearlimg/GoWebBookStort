package dao

import (
	"bookstore01/model"
	"bookstore01/utils"
	"fmt"
	"strconv"
)

func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	var books []*model.Book
	for rows.Next() {
		//var book model.Book
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	fmt.Println(books)
	return books, nil
}

func AddBook(book *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func DeleteBook(bookID string) error {
	sqlStr := "delete from books where id=?"
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

func GetBookByID(bookID string) (*model.Book, error) {
	sqlStr := "select * from books where id=?"
	row := utils.Db.QueryRow(sqlStr, bookID)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

func UpdateBook(book *model.Book) error {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sqlStr := "select count(*) from books"
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//fmt.Println(totalRecord)
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	sqlStr2 := "select * from books limit ?,?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println(err)
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//fmt.Println(book)
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	//fmt.Println(page)
	//fmt.Println(page.Books)
	return page, nil
}

func GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sqlStr := "select count(*) from books where price between ? and ?"
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	fmt.Println(totalRecord)
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	sqlStr2 := "select * from books where price between ? and ? limit ?,?"
	rows, err := utils.Db.Query(sqlStr2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println(err)
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		fmt.Println(book)
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	fmt.Println(page)
	fmt.Println(page.Books)
	return page, nil
}
