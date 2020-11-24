package dao

import (
	"book/model"
	"book/utils"
	"fmt"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	user := &model.User{}
	sql := "SELECT id, username, password, email FROM users WHERE username = ? AND password = ?"
	// 查询一行
	row := utils.Db.QueryRow(sql, username, password)
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	user := &model.User{}
	sql := "SELECT id, username, password, email FROM users WHERE username = ?"
	row := utils.Db.QueryRow(sql, username)
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func AddUser(username string, password string, email string) error {
	sql := "INSERT INTO users(username, password, email) VALUES (?, ?, ?)"
	_, err := utils.Db.Exec(sql, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
type User struct {
	UserName string
	Password string
}

func TestTx() bool {
	user := User{
		UserName: "John",
		Password: "1234",
	}
	// 开启事务
	tx, err := utils.Db.Begin()
	if err != nil {
		fmt.Println("Tx failed")
		return false
	}
	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO test (`name`,`password`) VALUES (?, ?)")
	//stmt, err := tx.Prepare("UPDATE test SET name = ?, password = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare failed")
		return false
	}
	// 将参数传递到sql语句中并执行
	result, err := stmt.Exec(user.UserName, user.Password)
	if err != nil {
		fmt.Println("Exec failed")
		return false
	}
	// 提交事务
	tx.Commit()
	// 获取上一次插入的自增Id
	fmt.Println(result.LastInsertId())
	return true
}

func TestSelect() []User {
	// 执行查询语句
	rows, err := utils.Db.Query("SELECT name, password FROM test")
	if err != nil {
		fmt.Println("query failed")
	}
	var users []User
	// 循环读取结果
	for rows.Next() {
		var user User
		// 将每一行的结果赋值给一个user对象
		err := rows.Scan(&user.UserName, &user.Password)
		if err != nil {
			fmt.Println("row fail")
		}
		// 将user追到到users这个数组种
		users = append(users, user)
	}
	return users
}
