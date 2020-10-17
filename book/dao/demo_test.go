package dao

import (
	"book/model"
	"book/utils"
	"fmt"
	"testing"
)

// 使用testing包进行测试
// 1.导入testing包

// 2.测试函数名前缀必须以Test开头,后缀名必须以大写字母.例如:
//func TestAdd(t *testing.T){ }
//func TestSum(t *testing.T){ }
//func TestLog(t *testing.T){ }

// 3.参数t用于报告测试失败和附加的日志信息
//func (c *T) Error(args ...interface{})
//func (c *T) Errorf(format string, args ...interface{})
//func (c *T) Fail()
//func (c *T) FailNow()
//func (c *T) Failed() bool
//func (c *T) Fatal(args ...interface{})
//func (c *T) Fatalf(format string, args ...interface{})
//func (c *T) Log(args ...interface{})
//func (c *T) Logf(format string, args ...interface{})
//func (c *T) Name() string
//func (t *T) Parallel()
//func (t *T) Run(name string, f func(t *T)) bool
//func (c *T) Skip(args ...interface{})
//func (c *T) SkipNow()
//func (c *T) Skipf(format string, args ...interface{})
//func (c *T) Skipped() bool

// 4.执行 go test 或 go test -v 查询详细信息

// 测试预编译执行
func TestAdd1(t *testing.T) {
	user := &model.User{
		Username: "DemoTest1",
		Password: "111",
		Email:    "demo1@126.com",
	}
	sql := "INSERT INTO users (username, password, email) VALUES(?, ?, ?)"
	// 预编译
	stmt, err := utils.Db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println("预编译错误")
	}
	// 执行
	_, err2 := stmt.Exec(user.Username, user.Password, user.Email)
	if err2 != nil {
		fmt.Println("执行错误")
	}
}

// 测试执行
func TestAdd2(t *testing.T) {
	user := &model.User{
		Username: "DemoTest2",
		Password: "222",
		Email:    "demo2@126.com",
	}
	sql := "INSERT INTO users (username, password, email) VALUES(?, ?, ?)"
	// 执行
	result, err := utils.Db.Exec(sql, user.Username, user.Password, user.Email)
	if err != nil {
		fmt.Println("执行错误")
	}
	// 插入后的自增id
	insertId, _ := result.LastInsertId() // 15 <nil>
	fmt.Println(insertId)
	// 受影响的行数
	fmt.Println(result.RowsAffected()) // 1 <nil>
}

func TestAdd3(t *testing.T) {
	user := &model.User{
		Username: "DemoTest3",
		Password: "333",
		Email:    "demo3@126.com",
	}
	err := AddUser(user.Username, user.Password, user.Email)
	if err != nil {
		err.Error()
	}
}