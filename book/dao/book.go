package dao

import  (
	"book/model"
	"book/utils"
	"fmt"
	"strconv"
)
// 获取所有图书
func GetBooks() ([]*model.Book, error) {
	// sql
	sql := "SELECT id, title, author, price, sales, stock, img_path FROM books"
	// 执行
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	// 定义一个切片
	var books []*model.Book
	// 遍历
	for rows.Next() {
		book := &model.Book{}
		err2 := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err2 != nil {
			return nil, err2
		}
		// 将book添加到books中
		books = append(books, book)
	}
	return books, nil
}

// Add
func AddBook(b *model.Book) error {
	sql := "INSERT INTO books(title, author, price, sales, stock, img_path) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// Update
func UpdateBook(b *model.Book) error {
	sql := "UPDATE books SET title = ?, author = ?, price = ?, sales = ?, stock = ? WHERE id = ?"
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func DeleteBook(bookId string) error {
	fmt.Println("bookId = ", bookId)
	sql := "DELETE FROM books WHERE id = ?"
	_, err := utils.Db.Exec(sql, bookId)
	if err != nil {
		return err
	}
	return nil
}

func GetBookById(bookID string) (*model.Book, error) {
	sql := "SELECT id,title,author,price,sales,stock,img_path FROM books WHERE id = ?"
	row := utils.Db.QueryRow(sql, bookID)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath) // 为book中的字段赋值
	return book, nil
}

// 带分页
func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64) // 将页码转成int64类型
	sql := "SELECT COUNT(*) FROM books"            // 获取数据库中的总记录数
	var totalRecord int64                          // 总的记录数
	row := utils.Db.QueryRow(sql)                  // 执行
	row.Scan(&totalRecord)

	var pageSize int64 = 4 // 每页显示4本书
	var totalPageNo int64  // 总页数
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	// 获取当前页中的图书
	sql2 := "SELECT id, title, author, price, sales, stock, img_path FROM books limit ?, ?"
	rows, err := utils.Db.Query(sql2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book) // 将book添加到books切片中
	}
	// 创建page
	data := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return data, nil
}
