package dao

import (
	"book/model"
	"fmt"
	"testing"
)

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本书是%v", k, v)
	}
}


func testAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   59.31,
		Sales:   91,
		Stock:   75,
		ImgPath: "/static/img/default.jpg",
	}
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	DeleteBook("17")
}

func testGetBookByID(t *testing.T) {
	book, _ := GetBookById("6")
	fmt.Println("book = ", book)
}


func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      14,
		Title:   "六体",
		Author:  "磁芯刘",
		Price:   66.66,
		Sales:   100,
		Stock:   1,
		ImgPath: "/static/img/default.jpg",
	}
	UpdateBook(book)
}