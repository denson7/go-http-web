package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

type Mysql struct {
	Db  *sql.DB
	Dsn string
}

func init() {
	//mysql := config.CreateMysql()
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysql.Username, mysql.Password, mysql.Server, mysql.Port, mysql.Database)

	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/books")
	if err != nil {
		panic(err.Error())
	}
}
