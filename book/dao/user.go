package dao

import (
	"book/model"
	"book/utils"
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
