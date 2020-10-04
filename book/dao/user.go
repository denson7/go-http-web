package dao

import (
	"book/model"
	"book/utils"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sql := "SELECT id, username, password, email FROM users WHERE username = ? AND password = ?"
	row := utils.Db.QueryRow(sql, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	sql := "SELECT id, username, password, email FROM users WHERE username = ?"
	row := utils.Db.QueryRow(sql, username)
	user := &model.User{}
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
