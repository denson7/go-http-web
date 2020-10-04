package controller

import (
	"book/dao"
	"book/model"
	"fmt"
	"net/http"
	"strconv"
)

func GetBookList(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本书是%v", k, v)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	err := dao.DeleteBook(bookId)
	if err != nil {
		//
	}
	Index(w, r)
}


func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	book, _ := dao.GetBookById(bookId)
	if book.ID > 0 {
		//
	} else {
		//
	}
}

// UpdateOrAddBook 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	// 获取POST提交数据
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	// 数据类型转换
	fPrice, _ := strconv.ParseFloat(price, 64) // 将价格、销量、库存进行转换
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	iBookID, _ := strconv.ParseInt(bookID, 10, 0)
	// 创建book
	book := &model.Book{
		ID:      int(iBookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   iSales,
		Stock:   iStock,
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		// update
		dao.UpdateBook(book)
	} else {
		// add
		dao.AddBook(book)
	}
	Index(w, r)
}
